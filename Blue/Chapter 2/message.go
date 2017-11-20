package main

import (
	"time"
)

type message struct {
	When         time.Time
	Name         string
	Message      string
	AvatarURL    string
}
