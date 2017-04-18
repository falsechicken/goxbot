package command

import (
	"github.com/falsechicken/goxbot"
)

var CommandTable map[string]*goxbot.Plugin
var CommandPrefix = '@'

func init() {
	CommandTable = make(map[string]*goxbot.Plugin)
}

//IsCommand returns true if the message starts with the command prefix.
func IsCommand(m string) bool {
	if []rune(m)[0] == CommandPrefix {
		return true
	} else {
		return false
	}
}

//SetCommandPrefix sets the character that will be used to run a message as a command or not.
func SetCommandPrefix(p rune) {
	CommandPrefix = p
}

/*
AddCommandToTable adds a command and its matching plugin to the command table.
Previous entries will be overwritten.
*/
func AddCommandToTable(command string, p *goxbot.Plugin) {
	CommandTable[command] = p
}

//parseCommand Takes a message and breaks it into an array of arguments.
func parseCommand(m string) []string {
	return []string{m}
}
