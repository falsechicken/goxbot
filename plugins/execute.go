package plugins

import (
	"fmt"
	"os/exec"

	"github.com/falsechicken/glogger"
	"github.com/falsechicken/goxbot"
	"github.com/falsechicken/goxbot/command"
	"github.com/mattn/go-xmpp"
)

const helpMessage = "\nUsage: @exec [command] <arguments>\n Ex: @exec touch test"

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

	command.Register("exec", a)
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

func (a *Execute) ProcessCommand(jid string, cmd string, arg []string) bool {
	if cmd != "" && len(arg) > 0 {
		glogger.LogMessage(glogger.Debug, a.Name+" | User "+jid+" running command "+arg[0])
		_, msg := executeCommand(arg[0], arg[1:])
		a.client.Send(xmpp.Chat{Remote: jid, Type: "chat", Text: msg})
	} else {
		a.client.Send(xmpp.Chat{Remote: jid, Type: "chat", Text: helpMessage})
	}
	return true
}

func executeCommand(cmd string, arg []string) (bool, string) {

	out, err := exec.Command(cmd, arg...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	s := string(out)

	return true, s

}
