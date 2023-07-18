package pixiv

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/EnergoStalin/osubot/internal/modules"
	"github.com/bwmarrin/discordgo"
	"github.com/codingconcepts/env"
	"github.com/everpcpc/pixiv"
)

type PixivModule struct {
	app *pixiv.AppPixivAPI
	acc *pixiv.Account

	id uint64
	mu sync.Mutex

	c struct {
		Token        string `env:"PIXIV_TOKEN" required:"true"`
		RefreshToken string `env:"PIXIV_REFRESH_TOKEN" required:"true"`
	}
}

func (p *PixivModule) Init() (err error) {
	err = env.Set(&p.c)
	if err != nil {
		return
	}
	p.app = pixiv.NewApp()
	p.acc, err = pixiv.LoadAuth(p.c.Token, p.c.RefreshToken, time.Now().Add(3600))
	if err != nil {
		return
	}
	id, _ := strconv.ParseInt(p.acc.ID, 10, 0)
	p.id = uint64(id)

	return
}

func (p *PixivModule) GetName() string {
	return "Pixiv"
}

func (p *PixivModule) GetCommands() modules.ModuleCommands {
	return modules.ModuleCommands{
		{
			Name:        "pixiv",
			Description: "Get random illustration from EnergoStalin's bookmarks",
		},
	}
}

func (d *PixivModule) GetCallbacks() modules.InteractionCallbacks {
	return modules.ReflectCallbacks[*PixivModule](d)
}

func (p *PixivModule) Pixiv(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	p.mu.Lock()
	ia, err := fetchAll(p.app, p.id)
	p.mu.Unlock()

	if err != nil {
		msg := err.Error()
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &msg,
		})
		return
	}

	ri := ia[rand.Int()%len(ia)]

	msg := fmt.Sprintf("https://www.pixiv.net/en/artworks/%d", ri.ID)
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &msg,
	})
}
