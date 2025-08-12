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

	fmt.Println("Strava API Wrapper Demo")

	athlete, err := client.GetAthlete()
	if err != nil {
		log.Fatalf("Error getting athlete: %v", err)
	}
	fmt.Printf("Hello, %s %s! (ID: %d)\n", athlete.FirstName, athlete.LastName, athlete.ID)

	stats, err := client.GetAthleteStats(athlete.ID)
	if err == nil {
		fmt.Printf("Recent ride total distance: %.2f meters\n", stats.RecentRideTotals.Distance)
	}

	activities, err := client.ListAthleteActivities(0, 0, 1, 3)
	if err == nil && len(activities) > 0 {
		fmt.Printf("Recent activities:\n")
		for _, act := range activities {
			fmt.Printf("- %s (ID: %d, Distance: %.2fm)\n", act.Name, act.ID, act.Distance)
		}
		activity, err := client.GetActivityByID(activities[0].ID, false)
		if err == nil {
			fmt.Printf("First activity details: %s, %s, %.2fm\n", activity.Name, activity.SportType, activity.Distance)
		}
		comments, err := client.ListActivityComments(activities[0].ID, 5, "")
		if err == nil && len(comments) > 0 {
			fmt.Printf("Comments on first activity:\n")
			for _, c := range comments {
				fmt.Printf("- %s: %s\n", c.Athlete.FirstName, c.Text)
			}
		}
	}

	bounds := [4]float64{37.82, -122.52, 37.84, -122.35}
	explorer, err := client.ExploreSegments(bounds, "riding", 0, 5)
	if err == nil && len(explorer.Segments) > 0 {
		seg := explorer.Segments[0]
		fmt.Printf("Explorer segment: %s (ID: %d)\n", seg.Name, seg.ID)
		segment, err := client.GetSegment(seg.ID)
		if err == nil {
			fmt.Printf("Detailed segment: %s, %.2fm, avg grade %.2f%%\n", segment.Name, segment.Distance, segment.AverageGrade)
		}
	}
	starred, err := client.ListStarredSegments(1, 3)
	if err == nil && len(starred) > 0 {
		fmt.Printf("Starred segments:\n")
		for _, s := range starred {
			fmt.Printf("- %s (ID: %d)\n", s.Name, s.ID)
		}
	}

	club, err := client.GetClub(1)
	if err == nil {
		fmt.Printf("Club: %s (ID: %d, Members: %d)\n", club.Name, club.ID, club.MemberCount)
	}

	if len(athlete.Bikes) > 0 {
		gear, err := client.GetDetailedGear(athlete.Bikes[0].ID)
		if err == nil {
			fmt.Printf("Bike: %s, Brand: %s, Model: %s\n", gear.Name, gear.BrandName, gear.ModelName)
		}
	}

	route, err := client.GetRoute(1)
	if err == nil {
		fmt.Printf("Route: %s (ID: %d, Distance: %.2fm)\n", route.Name, route.ID, route.Distance)
	}

	upload, err := client.GetUpload(1)
	if err == nil {
		fmt.Printf("Upload: %s (ID: %d, Status: %s)\n", upload.IDStr, upload.ID, upload.Status)
	}

	if len(activities) > 0 {
		streams, err := client.GetActivityStreams(activities[0].ID, []string{"distance", "latlng"}, true)
		if err == nil && streams != nil {
			fmt.Printf("Got streams for activity %d\n", activities[0].ID)
		}
	}
}
