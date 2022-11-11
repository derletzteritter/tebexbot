package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	tbc "github.com/itschip/tebexbot/commands"
	"github.com/itschip/tebexbot/internal"
	"github.com/itschip/tebexgo"
)

var (
	BotToken    = internal.EnvVariable("BOT_TOKEN")
	GuildID     = internal.EnvVariable("GUILD_ID")
	TebexSecret = internal.EnvVariable("TEBEX_SECRET")
)

var (
	s           *discordgo.Session
	ts          *tebexgo.Session
	allPackages string
)

func main() {
	ts = tebexgo.New(TebexSecret)

	var err error
	s, err = discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err.Error())
	}

	commands := []*discordgo.ApplicationCommand{
		tbc.RegisterCheckoutCommand(ts),
		tbc.RegisterPackageCommand(),
	}

	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"checkout": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			tbc.CreateCheckoutCommand(ts, s, i)
		},
		"packages": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			tbc.CreatePackageCommand(ts, s, i)
		},
		"ban": func(s *discordgo.Session, i *discordgo.InteractionCreate) {},
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
