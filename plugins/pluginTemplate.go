package plugins

import (
	"github.com/mattn/go-xmpp"
	"github.com/falsechicken/goxbot"
)

type PluginTemplate struct {
	Conf goxbot.PluginConf
}

//Initialize the plugin. Providing the configuration table.
func (a *PluginTemplate) Init(conf map[string]string) bool {
	return true
}

/*
 * Process an incoming message. Return true to signify the handling of the message.
 * Prevents the message from propagating further.
 */
func (a *PluginTemplate) ProcessMessage(m xmpp.Chat) bool {
	return true
}
