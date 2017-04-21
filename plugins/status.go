package plugins

import (
	"github.com/falsechicken/glogger"
	"github.com/falsechicken/goxbot"
	"github.com/falsechicken/goxbot/command"
	"github.com/mattn/go-xmpp"
)

type Status struct {
	Conf    goxbot.PluginConf
	Name    string
	Version string

	client *xmpp.Client
}

func (s *Status) Init(c *xmpp.Client, conf map[string]string) bool {
	s.Name = "Status"
	s.Version = "0.0.1"

	s.client = c

	glogger.LogMessage(glogger.Info, s.Name+" "+s.Version+" initializing...")

	command.Register("status", s)
	return true
}

func (s *Status) GetInfo() (string, string) {
	return s.Name, s.Version
}

func (s *Status) ProcessChat(m xmpp.Chat) bool {
	return false
}

func (s *Status) ProcessPresence(p xmpp.Presence) bool {
	return false
}

func (s *Status) ProcessCommand(cmd string, arg []string) bool {
	return false
}

func handleCmd() {
	glogger.LogMessage(glogger.Info, "Handled status command")
}
