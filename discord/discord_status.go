package discord

import "sync"

type DiscordStatus struct {
	MU sync.RWMutex `json:"-"`
	Username string
	DisplayName string
	Avatar string
	Activity string
	Status string
}

var Status DiscordStatus