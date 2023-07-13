package commands

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/EnergoStalin/osubot/internal/services"
	"github.com/EnergoStalin/osubot/pkg/gatari"
	"github.com/bwmarrin/discordgo"
)

func Rs(s *discordgo.Session, m *discordgo.MessageCreate) {
	sa := strings.Split(m.Content, " ")
	sid := services.KVStore.Get(m.Author.ID)
	if sid == "" {
		s.ChannelMessageSendReply(m.ChannelID, "Set user id first. via !bind <userid>", m.Reference())
		return
	}

	id, _ := strconv.Atoi(sid)

	special := 0
	if len(sa) > 1 && (sa[1] == "-relax" || sa[1] == "-rx") {
		special = 1
	}

	scores, err := services.GatariClient.GetUserRecentScores(&gatari.GetUserRecentScoresConfig{
		F:        1,
		Mode:     0,
		L:        1,
		PpFilter: 0,
		Id:       id,
		Special:  special,
	}, context.TODO())
	if err != nil {
		s.ChannelMessageSendReply(m.ChannelID, "Error occured durign retriving user recent scores...", m.Reference())
		return
	}

	bm := scores.Scores[0].Beatmap
	sc := scores.Scores[0]

	s.ChannelMessageSendEmbedReply(m.ChannelID, &discordgo.MessageEmbed{
		Type: "rich",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL:    fmt.Sprintf("https://b.ppy.sh/thumb/%dl.jpg", bm.BeatmapsetID),
			Width:  160,
			Height: 120,
		},
		Color: m.Author.AccentColor,
		Author: &discordgo.MessageEmbedAuthor{
			URL:     fmt.Sprintf("https://osu.ppy.sh/b/%d", bm.BeatmapID),
			Name:    fmt.Sprintf("%s [%s] [%.2f★]", bm.Title, bm.Version, bm.Difficulty),
			IconURL: fmt.Sprintf("https://a.gatari.pw/%d", id),
		},
		Description: fmt.Sprintf("▸ %.2fPP ▸ %.2f%%\n▸ %d ▸ %dx ▸ [%d/%d/%d/%d]", sc.Pp, sc.Accuracy, sc.Score, sc.MaxCombo, sc.Count300, sc.Count100, sc.Count50, sc.CountMiss),
	}, m.Reference())
}
