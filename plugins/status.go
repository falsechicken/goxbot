package plugins

import (
	"os/exec"

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

func (s *Status) ProcessCommand(jid string, cmd string, arg []string) bool {
	if cmd == "status" {
		handleCmd(s, jid)
		return true
	}
	return false
}

func handleCmd(s *Status, jid string) {

	memoryInfo, _ := exec.Command("free", "-m").Output()

	var msg = string(memoryInfo)

	s.client.Send(xmpp.Chat{Remote: jid, Type: "chat", Text: msg})

}
