package plugins

import (
	"github.com/falsechicken/goxbot"
	"github.com/mattn/go-xmpp"
)

type PluginTemplate struct {
	Conf goxbot.PluginConf
}

//Initialize the plugin. Providing the configuration table.
func (a *PluginTemplate) Init(conf map[string]string) bool {
	return true
}

//ProcessChat is called when the intentEngine has determined that a message has matched a signature of a plugin.
func (a *PluginTemplate) ProcessChat(m xmpp.Chat) bool {
	return false
}

func (a *PluginTemplate) ProcessPresence(p xmpp.Presence) bool {
	return false
}

//ProcessCommand is called when a command is issued and the plugin has registered the command.
func (a *PluginTemplate) ProcessCommand(cmd string, arg string) {

}
