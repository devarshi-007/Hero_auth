package models

import (
	"time"
	"github.com/devarshitrivedi01/hero_auth/database"
)

type userDetail struct {
	Id int64 `json:"userId"`
	Username string `json:"userName"` 
	Password string `json:"password"`
	Name string `json:"name"`
	Birthdate time.Time `json:"birthDate"`
	Deathdate time.Time `json:"deathDate"`
	Originid int64 `json:"originId"`
	Organizationid int64 `json:"organizationId"`
	Retired bool `json:"retired"`
}

func GetUserDetail() []userDetail {
	db := database.CreateConnection()
	rows, err := db.Query("Select * from user_detail;")
	
	var users []userDetail

	if err !=nil {
		panic(err)
	}
	
	for rows.Next() {
		var rec userDetail
		rows.Scan(&rec.Id, &rec.Username, &rec.Password, &rec.Name, &rec.Birthdate, &rec.Deathdate, &rec.Originid, &rec.Organizationid, &rec.Retired)		
		users = append(users, rec)
	}
	
	rows.Close()
	db.Close()

	return users
}