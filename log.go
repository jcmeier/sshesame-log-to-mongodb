package main

import (
	"fmt"
	"net"
	"regexp"
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

type LogLine struct {
	Time          time.Time     `json:"time" bson:"time"`
	Source        string        `json:"source" bson:"source"`
	SourceDetails SourceDetails `json:"source_details" bson: "-"`
	EventType     string        `json:"event_type" "bson:event_type"`
	Event         Event         `json:"event" "bson:event"`
}

type SourceDetails struct {
	Ip         string `json:"ip" bson"ip"`
	Port       string `json:"port" bson"port"`
	ReverseDNS string `json:"reverse_dns" bson"reverse_dns"`
}

var regex *regexp.Regexp = regexp.MustCompile(`((?:\d{1,3}\.){3}\d{1,3}|\[?(?:[0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}\]?)\s*:\s*(\d+)`)

func (l *LogLine) ExtractIpAndPort() {
	matches := regex.FindAllStringSubmatch(l.Source, -1)
	for _, match := range matches {
		addr, err := net.LookupAddr(match[1])
		reverseDNS := `<unknown>`
		if err == nil {
			reverseDNS = addr[0]
		}
		l.SourceDetails = SourceDetails{
			match[1],
			match[2],
			reverseDNS,
		}
	}

	fmt.Println(l.SourceDetails)
}
