package modules

import (
	"reflect"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type InteractionCallback func(*discordgo.Session, *discordgo.InteractionCreate)
type InteractionCallbacks map[string]InteractionCallback
type ModuleCommand discordgo.ApplicationCommand
type ModuleCommands []*ModuleCommand

type BaseModule interface {
	Init() error
	GetName() string
	GetCallbacks() InteractionCallbacks
	GetCommands() ModuleCommands
}

func ReflectCallbacks[B any](d B) InteractionCallbacks {
	c := InteractionCallbacks{}

	t := reflect.TypeOf(d)
	v := reflect.ValueOf(d)
	for i := 0; i < v.NumMethod(); i++ {
		mt := t.Method(i)
		cb, ok := mt.Func.Interface().(func(B, *discordgo.Session, *discordgo.InteractionCreate))
		if !ok {
			continue
		}
		c[strings.ToLower(mt.Name)] = func(s *discordgo.Session, ic *discordgo.InteractionCreate) {
			cb(d, s, ic)
		}
	}

	return c
}
