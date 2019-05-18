package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// our main function
func main() {
	d, err := discordgo.New("Bot 579197034234511361")

	if err != nil {
		fmt.Println("failed to create discord session", err)
	}

	bot, err := d.User("@me")

	if err != nil {
		fmt.Println("failed to access account", err)
	}

	d.AddHandler(handleCmd)
	err = d.Open()

	if err != nil {
		fmt.Println("unable to establish connection", err)
	}

	defer d.Close()

	<-make(chan struct{})
}

// our command handler function
func handleCmd(d *discordgo.Session, msg *discordgo.MessageCreate) {
	user := msg.Author
	if user.ID == bid || user.Bot {
		return
	}

	content := msg.Content

	if (content == "!test") {
		d.ChannelMessageSend(msg.ChannelID, "Testing..")
	}

	fmt.Printf("Message: %+v\n", msg.Message)
}