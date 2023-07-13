package commands

import (
	"strings"

	"github.com/EnergoStalin/osubot/internal/services"
	"github.com/bwmarrin/discordgo"
)

func Bind(s *discordgo.Session, m *discordgo.MessageCreate) {
	sa := strings.Split(m.Content, " ")
	services.KVStore.Put(m.Author.ID, sa[1])

	s.ChannelMessageSendReply(m.ChannelID, "Successfully binded.", m.Reference())
}
