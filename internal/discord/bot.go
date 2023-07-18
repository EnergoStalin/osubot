package discord

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/EnergoStalin/osubot/internal/modules"
	"github.com/EnergoStalin/osubot/internal/modules/gatari"
	"github.com/EnergoStalin/osubot/internal/modules/pixiv"
	"github.com/bwmarrin/discordgo"
)

func Run() (err error) {
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	mm := modules.NewModuleManager()

	dg.AddHandler(mm.Invoke)
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentMessageContent

	mm.RegiesterModule(&pixiv.PixivModule{})
	mm.RegiesterModule(&gatari.GatariModule{})

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	mm.RegisterCommands(dg)

	fmt.Println("Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	mm.UnRegisterAllCommands(dg)

	dg.Close()
	return
}
