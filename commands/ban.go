package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/itschip/tebexgo"
)

func RegisterBanCommand(ts *tebexgo.Session) *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "ban",
		Description: "Ban a user from the tebex store.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "user",
				Description: "The username or UUID of the player to ban",
				Required:    true,
				Type:        discordgo.ApplicationCommandOptionString,
			},
			{
				Description: "The reason for the ban",
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Required:    true,
				Choices:     CreateCheckoutCommandChoices(ts),
			},
			{
				Description: "The IP address to also ban",
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "ip",
				Required:    false,
				Choices:     CreateCheckoutCommandChoices(ts),
			},
		},
	}
}
