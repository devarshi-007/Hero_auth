package utils

import (
	"time"
)


type SessionDetail struct {
	Session_id string
	Username string
	Password string
	Expires time.Time
}