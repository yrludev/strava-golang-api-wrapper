package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yrludev/strava-golang-api-wrapper/strava"
	"golang.org/x/oauth2"
)

func main() {
	accessToken := os.Getenv("STRAVA_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("STRAVA_ACCESS_TOKEN environment variable not set")
	}
	token := &oauth2.Token{AccessToken: accessToken}
	client := strava.NewClient(token)
	athlete, err := client.GetAthlete()
	if err != nil {
		log.Fatalf("Error getting athlete: %v", err)
	}
	fmt.Printf("Hello, %s %s! (ID: %d)\n", athlete.FirstName, athlete.LastName, athlete.ID)
}
