package commands

import (
	"github.com/disgoorg/disgo/discord"
)

var Commands = []discord.ApplicationCommandCreate{
	discord.SlashCommandCreate{
		Name:        "ping",
		Description: "Respond with Pong!",
	},
	discord.SlashCommandCreate{
		Name:        "collection",
		Description: "Edit own collection or check someone else's",
	},
	discord.SlashCommandCreate{
		Name:        "syncjson",
		Description: "Sync items JSON with the database",
	},
}