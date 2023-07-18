package gatari

import (
	"context"
	"fmt"
	"strconv"

	"github.com/EnergoStalin/osubot/internal/modules"
	"github.com/EnergoStalin/osubot/internal/services"
	"github.com/EnergoStalin/osubot/pkg/gatari"
	"github.com/bwmarrin/discordgo"
	"github.com/go-resty/resty/v2"
)

type GatariModule struct {
	client *gatari.GatariClient
}

func (g *GatariModule) Init() (err error) {
	g.client = &gatari.GatariClient{
		Client: resty.New(),
	}

	return
}

func (g *GatariModule) GetName() string {
	return "Gatari"
}

func (g *GatariModule) GetCommands() modules.ModuleCommands {
	return modules.ModuleCommands{
		{
			Name:        "recent",
			Description: "Get recent relax score on gatari server for linked account",
		},
		{
			Name:        "link",
			Description: "Linking gatari accout to discord",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "id",
					Type:        discordgo.ApplicationCommandOptionInteger,
					Description: "Gatari user id",
					Required:    true,
				},
			},
		},
	}
}

func (g *GatariModule) GetCallbacks() modules.InteractionCallbacks {
	return modules.ReflectCallbacks[*GatariModule](g)
}

func (g *GatariModule) checkLink(id string) bool {
	return services.KVStore.Get(id) == ""
}

func (g *GatariModule) getLastScore(sid string, special int) (s gatari.Score, err error) {
	id, _ := strconv.Atoi(sid)
	scores, err := g.client.GetUserRecentScores(&gatari.GetUserRecentScoresConfig{
		F:        1,
		Mode:     0,
		L:        1,
		PpFilter: 0,
		Id:       id,
		Special:  special,
	}, context.Background())
	if err != nil {
		return
	}
	s = scores.Scores[0]

	return
}

func (g *GatariModule) Recent(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	var msg string
	if g.checkLink(i.Member.User.ID) {
		msg = "Account not linked!"
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &msg,
		})
		return
	}

	id := services.KVStore.Get(i.Member.User.ID)

	sc, err := g.getLastScore(id, 1)
	if err != nil {
		msg = err.Error()
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &msg,
		})
		return
	}

	bm := sc.Beatmap
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{
			{
				Type: "rich",
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL:    fmt.Sprintf("https://b.ppy.sh/thumb/%dl.jpg", bm.BeatmapsetID),
					Width:  160,
					Height: 120,
				},
				Color: i.Member.User.AccentColor,
				Author: &discordgo.MessageEmbedAuthor{
					URL:     fmt.Sprintf("https://osu.ppy.sh/b/%d", bm.BeatmapID),
					Name:    fmt.Sprintf("%s [%s] [%.2f★]", bm.Title, bm.Version, bm.Difficulty),
					IconURL: fmt.Sprintf("https://a.gatari.pw/%s", id),
				},
				Description: fmt.Sprintf("▸ %.2fPP ▸ %.2f%%\n▸ %d ▸ %dx ▸ [%d/%d/%d/%d]", sc.Pp, sc.Accuracy, sc.Score, sc.MaxCombo, sc.Count300, sc.Count100, sc.Count50, sc.CountMiss),
			},
		},
	})
}

func (g *GatariModule) Link(s *discordgo.Session, i *discordgo.InteractionCreate) {
	o := i.ApplicationCommandData().Options[0]
	id := fmt.Sprint(o.IntValue())

	services.KVStore.Put(i.Member.User.ID, id)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Linked to " + id,
		},
	})
}
