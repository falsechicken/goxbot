package plugins

import (
	"github.com/falsechicken/goxbot"
	"github.com/mattn/go-xmpp"
)

type PluginTemplate struct {
	Conf    goxbot.PluginConf
	Name    string
	Version string

	client *xmpp.Client
}

//Initialize the plugin. Providing the configuration table and a pointer to the xmpp client.
func (a *PluginTemplate) Init(c *xmpp.Client, conf map[string]string) bool {
	a.Name = "PluginTemplate"
	a.Version = "0.0.1"
	a.client = c
	return true
}

//GetInfo returns the name and version number of the plugin.
func (a *PluginTemplate) GetInfo() (string, string) {
	return a.Name, a.Version
}

//ProcessChat is called when a message is received by the bot. Return true to prevent the message from being processed by the rest of the plugins.
func (a *PluginTemplate) ProcessChat(m xmpp.Chat) bool {
	return false
}

//ProcessPresence is called when a change in a contacts presence has occured. Return true to prevent the presence from being processed by the rest of the plugins.
func (a *PluginTemplate) ProcessPresence(p xmpp.Presence) bool {
	return false
}

//ProcessCommand is called when a command is issued and the plugin has registered the command. Return true if the command was successful.
func (a *PluginTemplate) ProcessCommand(jid string, cmd string, arg []string) bool {
	return true
}
