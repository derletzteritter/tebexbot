package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/itschip/tebexgo"
)

func CreatePackageCommand(ts *tebexgo.Session, s *discordgo.Session, i *discordgo.InteractionCreate) {
	packages, err := ts.GetAllPackages()
	if err != nil {
		log.Fatalln(err.Error())
	}

	fields := make([]*discordgo.MessageEmbedField, 0)

	for _, p := range packages {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  p.Name,
			Value: fmt.Sprint(p.Price),
		})
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Just testing",
			Embeds: []*discordgo.MessageEmbed{
				{
					Fields: fields,
				},
			},
		},
	})
}

func RegisterPackageCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "packages",
		Description: "List all packages",
	}
}
