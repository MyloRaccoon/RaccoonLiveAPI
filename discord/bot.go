package discord

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func Run() (*discordgo.Session, error) {
	println("Connecting bot...")
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	token := os.Getenv("DISCORD_TOKEN")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil,err
	}

	dg.AddHandler(messageCreate)
	dg.AddHandler(presenceUpdate)
	dg.Identify.Intents = discordgo.IntentsGuildMembers | 
				discordgo.IntentsGuildPresences |
				discordgo.IntentsGuildMessages |
				discordgo.IntentsMessageContent

	err = dg.Open()
	if err != nil {
		return nil, err
	}
	println("bot connected!")

	return dg, nil
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
	if m.Author.ID == s.State.User.ID {
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