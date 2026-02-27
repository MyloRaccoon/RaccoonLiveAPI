package discord

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func Run() *discordgo.Session {
	println("Connecting bot...")
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

	dg.AddHandler(messageCreate)
	dg.AddHandler(presenceUpdate)
	dg.Identify.Intents = discordgo.IntentsGuildMembers | 
				discordgo.IntentsGuildPresences |
				discordgo.IntentsGuildMessages |
				discordgo.IntentsMessageContent

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return nil
	}
	println("bot connected!")

	return dg
}

func presenceUpdate(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	activity := ""
	if len(p.Activities) > 0 {
		activity = p.Activities[0].Name
	}
	username := p.User.Username
	avatar := p.User.Avatar
	updateStatus(username, avatar, activity, "")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
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
}