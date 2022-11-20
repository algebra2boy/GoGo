package main

import (
	"flag"

	"github.com/bwmarrin/discordgo"
)

// important variables for command line parameters
var (
	Token string
)

const KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse() // parses the command-line flags from os.Args[1:]
}

func main() {
	// constrct a new Discord client which can be used to access the Discord API
	discord, err := discordgo.New("Bot " + Token)
}
