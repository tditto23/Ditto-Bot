package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var (
	commandPrefix string
	botID         string
)

func main() {
	discord, err := discordgo.New("Bot NTc5MjA1NjU3OTIwNTM2NTc3.XN-xWg.uL2-1PVqTSs9cvWqaNUvPy9OaOA")
	errCheck("error creating discord session", err)
	user, err := discord.User("@me")
	errCheck("error retrieving account", err)

	botID = user.ID
	discord.AddHandler(commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err = discord.UpdateStatus(0, "A friendly helpful bot!")
		if err != nil {
			fmt.Println("Error attempting to set my status")
		}
		servers := discord.State.Guilds
		fmt.Printf("SuperAwesomeOmegaTutorBot has started on %d servers", len(servers))
	})

	err = discord.Open()
	errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	commandPrefix = "D!"

	<-make(chan struct{})

}

func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botID || user.Bot {
		//Do nothing because the bot is talking
		return
	}

	content := message.Content
	if len(content) <= len(commandPrefix) {
		return
	}
	if content[:len(commandPrefix)] != commandPrefix {
		return
	}
	content = content[len(commandPrefix)+1:]
	if len(content) < 1 {
		return
	}
	args := strings.Fields(content)
	name := strings.ToLower(args[0])
	if name == "rpg" {
		discord.ChannelMessageSend(message.ChannelID, "rpg hunt")
	}
	if name == "test" {
		discord.ChannelMessageSend(message.ChannelID, "Testing..")

	}

	fmt.Printf("Message: %+v || From: %s\n", message.Message, message.Author, content)
}
