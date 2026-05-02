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
	id := p.User.ID
	username := p.User.Username
	displayName := p.User.DisplayName()
	avatar := p.User.Avatar
	updateStatus(id, username, displayName, avatar, activity, "")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	msg := ""
	if m.Content != "&remove" {
		msg = fmt.Sprintf("Status updated to: '%s'", m.Content)
	} else {
		msg = "Status removed"
	}
	user := m.Author
	updateStatus(user.ID, user.Username, user.DisplayName(), user.Avatar, "&dontupdate", m.Content)
	s.ChannelMessageSend(m.ChannelID, msg)
}

func updateStatus(id string, username string, displayName string, avatar string, activity string, status string) {
	Status.MU.Lock()
	defer Status.MU.Unlock()
	if id != "" {
		Status.ID = id
	}
	if username != "" {
		Status.Username = username
	}
	if displayName != "" {
		Status.DisplayName = displayName
	}
	if avatar != "" {
		Status.Avatar = fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", id, avatar)
	}
	if activity != "&dontupdate" {
		Status.Activity = activity
	}
	if status == "&remove" {
		Status.Status = ""
	} else if status != "" {
		Status.Status = status
	} 
}