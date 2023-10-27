package utils

import (
	"fmt"
	"time"

	"github.com/devarshitrivedi01/hero_auth/database"
)

func Valid(c string) bool {
	var expires time.Time
	db := database.CreateConnection()
	err := db.QueryRow("Select expires from session_detail where session_id=$1;", c).Scan(&expires)
	if err != nil {
		fmt.Println(err)
	}
	if expires.Before(time.Now()) {
		return false
	} else {
		return true
	}
}