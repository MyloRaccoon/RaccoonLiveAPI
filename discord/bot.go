package discord

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func Run() *discordgo.Session {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil
	}
	token := os.Getenv("DISCORD_TOKEN")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord sessions: ", err)
		return nil
	}
	println("bot created")

	dg.AddHandler(messageCreate)
	println("added messageCreate event")
	dg.AddHandler(presenceUpdate)
	println("added presenceUpdate event")
	dg.Identify.Intents = discordgo.IntentsGuildMembers | 
				discordgo.IntentsGuildPresences |
				discordgo.IntentsGuildMessages |
				discordgo.IntentsMessageContent
	println("added intents")

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return nil
	}
	println("bot connected")

	return dg
}

func presenceUpdate(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	println("Presence updated")
	activity := ""
	if len(p.Activities) > 0 {
		activity = p.Activities[0].Name
	}
	username := p.User.Username
	avatar := p.User.Avatar
	updateStatus(username, avatar, activity, "")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("Message created!")
	fmt.Printf("[%s] %s ", m.Author.Username, m.Content)

	if m.Author.ID == s.State.User.ID {
		fmt.Println("Self message: abort")
		return
	}

	msg := ""
	if m.Content != "&remove" {
		user := m.Author
		updateStatus(user.Username, user.Avatar, "", m.Content)
		msg = fmt.Sprintf("Status updated to: '%s'", m.Content)
	} else {
		msg = "Status removed"
	}
	s.ChannelMessageSend(m.ChannelID, msg)
}

func updateStatus(username string, avatar string, activity string, status string) {
	Status.MU.Lock()
	defer Status.MU.Unlock()
	if username != "" {
		Status.Username = username
	}
	if avatar != "" {
		Status.Avatar = avatar
	}
	if activity != "" {
		Status.Activity = activity
	}
	if status != "" {
		Status.Status = status
	} else if status == "&remove" {
		Status.Status = ""
	}
	println("status updated")
}