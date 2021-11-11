package models

import "time"

type Session struct {
	SessionId   string
	UserId      string
	CurrentTime time.Time
}
