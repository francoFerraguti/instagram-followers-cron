package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/liteByte/frango"
)

var db *sql.DB

func ConnectToDatabase() {
	var err error

	db, err = sql.Open(GetConfig().DB_TYPE, GetConfig().DB_USERNAME+":"+GetConfig().DB_PASSWORD+"@tcp("+GetConfig().DB_HOST+":"+GetConfig().DB_PORT+")/"+GetConfig().DB_NAME)
	frango.PrintErr(err)

	err = db.Ping()
	frango.PrintErr(err)

	createSchema()
}

func GetDatabase() *sql.DB {
	return db
}

func createSchema() {
	createUserTable()
	mockUsers()
}

func createUserTable() {
	query := `
		CREATE TABLE IF NOT EXISTS Users (
			ID int UNIQUE NOT NULL,
			Username varchar(255) UNIQUE NOT NULL,
			FullName varchar(255),
			InstagramID int UNIQUE NOT NULL,
			IsFollower boolean NOT NULL,
			ImFollowing boolean NOT NULL,
			FollowDate date,
			UnfollowDate date
		);`

	_, err := db.Exec(query)
	frango.PrintErr(err)
}

func mockUsers() {

}
