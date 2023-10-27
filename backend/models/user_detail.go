package models

import (
	"time"
	"fmt"
	"github.com/devarshitrivedi01/hero_auth/database"
	"github.com/devarshitrivedi01/hero_auth/utils"
)

type userDetail struct {
	Id             int64     `json:"userId"`
	Username       string    `json:"userName"`
	Password       string    `json:"password"`
	Name           string    `json:"name"`
	Birthdate      time.Time `json:"birthDate"`
	Deathdate      time.Time `json:"deathDate"`
	Originid       int64     `json:"originId"`
	Organizationid int64     `json:"organizationId"`
	Retired        bool      `json:"retired"`
}

type User struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

func GetUserDetail() []userDetail {
	db := database.CreateConnection()
	rows, err := db.Query("Select * from user_detail;")

	var users []userDetail

	if err != nil {
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

func CheckUser(user User) bool {
	db := database.CreateConnection()
	rows, err := db.Query("SELECT * FROM user_detail WHERE username=$1 AND password=$2;", user.Username, user.Password)	
	if err != nil {
		fmt.Println(err)
	}
	count := 0
	
	for rows.Next() {
		count = count + 1
	}
	
		fmt.Println(count)
	if count >= 1 {
		return true
	}

	return false
}

func AddSession(s utils.SessionDetail) error {
	db := database.CreateConnection()
	var uId int
	err := db.QueryRow("select id FROM user_detail WHERE username=$1 AND password=$2;", s.Username, s.Password).Scan(&uId)	
	if err != nil {
		panic(err)
	}
	q := "INSERT INTO session_detail (session_id, user_id, expires) VALUES ($1, $2, $3);"
	insert, err := db.Prepare(q)
	if err != nil {
		
		fmt.Println(err)
	}
	_, err = insert.Exec(s.Session_id, uId, s.Expires)
	insert.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}