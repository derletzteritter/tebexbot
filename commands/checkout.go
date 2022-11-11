package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/itschip/tebexgo"
)

func CreateCheckoutCommand(ts *tebexgo.Session, s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionsMap[opt.Name] = opt
	}

	packageId, ok := optionsMap["package-id"]
	user, ok := optionsMap["user"]

	if !ok {
		log.Println("Failed to find package")
	}
	checkoutObject := &tebexgo.PutCheckoutObject{
		PackageId: packageId.StringValue(),
		Username:  user.StringValue(),
	}

	checkout, err := ts.CreateCheckoutUrl(checkoutObject)
	if err != nil {
		log.Fatal(err.Error())
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Creating a checkout link for the package. URL: " + checkout.Url,
		},
	})
}

func CreateCheckoutCommandChoices(ts *tebexgo.Session) []*discordgo.ApplicationCommandOptionChoice {
	pkgs, _ := ts.GetAllPackages()

	c := make([]*discordgo.ApplicationCommandOptionChoice, 0)
	for _, p := range pkgs {
		c = append(c, &discordgo.ApplicationCommandOptionChoice{
			Name:  p.Name,
			Value: fmt.Sprint(p.Id),
		})
	}

	return c
}

func RegisterCheckoutCommand(ts *tebexgo.Session) *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "checkout",
		Description: "Creates a checkout url for a package",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "user",
				Description: "The username which is to buy the package",
				Required:    true,
				Type:        discordgo.ApplicationCommandOptionString,
			},
			{
				Description: "The package to buy",
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "package-id",
				Required:    true,
				Choices:     CreateCheckoutCommandChoices(ts),
			},
		},
	}
}
