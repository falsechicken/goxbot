package command

import (
	"errors"
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

//Register registers a plugin to act upon a command. Only one plugin can subscribe to a command at a time.
func Register(cmd string, plugin goxbot.Plugin) {
	var pName, pVersion = plugin.GetInfo()
	glogger.LogMessage(glogger.Debug, "Plugin "+pName+"(v"+pVersion+") registered command "+cmd+".")

	if _, exists := commandTable[cmd]; exists {
		var cPName, _ = commandTable[cmd].GetInfo()
		glogger.LogMessage(glogger.Warning, "Plugin "+pName+" is overwriting command "+cmd+" registered by "+cPName)
	}

	commandTable[cmd] = plugin
}

//Execute runs a command. Accepts the command to be run and a slice of arguments.
func Execute(cmd string, args []string) (bool, error) {
	if !Exists(cmd) {
		glogger.LogMessage(glogger.Debug, "Command "+cmd+" does not exist.")
		return false, errors.New("Command " + cmd + " does not exist.")
	} else {
		glogger.LogMessage(glogger.Debug, "Executing command "+cmd)
		commandTable[cmd].ProcessCommand(cmd, args)
		return true, nil
	}
}

//HasCommandPrefix returns true is the strings first element is the command prefix.
func HasCommandPrefix(cmd string) bool {
	s := strings.Split(cmd, "")
	if len(s) != 0 && s[0] == string(CommandPrefix) {
		return true
	} else {
		return false
	}
}
