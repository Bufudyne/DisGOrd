package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

var s *discordgo.Session

func (app *application) startBot() error {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + app.config.bot.token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return err
	}
	s = dg
	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(app.messageCreate)

	// Just like the ping pong example, we only care about receiving message
	// events in this example.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return err
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	err = dg.Close()
	if err != nil {
		return err
	}
	return nil
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
//
// It is called whenever a message is created but only when it's sent through a
// server as we did not request IntentsDirectMessages.
func (app *application) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	var response WsJsonResponse

	response.Action = "user_message"
	m.Author.Avatar = m.Author.AvatarURL("128")
	response.ChatMessage = *m
	app.broadcastToAll(response)
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	fmt.Printf("%+v\n", m.Author.Username)
	fmt.Printf("%+v\n", m.Author.Avatar)
	fmt.Printf("%+v\n", m.Timestamp)
	fmt.Printf("%+v\n", m.EditedTimestamp)
	fmt.Printf("%+v\n", m.Content)

	// In this example, we only care about messages that are "ping".
	if m.Content != "ping" {
		return
	}

	// We create the private channel with the user who sent the message.
	channel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		// If an error occurred, we failed to create the channel.
		//
		// Some common causes are:
		// 1. We don't share a server with the user (not possible here).
		// 2. We opened enough DM channels quickly enough for Discord to
		//    label us as abusing the endpoint, blocking us from opening
		//    new ones.
		fmt.Println("error creating channel:", err)
		_, err := s.ChannelMessageSend(
			m.ChannelID,
			"Something went wrong while sending the DM!",
		)
		if err != nil {
			return
		}
		return
	}
	// Then we send the message through the channel we created.
	_, err = s.ChannelMessageSend(channel.ID, "Pong!")
	if err != nil {
		// If an error occurred, we failed to send the message.
		//
		// It may occur either when we do not share a server with the
		// user (highly unlikely as we just received a message) or
		// the user disabled DM in their settings (more likely).
		fmt.Println("error sending DM message:", err)
		_, err := s.ChannelMessageSend(
			m.ChannelID,
			"Failed to send you a DM. "+
				"Did you disable DM in your privacy settings?",
		)
		if err != nil {
			return
		}
	}
}
func (app *application) sendMessage(m string) {
	// We create the private channel with the user who sent the message.
	channel, err := s.Channel("602892529997840399")
	if err != nil {
		// If an error occurred, we failed to create the channel.
		//
		// Some common causes are:
		// 1. We don't share a server with the user (not possible here).
		// 2. We opened enough DM channels quickly enough for Discord to
		//    label us as abusing the endpoint, blocking us from opening
		//    new ones.
		fmt.Println("error creating channel:", err)
		_, err := s.ChannelMessageSend(
			m,
			"Something went wrong while sending the DM!",
		)
		if err != nil {
			return
		}
		return
	}
	// Then we send the message through the channel we created.
	_, err = s.ChannelMessageSend(channel.ID, m)
	if err != nil {
		// If an error occurred, we failed to send the message.
		//
		// It may occur either when we do not share a server with the
		// user (highly unlikely as we just received a message) or
		// the user disabled DM in their settings (more likely).
		fmt.Println("error sending DM message:", err)
		_, err := s.ChannelMessageSend(
			m,
			"Failed to send you a DM. "+
				"Did you disable DM in your privacy settings?",
		)
		if err != nil {
			return
		}
	}
}
