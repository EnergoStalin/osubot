package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	for k, v := range messageHandlerMap {
		if strings.HasPrefix(m.Content, "!"+k) {
			v(s, m)
			break
		}
	}
}
