package strava

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"

	"golang.org/x/oauth2"
)

const (
	stravaAPIBase = "https://www.strava.com/api/v3"
)

// Client is the Strava API client
type Client struct {
	HTTPClient *http.Client
	Token      *oauth2.Token
}

// NewClient creates a new Strava API client
func NewClient(token *oauth2.Token) *Client {
	return &Client{
		HTTPClient: oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token)),
		Token:      token,
	}
}

// Athlete represents a Strava athlete
// Add more fields as needed
// https://developers.strava.com/docs/reference/#api-models-Athlete

type Athlete struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// DetailedActivity represents a Strava activity (partial, add more fields as needed)
type DetailedActivity struct {
	ID               int64   `json:"id"`
	Name             string  `json:"name"`
	Type             string  `json:"type"`
	SportType        string  `json:"sport_type"`
	StartDateLocal   string  `json:"start_date_local"`
	ElapsedTime      int     `json:"elapsed_time"`
	Description      string  `json:"description"`
	Distance         float64 `json:"distance"`
	Trainer          int     `json:"trainer"`
	Commute          int     `json:"commute"`
}

// CreateActivity creates a manual activity for an athlete
func (c *Client) CreateActivity(name, activityType, sportType, startDateLocal string, elapsedTime int, description string, distance float64, trainer, commute int) (*DetailedActivity, error) {
	url := stravaAPIBase + "/activities"
	data := make(map[string]interface{})
	data["name"] = name
	if activityType != "" {
		data["type"] = activityType
	}
	data["sport_type"] = sportType
	data["start_date_local"] = startDateLocal
	data["elapsed_time"] = elapsedTime
	if description != "" {
		data["description"] = description
	}
	if distance > 0 {
		data["distance"] = distance
	}
	if trainer != 0 {
		data["trainer"] = trainer
	}
	if commute != 0 {
		data["commute"] = commute
	}
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var activity DetailedActivity
	if err := json.NewDecoder(resp.Body).Decode(&activity); err != nil {
		return nil, err
	}
	return &activity, nil
}

// GetActivityByID returns the activity with the given ID
func (c *Client) GetActivityByID(id int64, includeAllEfforts bool) (*DetailedActivity, error) {
	url := fmt.Sprintf("%s/activities/%d", stravaAPIBase, id)
	if includeAllEfforts {
		url += "?include_all_efforts=true"
	}
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var activity DetailedActivity
	if err := json.NewDecoder(resp.Body).Decode(&activity); err != nil {
		return nil, err
	}
	return &activity, nil
}

// GetAthlete returns the currently authenticated athlete
func (c *Client) GetAthlete() (*Athlete, error) {
	url := fmt.Sprintf("%s/athlete", stravaAPIBase)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var athlete Athlete
	if err := json.NewDecoder(resp.Body).Decode(&athlete); err != nil {
		return nil, err
	}
	return &athlete, nil
}
