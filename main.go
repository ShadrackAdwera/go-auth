package main

import (
	"log"

	"github.com/ShadrackAdwera/go-auth/api"
	"github.com/ShadrackAdwera/go-auth/authenticator"
	"github.com/joho/godotenv"
)

// vocek57061@oniecan.com
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	srv := api.NewServer(auth)

	err = srv.StartServer("0.0.0.0:3000")

	if err != nil {
		panic(err)
	}
}
