package commands

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/EnergoStalin/osubot/internal/services"
	"github.com/EnergoStalin/pixiv"
	"github.com/bwmarrin/discordgo"
)

func fetchAll(id uint64) ([]pixiv.Illust, error) {
	var (
		err     error
		next    int            = 0
		illusts                = []pixiv.Illust{}
		s       []pixiv.Illust = []pixiv.Illust{}
	)

	PIXIV := services.KVStore.Get("PIXIV" + fmt.Sprint(id))
	if PIXIV == "" {
		for {
			s, next, err = services.PixivApp.UserBookmarksIllust(id, "private", next, "")
			illusts = append(illusts, s...)
			if next == 0 {
				break
			}
		}
		text, _ := json.Marshal(illusts)
		services.KVStore.Put("PIXIV"+fmt.Sprint(id), string(text))
	} else {
		json.Unmarshal([]byte(PIXIV), &illusts)
	}

	return illusts, err
}

func Pixiv(s *discordgo.Session, m *discordgo.MessageCreate) {
	id, _ := strconv.Atoi(services.PixivAcc.ID)

	illust, err := fetchAll(uint64(id))
	if err != nil {
		s.ChannelMessageSendReply(m.ChannelID, err.Error(), m.Reference())
		return
	}

	mm := illust[rand.Int()%len(illust)]

	s.ChannelMessageSendReply(m.ChannelID, fmt.Sprintf("https://www.pixiv.net/en/artworks/%d", mm.ID), m.Reference())
}
