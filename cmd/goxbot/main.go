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
	"github.com/falsechicken/goxbot/plugins"
	"github.com/mattn/go-xmpp"
)

var server = flag.String("server", "", "server:port")
var username = flag.String("username", "", "username")
var password = flag.String("password", "", "password")
var status = flag.String("status", "xa", "status")
var statusMessage = flag.String("status-msg", "I for one welcome our new codebot overlords.", "status message")
var notls = flag.Bool("notls", true, "No TLS")
var starttls = flag.Bool("starttls", true, "Enable StartTLS")
var debug = flag.Bool("debug", false, "debug output")
var session = flag.Bool("session", false, "use server session")

var loadedPlugins = [1]goxbot.Plugin{new(plugins.AutoSubscribe)}

var talk *xmpp.Client

func main() {

	glogger.LogMessage(glogger.Debug, "GoXBot starting up...")

	parseFlags()

	initXMPP()

	initPlugins()

	go func() {
		for {
			chat, err := talk.Recv()
			if err != nil {
				log.Fatal(err)
			}
			switch v := chat.(type) {
			case xmpp.Chat:
                fmt.Println(v.Remote, v.Text)
				fmt.Println()
			case xmpp.Presence:
				fmt.Println(v.From, v.Show)
				fmt.Println()
			}
		}
	}()

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

//Run the plugin's init function
func initPlugins() {
	for _, v := range loadedPlugins {
		v.Init(make(map[string]string))
	}
}

func loadConfig() {}

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
			ServerName:         serverName(*server),
			InsecureSkipVerify: false,
		}
	}

	xmpp.DefaultConfig = tls.Config{
		ServerName:         serverName(*server),
		InsecureSkipVerify: true,
	}

	var err error
	options := xmpp.Options{
		Host:                         *server,
		User:                         *username,
		Password:                     *password,
		NoTLS:                        *notls,
		Debug:                        *debug,
		Session:                      *session,
		Status:                       *status,
		StatusMessage:                *statusMessage,
		InsecureAllowUnencryptedAuth: false,
		StartTLS:                     *starttls,
	}

	talk, err = options.NewClient()

	if err != nil {
		log.Fatal(err)
	}
}

//Seperate domain name and port.
func serverName(host string) string {
	return strings.Split(host, ":")[0]
}
