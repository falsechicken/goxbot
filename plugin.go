package goxbot

import "github.com/mattn/go-xmpp"

type Plugin interface {
	//Initialize Plugin
	Init(c map[string]string) bool
	ProcessMessage(m xmpp.Chat) bool
}

type PluginConf struct {
	Enabled   bool
	ConfTable map[string]string
}
