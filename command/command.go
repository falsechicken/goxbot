package command

import (
	"github.com/falsechicken/glogger"
	"github.com/falsechicken/goxbot"
	"strings"
)

var CommandPrefix = '@'

var commandTable map[string]goxbot.Plugin

func init() {
	commandTable = make(map[string]goxbot.Plugin)
}

/*
ParseCommand Takes a message and breaks it into an array of arguments. Also removes the command prefix.
*/
func Parse(m string) []string {
	s := (strings.Split(m, " "))

	w := []rune(s[0])

	w = append(w[:0], w[0+1:]...)

	s[0] = string(w)
	return s
}

//IsCommand checks if an entry exists in the command table.
func Exists(cmd string) bool {
	if _, exists := commandTable[cmd]; exists {
		return true
	} else {
		return false
	}
}

//SetCommandPrefix sets the character needed to decide to check a message as a command.
func SetCommandPrefix(p rune) {
	CommandPrefix = p
}

//Register register a plugin to act upon a command.
func Register(cmd string, plugin goxbot.Plugin) {
	glogger.LogMessage(glogger.Debug, "Registered command "+cmd)
	commandTable[cmd] = plugin
}

//Execute runs a command. Accepts the command to be run a slice of arguments.
func Execute(cmd string, args []string) bool {
	commandTable[cmd].ProcessCommand(cmd, args)
	return true
}

//HasCommandPrefix returns true is the strings first element is the command prefix.
func HasCommandPrefix(cmd string) bool {
	s := strings.Split(cmd, "")
	if s[0] == string(CommandPrefix) {
		return true
	} else {
		return false
	}
}
