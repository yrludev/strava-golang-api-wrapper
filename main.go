package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yrludev/strava-golang-api-wrapper/strava"
	"golang.org/x/oauth2"
)

func main() {
	createActivity := flag.Bool("create-activity", false, "Create a new activity")
	activityName := flag.String("activity-name", "", "Activity name")
	activityType := flag.String("activity-type", "", "Activity type (e.g., Ride)")
	sportType := flag.String("sport-type", "", "Sport type (e.g., Cycling)")
	startDateLocal := flag.String("start-date-local", "", "Start date local (e.g., 2025-08-12T07:00:00Z)")
	elapsedTime := flag.Int("elapsed-time", 0, "Elapsed time in seconds")
	description := flag.String("description", "", "Description")
	distance := flag.Float64("distance", 0, "Distance in meters")
	trainer := flag.Int("trainer", 0, "Trainer (0 or 1)")
	commute := flag.Int("commute", 0, "Commute (0 or 1)")
	var (
		fetchAthlete    = flag.Bool("athlete", false, "Fetch athlete profile")
		fetchStats      = flag.Bool("stats", false, "Fetch athlete stats")
		fetchActivities = flag.Bool("activities", false, "Fetch recent activities")
		fetchSegments   = flag.Bool("segments", false, "Fetch segments in SF area")
		fetchClub       = flag.Int64("club", 0, "Fetch club by ID (0 to skip)")
		fetchGear       = flag.Bool("gear", false, "Fetch first bike's gear details")
		athleteID       = flag.Int64("athlete-id", 0, "Fetch public info for this athlete ID (if allowed)")
		outputJSON      = flag.Bool("json", true, "Output as JSON (default true)")
	)
	flag.Parse()

	accessToken := os.Getenv("STRAVA_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("STRAVA_ACCESS_TOKEN environment variable not set")
	}
	token := &oauth2.Token{AccessToken: accessToken}
	client := strava.NewClient(token)

	result := make(map[string]interface{})

	if *createActivity {
		if *activityName == "" || *activityType == "" || *sportType == "" || *startDateLocal == "" || *elapsedTime == 0 {
			log.Fatal("Missing required fields for activity creation. Required: --activity-name, --activity-type, --sport-type, --start-date-local, --elapsed-time")
		}
		activity, err := client.CreateActivity(
			*activityName,
			*activityType,
			*sportType,
			*startDateLocal,
			*elapsedTime,
			*description,
			*distance,
			*trainer,
			*commute,
		)
		if err != nil {
			log.Fatalf("Error creating activity: %v", err)
		}
		if *outputJSON {
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			enc.Encode(activity)
		} else {
			fmt.Printf("Created activity: %+v\n", activity)
		}
		return
	}

	if *athleteID != 0 {
		// Fetch public info for a specific athlete ID
		athlete, err := client.GetAthleteByID(*athleteID)
		if err != nil {
			log.Fatalf("Error getting athlete by ID: %v", err)
		}
		result["athlete"] = athlete
	} else if *fetchAthlete || *fetchStats || *fetchGear || *fetchClub != 0 {
		athlete, err := client.GetAthlete()
		if err != nil {
			log.Fatalf("Error getting athlete: %v", err)
		}
		if *fetchAthlete {
			result["athlete"] = athlete
		}
		if *fetchStats {
			stats, err := client.GetAthleteStats(athlete.ID)
			if err == nil {
				result["stats"] = stats
			}
		}
		if *fetchGear && len(athlete.Bikes) > 0 {
			gear, err := client.GetDetailedGear(athlete.Bikes[0].ID)
			if err == nil {
				result["gear"] = gear
			}
		}
	}

	if *fetchActivities {
		activities, err := client.ListAthleteActivities(0, 0, 1, 10)
		if err == nil {
			result["activities"] = activities
		}
	}

	if *fetchSegments {
		bounds := [4]float64{37.82, -122.52, 37.84, -122.35}
		explorer, err := client.ExploreSegments(bounds, "riding", 0, 5)
		if err == nil {
			result["segments"] = explorer.Segments
		}
	}

	if *fetchClub != 0 {
		club, err := client.GetClub(*fetchClub)
		if err == nil {
			result["club"] = club
		}
	}

	if *outputJSON {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(result)
	} else {
		fmt.Printf("%+v\n", result)
	}
}
