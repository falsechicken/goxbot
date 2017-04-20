/*
AutoSubscribe will always accept a request to see the bots status.
*/
package plugins

import (
	"fmt"
	"github.com/falsechicken/goxbot"
	"github.com/mattn/go-xmpp"
)

type AutoSubscribe struct {
	Conf    goxbot.PluginConf
	Name    string
	Version string

	client *xmpp.Client
}

//Initialize the plugin. Providing the configuration table.
func (a *AutoSubscribe) Init(c *xmpp.Client, conf map[string]string) bool {
	a.Name = "AutoSubscribe"
	a.Version = "0.0.1"

	a.client = c

	return true
}

//GetInfo returns the name and version of the plugin.
func (a *AutoSubscribe) GetInfo() (string, string) {
	return a.Name, a.Version
}

/*
 * Process an incoming message. Return true to signify the handling of the message.
 * Prevents the message from propagating further.
 */
func (a *AutoSubscribe) ProcessChat(m xmpp.Chat) bool {
	return false
}

func (a *AutoSubscribe) ProcessPresence(m xmpp.Presence) bool {
	fmt.Println(m.Type)
	if m.Type == "subscribe" {
		a.client.ApproveSubscription(m.From)
	}
	return false
}

func (a *AutoSubscribe) ProcessCommand(cmd string, arg []string) {

}
