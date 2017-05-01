/*
AutoSubscribe will always accept a request to see the bots status.
*/
package plugins

import (
	"github.com/falsechicken/glogger"
	"github.com/falsechicken/goxbot"
	"github.com/mattn/go-xmpp"
)

type AutoSubscribe struct {
	Conf    goxbot.PluginConf
	Name    string
	Version string

	client *xmpp.Client
}

func (a *AutoSubscribe) Init(c *xmpp.Client, conf map[string]string) bool {
	a.Name = "AutoSubscribe"
	a.Version = "0.0.1"

	a.client = c

	return true
}

func (a *AutoSubscribe) GetInfo() (string, string) {
	return a.Name, a.Version
}

func (a *AutoSubscribe) ProcessChat(m xmpp.Chat) bool {
	return false
}

func (a *AutoSubscribe) ProcessPresence(m xmpp.Presence) bool {
	if m.Type == "subscribe" {
		glogger.LogMessage(glogger.Debug, a.Name+": accepted subscription request from "+m.From)
		a.client.ApproveSubscription(m.From)
	}
	return false
}

func (a *AutoSubscribe) ProcessCommand(jid string, cmd string, arg []string) bool {
	return false
}
