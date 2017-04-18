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
	Conf goxbot.PluginConf
}

//Initialize the plugin. Providing the configuration table.
func (a *AutoSubscribe) Init(conf map[string]string) bool {
	fmt.Println("INIT AUTO")

	return true
}

/*
 * Process an incoming message. Return true to signify the handling of the message.
 * Prevents the message from propagating further.
 */
func (a *AutoSubscribe) ProcessMessage(m xmpp.Chat) bool {
	return true
}
