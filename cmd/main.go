package main

import "github.com/Smirnov-O/noter/pkg/database"

func main() {
	_, err := database.NewConnection(database.ConnInfo{
		Host:     "",
		Port:     "",
		Username: "",
		Password: "",
		DBName:   "",
		SSLMode:  "",
	})
	if err != nil {
		panic(err)
	}
}
