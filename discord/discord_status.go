package discord

import "sync"

type DiscordStatus struct {
	MU sync.RWMutex `json:"-"`
	ID string
	Username string
	DisplayName string
	Avatar string
	Activity string
	Status string
}

var Status DiscordStatus