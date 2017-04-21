package plugins

import (
	"github.com/falsechicken/goxbot"
	"github.com/mattn/go-xmpp"
)

type Execute struct {
	Conf    goxbot.PluginConf
	Name    string
	Version string

	client *xmpp.Client
}

func (a *Execute) Init(c *xmpp.Client, conf map[string]string) bool {
	a.Name = "Execute"
	a.Version = "0.0.1"
	a.client = c
	return true
}

func (a *Execute) GetInfo() (string, string) {
	return a.Name, a.Version
}

func (a *Execute) ProcessChat(m xmpp.Chat) bool {
	return false
}

func (a *Execute) ProcessPresence(p xmpp.Presence) bool {
	return false
}

func (a *Execute) ProcessCommand(cmd string, arg string) bool {
	return true
}
