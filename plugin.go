package goxbot

import "github.com/mattn/go-xmpp"

type Plugin interface {
	//Init initializes the plugin.Provides a pointer to the client and  returns true if successful.
	Init(*xmpp.Client, map[string]string) bool

	//GetInfo should return the name of the plugin followed by the version.
	GetInfo() (string, string)

	//ProcessChat is called when a message matches a plugins signature.
	ProcessChat(m xmpp.Chat) bool

	//ProcessPresence is called when a presense message is received.
	ProcessPresence(p xmpp.Presence) bool

	//ProcessCommand is called when a command is made that the plugin registered. Commands are any message starting with the command prefix. @ by default.
	ProcessCommand(c string, a []string) bool
}

type PluginConf struct {
	Enabled   bool
	ConfTable map[string]string
}
