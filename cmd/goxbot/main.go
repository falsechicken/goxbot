// File: main.go

package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/falsechicken/glogger"
	"github.com/falsechicken/goxbot"
	"github.com/falsechicken/goxbot/command"
	"github.com/falsechicken/goxbot/config"
	"github.com/falsechicken/goxbot/permissions"
	"github.com/falsechicken/goxbot/plugins"
	"github.com/mattn/go-xmpp"
)

//Goxbot Version
const Version = "0.0.1"

var server = flag.String("server", "", "server:port")
var username = flag.String("username", "", "username")
var password = flag.String("password", "", "password")
var status = flag.String("status", "", "status")
var statusMessage = flag.String("status-msg", "I for one welcome our new codebot overlords.", "status message")
var notls = flag.Bool("notls", true, "No TLS")
var starttls = flag.Bool("starttls", true, "Enable StartTLS")
var debug = flag.Bool("debug", false, "debug output")
var session = flag.Bool("session", false, "use server session")
var console = flag.Bool("console", false, "enable the command console.")
var configPath = flag.String("config", ".", "set configuration file location.")

var loadedPlugins = [4]goxbot.Plugin{new(plugins.AutoSubscribe), new(plugins.Status), new(plugins.PluginTemplate), new(plugins.Execute)}

var talk *xmpp.Client

var conf config.Config

func main() {

	glogger.LogMessage(glogger.Info, "GoXBot "+Version+" starting up...")

	parseFlags()

	loadConfig()

	loadPerms()

	initXMPP()

	initPlugins()

	go listen()

	if !*console {
		for {
		}
	} else {
		for {
			in := bufio.NewReader(os.Stdin)
			line, err := in.ReadString('\n')
			if err != nil {
				continue
			}
			line = strings.TrimRight(line, "\n")

			tokens := strings.SplitN(line, " ", 2)
			if len(tokens) == 2 {
				talk.Send(xmpp.Chat{Remote: tokens[0], Type: "chat", Text: tokens[1]})
			}
		}
	}
}

//Run the plugin's init function
func initPlugins() {
	for _, p := range loadedPlugins {
		p.Init(talk, make(map[string]string))
		var name, version = p.GetInfo()
		glogger.LogMessage(glogger.Info, name+" "+version+" initialized.")
	}
}

func processChat(c xmpp.Chat) {
	for _, p := range loadedPlugins {
		if p.ProcessChat(c) {
			break
		}
	}
}

func processPresence(p xmpp.Presence) {
	for _, v := range loadedPlugins {
		if v.ProcessPresence(p) {
			break
		}
	}
}

func parseFlags() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: example [options]\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()
}

func initXMPP() {
	if *username == "" || *password == "" {
		if *debug && *username == "" && *password == "" {
			fmt.Fprintf(os.Stderr, "no username or password were given; attempting ANONYMOUS auth\n")
		} else if *username != "" || *password != "" {
			flag.Usage()
		}
	}

	if !*notls {
		xmpp.DefaultConfig = tls.Config{
			ServerName:         serverName(conf.Server),
			InsecureSkipVerify: false,
		}
	}

	xmpp.DefaultConfig = tls.Config{
		ServerName:         serverName(conf.Server),
		InsecureSkipVerify: true,
	}

	var err error
	options := xmpp.Options{
		Host:                         conf.Server,
		User:                         conf.Username,
		Password:                     conf.Password,
		NoTLS:                        *notls,
		Debug:                        conf.Debug,
		Session:                      conf.Session,
		Status:                       conf.Status,
		StatusMessage:                conf.StatusMessage,
		InsecureAllowUnencryptedAuth: false,
		StartTLS:                     conf.StartTLS,
	}

	talk, err = options.NewClient()

	if err != nil {
		log.Fatal(err)
	}
}

func listen() {
	for {
		chat, err := talk.Recv()
		if err != nil {
			log.Fatal(err)
		}
		switch v := chat.(type) {
		case xmpp.Chat:
			if command.HasCommandPrefix(v.Text) {
				var s = command.Parse(v.Text)
				if permissions.HasPermission(v.Remote, s[0]) {
					go command.Execute(v.Remote, s[0], s[1:])
				} else {
					talk.Send(xmpp.Chat{Remote: v.Remote, Type: "chat", Text: "You do not have permission to execute this command."})
				}
			} else {
				processChat(v)
			}

		case xmpp.Presence:
			processPresence(v)
		}
	}
}

func loadConfig() {
	conf = config.Load(*configPath)
	glogger.SetDebugMode(conf.Debug)
}

func loadPerms() {
	permissions.Load(*configPath)
}

//Seperate domain name and port.
func serverName(host string) string {
	return strings.Split(host, ":")[0]
}
