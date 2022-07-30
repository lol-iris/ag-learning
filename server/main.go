package main

import (
	"encoding/json"
	"log"

	"github.com/lol-iris/aglearning/Config"
	"github.com/lol-iris/aglearning/Models"
)

var err error

func main() {
	// Initialize DB
	Config.Init()

	user := Models.User{
		Username: "admin",
		Password: "pwd",
		Email:    "admin@mail.org",
		IsAdmin:  true,
	}
	Config.DB.FirstOrCreate(&user)
	output, _ := json.MarshalIndent(user, "", "  ")
	log.Println(string(output))
}
