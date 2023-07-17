package commands

import "github.com/EnergoStalin/osubot/internal/discord/handlers"

func RegisterCommands() {
	handlers.RegisterCommand("rs", Rs)
	handlers.RegisterCommand("bind", Bind)
	handlers.RegisterCommand("pixiv", Pixiv)
}
