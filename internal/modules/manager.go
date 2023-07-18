package modules

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type ModuleManager struct {
	modules   []BaseModule
	callbacks InteractionCallbacks
	logger    *logrus.Logger
}

func NewModuleManager() *ModuleManager {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	return &ModuleManager{
		modules:   []BaseModule{},
		callbacks: InteractionCallbacks{},
		logger:    log,
	}
}

func (m *ModuleManager) log() *logrus.Entry {
	return m.logger.WithTime(time.Now()).WithField("prefix", "ModuleManager")
}

func (m *ModuleManager) registerCommands(b BaseModule) {
	for k, v := range b.GetCallbacks() {
		m.callbacks[k] = v
	}
}

func (m *ModuleManager) RegiesterModule(b BaseModule) (err error) {
	m.modules = append(m.modules, b)
	err = b.Init()
	if err != nil {
		return
	}
	m.registerCommands(b)

	return
}

func (m *ModuleManager) RegisterCommands(s *discordgo.Session) {
	for _, v := range m.modules {
		m.log().Debug(fmt.Sprintf("Registering commands from module [%s]...", v.GetName()))
		mc := v.GetCommands()
		for _, c := range mc {
			_, err := s.ApplicationCommandCreate(s.State.User.ID, "", (*discordgo.ApplicationCommand)(c))
			if err != nil {
				m.log().Warn(fmt.Sprintf("Registering command [%s] failed for module [%s]...", c.Name, v.GetName()))
			}
		}
	}
}

func (m *ModuleManager) UnRegisterAllCommands(s *discordgo.Session) {
	m.log().Debug("Unregistering commands...")

	c, _ := s.ApplicationCommands(s.State.User.ID, "")
	for _, v := range c {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			m.log().Warn(fmt.Sprintf("Unregistering command [%s] failed...", v.Name))
		}
		m.log().Debug(fmt.Sprintf("Unregistered [%s]...", v.Name))
	}
}

func (m *ModuleManager) Invoke(s *discordgo.Session, i *discordgo.InteractionCreate) {
	name := i.ApplicationCommandData().Name
	if v, ok := m.callbacks[name]; ok {
		v(s, i)
	} else {
		m.log().Warn(fmt.Sprintf("Command [%s] not found...", name))
	}
}
