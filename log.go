package main

import (
	"time"
)

type Event struct {
	ChannelID     int    `json:"channel_id" bson:"channel_id"`
	Command       string `json:"command" bson:"command,omitempty"`
	User          string `json:"user" bson:"user,omitempty"`
	Password      string `json:"password" bson:"password,omitempty"`
	Accepted      bool   `json:"accepted" bson:"accepted,omitempty"`
	ClientVersion string `json:"client_version" bson:"client_version,omitempty"`
	Terminal      string `json:"terminal" bson:"terminal,omitempty"`
	Width         int    `json:"width" bson:"width,omitempty"`
	Height        int    `json:"height" bson:"height,omitempty"`
	Name          string `name:"client_version" bson:"name,omitempty"`
	Value         string `name:"value" bson:"value,omitempty"`
	Input         string `name:"input" bson:"input,omitempty"`
}

// MyClass represents the main class with the specified attributes
type LogLine struct {
	Time      time.Time `json:"time" bson:"time"`
	Source    string    `json:"source" bson:"source"`
	EventType string    `json:"event_type" "bson:event_type"`
	Event     Event     `json:"event" "bson:event"`
}
