package handlers

import "github.com/bwmarrin/discordgo"

type MessageCallback func(s *discordgo.Session, m *discordgo.MessageCreate)

var (
	messageHandlerMap = make(map[string]MessageCallback)
)
