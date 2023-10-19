package database

import "database/sql"

func CreateConnection() *sql.DB {
	connectionStr := "postgres://devarshi:eA234123Pq@localhost:5432/supreme_org?sslmode=disable"
	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}
	return conn
}