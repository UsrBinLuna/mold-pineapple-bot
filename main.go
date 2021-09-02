package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
	gid   string
	/*chname      []string
	nchname     string*/
	ChannelType string
)

// channel types
const (
	ChannelTypeGuildText discordgo.ChannelType = iota
	ChannelTypeDM
	ChannelTypeGuildVoice
	ChannelTypeGroupDM
	ChannelTypeGuildCategory
	ChannelTypeGuildNews
	ChannelTypeGuildStore
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called every time a new message is created on any channel
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	var msg string = m.Content

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "m!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	} else if m.Content == "m!pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	} else if m.Content == "m!cc" { // create channel thing
		// not working

		gid = m.GuildID
		//chname := m.Content
		s.GuildChannelCreate(gid, strings.Split(msg, " ")[1], ChannelTypeGuildText)
		s.ChannelMessageSend(m.ChannelID, "Created channel #"+strings.Fields(msg)[1])
	}
}
