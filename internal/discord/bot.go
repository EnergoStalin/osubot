package discord

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/EnergoStalin/osubot/internal/discord/commands"
	"github.com/EnergoStalin/osubot/internal/discord/handlers"
	"github.com/bwmarrin/discordgo"
)

func Run() (err error) {
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(handlers.MessageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentMessageContent

	commands.RegisterCommands()

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	return
}
