package strava

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

// ActivityTotal represents a roll-up of metrics pertaining to a set of activities.
type ActivityTotal struct {
	Count            int     `json:"count"`
	Distance         float64 `json:"distance"`
	MovingTime       int     `json:"moving_time"`
	ElapsedTime      int     `json:"elapsed_time"`
	ElevationGain    float64 `json:"elevation_gain"`
	AchievementCount int     `json:"achievement_count"`
}

// ActivityStats is a set of rolled-up statistics and totals for an athlete
type ActivityStats struct {
	BiggestRideDistance         float64      `json:"biggest_ride_distance"`
	BiggestClimbElevationGain   float64      `json:"biggest_climb_elevation_gain"`
	RecentRideTotals            ActivityTotal `json:"recent_ride_totals"`
	RecentRunTotals             ActivityTotal `json:"recent_run_totals"`
	RecentSwimTotals            ActivityTotal `json:"recent_swim_totals"`
	YtdRideTotals               ActivityTotal `json:"ytd_ride_totals"`
	YtdRunTotals                ActivityTotal `json:"ytd_run_totals"`
	YtdSwimTotals               ActivityTotal `json:"ytd_swim_totals"`
	AllRideTotals               ActivityTotal `json:"all_ride_totals"`
	AllRunTotals                ActivityTotal `json:"all_run_totals"`
	AllSwimTotals               ActivityTotal `json:"all_swim_totals"`
}

// ActivityType is a stub for the ActivityType model
type ActivityType string

// ActivityZone is a stub for the ActivityZone model
type ActivityZone struct{}

// BaseStream is a stub for the BaseStream model
type BaseStream struct{}

// ClubActivity is a stub for the ClubActivity model
type ClubActivity struct{}

// ClubAthlete is a stub for the ClubAthlete model
type ClubAthlete struct{}

// Error is a stub for the Error model
type Error struct{}

// Fault is a stub for the Fault model
type Fault struct{}

// HeartRateZoneRanges is a stub for the HeartRateZoneRanges model
type HeartRateZoneRanges struct{}

// LatLng is a stub for the LatLng model
type LatLng []float64

// MetaActivity is a stub for the MetaActivity model
type MetaActivity struct{}

// MetaAthlete is a stub for the MetaAthlete model
type MetaAthlete struct{}

// MetaClub is a stub for the MetaClub model
type MetaClub struct{}

// PhotosSummary is a stub for the PhotosSummary model
type PhotosSummary struct{}

// PhotosSummary_primary is a stub for the PhotosSummary_primary model
type PhotosSummary_primary struct{}

// PolylineMap is a stub for the PolylineMap model
type PolylineMap struct{}

// PowerZoneRanges is a stub for the PowerZoneRanges model
type PowerZoneRanges struct{}

// Route is a stub for the Route model
type Route struct{}

// Split is a stub for the Split model
type Split struct{}

// SportType is a stub for the SportType model
type SportType string

// StreamSet is a stub for the StreamSet model
type StreamSet struct{}

// SummaryGear represents a summary of gear
type SummaryGear struct {
	ID            string  `json:"id"`
	ResourceState int     `json:"resource_state"`
	Primary       bool    `json:"primary"`
	Name          string  `json:"name"`
	Distance      float64 `json:"distance"`
}

// SummaryPRSegmentEffort is a stub for the SummaryPRSegmentEffort model
type SummaryPRSegmentEffort struct{}

// SummarySegment is a stub for the SummarySegment model
type SummarySegment struct{}

// SummarySegmentEffort is a stub for the SummarySegmentEffort model
type SummarySegmentEffort struct{}

// TimedZoneDistribution is a stub for the TimedZoneDistribution model
type TimedZoneDistribution struct{}

// UpdatableActivity is a stub for the UpdatableActivity model
type UpdatableActivity struct{}

// Upload is a stub for the Upload model
type Upload struct{}

// Waypoint is a stub for the Waypoint model
type Waypoint struct{}

// ZoneRange is a stub for the ZoneRange model
type ZoneRange struct{}

// ZoneRanges is a stub for the ZoneRanges model
type ZoneRanges struct{}

// Zones is a stub for the Zones model
type Zones struct{}

// AltitudeStream is a stub for the AltitudeStream model
type AltitudeStream struct{}

// CadenceStream is a stub for the CadenceStream model
type CadenceStream struct{}

// DetailedGear is a stub for the DetailedGear model
type DetailedGear struct{}

// DetailedSegment is a stub for the DetailedSegment model
type DetailedSegment struct{}

// DetailedSegmentEffort is a stub for the DetailedSegmentEffort model
type DetailedSegmentEffort struct{}

// DistanceStream is a stub for the DistanceStream model
type DistanceStream struct{}

// HeartrateStream is a stub for the HeartrateStream model
type HeartrateStream struct{}

// LatLngStream is a stub for the LatLngStream model
type LatLngStream struct{}

// MovingStream is a stub for the MovingStream model
type MovingStream struct{}

// PowerStream is a stub for the PowerStream model
type PowerStream struct{}

// SmoothGradeStream is a stub for the SmoothGradeStream model
type SmoothGradeStream struct{}

// SmoothVelocityStream is a stub for the SmoothVelocityStream model
type SmoothVelocityStream struct{}

// SummaryClub represents a summary of a club
type SummaryClub struct {
	ID            int64  `json:"id"`
	ResourceState int    `json:"resource_state"`
	Name          string `json:"name"`
	ProfileMedium string `json:"profile_medium"`
	CoverPhoto    string `json:"cover_photo"`
	CoverPhotoSmall string `json:"cover_photo_small"`
	SportType     string `json:"sport_type"`
	ActivityTypes []string `json:"activity_types"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	Private       bool   `json:"private"`
	MemberCount   int    `json:"member_count"`
	Featured      bool   `json:"featured"`
	Verified      bool   `json:"verified"`
	URL           string `json:"url"`
}

// TemperatureStream is a stub for the TemperatureStream model
type TemperatureStream struct{}

// TimeStream is a stub for the TimeStream model
type TimeStream struct{}

// TimedZoneRange is a stub for the TimedZoneRange model
type TimedZoneRange struct{}

// DetailedAthlete represents a detailed athlete
type DetailedAthlete struct {
	ID                int64   `json:"id"`
	ResourceState     int     `json:"resource_state"`
	FirstName         string  `json:"firstname"`
	LastName          string  `json:"lastname"`
	ProfileMedium     string  `json:"profile_medium"`
	Profile           string  `json:"profile"`
	City              string  `json:"city"`
	State             string  `json:"state"`
	Country           string  `json:"country"`
	Sex               string  `json:"sex"`
	Premium           bool    `json:"premium"`
	Summit            bool    `json:"summit"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
	FollowerCount     int     `json:"follower_count"`
	FriendCount       int     `json:"friend_count"`
	MeasurementPreference string `json:"measurement_preference"`
	FTP               int     `json:"ftp"`
	Weight            float64 `json:"weight"`
	Clubs             []SummaryClub `json:"clubs"`
	Bikes             []SummaryGear `json:"bikes"`
	Shoes             []SummaryGear `json:"shoes"`
}

// DetailedGear represents detailed information about gear
type DetailedGear struct {
	ID            string  `json:"id"`
	Primary       bool    `json:"primary"`
	Name          string  `json:"name"`
	Distance      float64 `json:"distance"`
	BrandName     string  `json:"brand_name"`
	ModelName     string  `json:"model_name"`
	FrameType     int     `json:"frame_type"`
	Description   string  `json:"description"`
}

// DetailedClub represents a detailed club
type DetailedClub struct {
	ID            int64  `json:"id"`
	ResourceState int    `json:"resource_state"`
	Name          string `json:"name"`
	ProfileMedium string `json:"profile_medium"`
	CoverPhoto    string `json:"cover_photo"`
	CoverPhotoSmall string `json:"cover_photo_small"`
	SportType     string `json:"sport_type"`
	ActivityTypes []string `json:"activity_types"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	Private       bool   `json:"private"`
	MemberCount   int    `json:"member_count"`
	Featured      bool   `json:"featured"`
	Verified      bool   `json:"verified"`
	URL           string `json:"url"`
	Membership    string `json:"membership"`
	Admin         bool   `json:"admin"`
	Owner         bool   `json:"owner"`
	FollowingCount int   `json:"following_count"`
}

// Route represents a Strava route
type Route struct {
	ID              int64        `json:"id"`
	IDStr           string       `json:"id_str"`
	Name            string       `json:"name"`
	Description     string       `json:"description"`
	Distance        float64      `json:"distance"`
	ElevationGain   float64      `json:"elevation_gain"`
	Private         bool         `json:"private"`
	Starred         bool         `json:"starred"`
	Timestamp       int64        `json:"timestamp"`
	Type            int          `json:"type"`
	SubType         int          `json:"sub_type"`
	CreatedAt       string       `json:"created_at"`
	UpdatedAt       string       `json:"updated_at"`
	EstimatedMovingTime int      `json:"estimated_moving_time"`
	Athlete         SummaryAthlete `json:"athlete"`
	Map             PolylineMap  `json:"map"`
}

// Upload represents an upload response
type Upload struct {
	ID        int64  `json:"id"`
	IDStr     string `json:"id_str"`
	ExternalID string `json:"external_id"`
	Error     string `json:"error"`
	Status    string `json:"status"`
	ActivityID int64 `json:"activity_id"`
}

// StreamSet represents a set of activity streams
type StreamSet struct {
	Time           *TimeStream        `json:"time,omitempty"`
	Distance       *DistanceStream    `json:"distance,omitempty"`
	LatLng         *LatLngStream      `json:"latlng,omitempty"`
	Altitude       *AltitudeStream    `json:"altitude,omitempty"`
	VelocitySmooth *SmoothVelocityStream `json:"velocity_smooth,omitempty"`
	Heartrate      *HeartrateStream   `json:"heartrate,omitempty"`
	Cadence        *CadenceStream     `json:"cadence,omitempty"`
	Watts          *PowerStream       `json:"watts,omitempty"`
	Temp           *TemperatureStream `json:"temp,omitempty"`
	Moving         *MovingStream      `json:"moving,omitempty"`
	GradeSmooth    *SmoothGradeStream `json:"grade_smooth,omitempty"`
}

// GetAthleteStats returns the activity stats of an athlete
func (c *Client) GetAthleteStats(athleteID int64) (*ActivityStats, error) {
	url := fmt.Sprintf("%s/athletes/%d/stats", stravaAPIBase, athleteID)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var stats ActivityStats
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, err
	}
	return &stats, nil
}

// GetSummaryGear returns a summary of gear by gear ID
func (c *Client) GetSummaryGear(gearID string) (*SummaryGear, error) {
	url := fmt.Sprintf("%s/gear/%s", stravaAPIBase, gearID)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var gear SummaryGear
	if err := json.NewDecoder(resp.Body).Decode(&gear); err != nil {
		return nil, err
	}
	return &gear, nil
}

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
//
// Athlete holds basic information about a Strava user
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

// Comment represents a comment on an activity
type Comment struct {
	ID         int64  `json:"id"`
	ActivityID int64  `json:"activity_id"`
	Text       string `json:"text"`
	CreatedAt  string `json:"created_at"`
	Athlete    struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
	} `json:"athlete"`
	Cursor string `json:"cursor"`
}

// SummaryAthlete represents a summary of an athlete (for kudoers)
type SummaryAthlete struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Lap represents a lap in an activity
type Lap struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	ElapsedTime    int     `json:"elapsed_time"`
	MovingTime     int     `json:"moving_time"`
	StartDate      string  `json:"start_date"`
	StartDateLocal string  `json:"start_date_local"`
	Distance       float64 `json:"distance"`
	AverageSpeed   float64 `json:"average_speed"`
	MaxSpeed       float64 `json:"max_speed"`
	LapIndex       int     `json:"lap_index"`
}

// SummaryActivity represents a summary of an activity (for activity lists)
type SummaryActivity struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
}

// ListActivityComments returns the comments on the given activity
func (c *Client) ListActivityComments(activityID int64, pageSize int, afterCursor string) ([]Comment, error) {
	url := fmt.Sprintf("%s/activities/%d/comments?page_size=%d", stravaAPIBase, activityID, pageSize)
	if afterCursor != "" {
		url += "&after_cursor=" + afterCursor
	}
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var comments []Comment
	if err := json.NewDecoder(resp.Body).Decode(&comments); err != nil {
		return nil, err
	}
	return comments, nil
}

// ListActivityKudoers returns the athletes who kudoed an activity
func (c *Client) ListActivityKudoers(activityID int64, page, perPage int) ([]SummaryAthlete, error) {
	url := fmt.Sprintf("%s/activities/%d/kudos?page=%d&per_page=%d", stravaAPIBase, activityID, page, perPage)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var athletes []SummaryAthlete
	if err := json.NewDecoder(resp.Body).Decode(&athletes); err != nil {
		return nil, err
	}
	return athletes, nil
}

// ListActivityLaps returns the laps of an activity
func (c *Client) ListActivityLaps(activityID int64) ([]Lap, error) {
	url := fmt.Sprintf("%s/activities/%d/laps", stravaAPIBase, activityID)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var laps []Lap
	if err := json.NewDecoder(resp.Body).Decode(&laps); err != nil {
		return nil, err
	}
	return laps, nil
}

// ListAthleteActivities returns the activities of the authenticated athlete
func (c *Client) ListAthleteActivities(before, after int64, page, perPage int) ([]SummaryActivity, error) {
	url := fmt.Sprintf("%s/athlete/activities?before=%d&after=%d&page=%d&per_page=%d", stravaAPIBase, before, after, page, perPage)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var activities []SummaryActivity
	if err := json.NewDecoder(resp.Body).Decode(&activities); err != nil {
		return nil, err
	}
	return activities, nil
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

// GetDetailedGear returns detailed gear by gear ID
func (c *Client) GetDetailedGear(gearID string) (*DetailedGear, error) {
	url := fmt.Sprintf("%s/gear/%s", stravaAPIBase, gearID)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var gear DetailedGear
	if err := json.NewDecoder(resp.Body).Decode(&gear); err != nil {
		return nil, err
	}
	return &gear, nil
}

// GetClub returns a detailed club by club ID
func (c *Client) GetClub(clubID int64) (*DetailedClub, error) {
	url := fmt.Sprintf("%s/clubs/%d", stravaAPIBase, clubID)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var club DetailedClub
	if err := json.NewDecoder(resp.Body).Decode(&club); err != nil {
		return nil, err
	}
	return &club, nil
}

// GetRoute returns a route by route ID
func (c *Client) GetRoute(routeID int64) (*Route, error) {
	url := fmt.Sprintf("%s/routes/%d", stravaAPIBase, routeID)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var route Route
	if err := json.NewDecoder(resp.Body).Decode(&route); err != nil {
		return nil, err
	}
	return &route, nil
}

// GetUpload returns an upload by upload ID
func (c *Client) GetUpload(uploadID int64) (*Upload, error) {
	url := fmt.Sprintf("%s/uploads/%d", stravaAPIBase, uploadID)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var upload Upload
	if err := json.NewDecoder(resp.Body).Decode(&upload); err != nil {
		return nil, err
	}
	return &upload, nil
}

// GetActivityStreams returns the streams for an activity
func (c *Client) GetActivityStreams(activityID int64, keys []string, keyByType bool) (*StreamSet, error) {
	url := fmt.Sprintf("%s/activities/%d/streams?keys=%s&key_by_type=%t", stravaAPIBase, activityID, joinKeys(keys), keyByType)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}
	var streams StreamSet
	if err := json.NewDecoder(resp.Body).Decode(&streams); err != nil {
		return nil, err
	}
	return &streams, nil
}

// joinKeys joins stream keys for the query string
func joinKeys(keys []string) string {
	if len(keys) == 0 {
		return ""
	}
	result := keys[0]
	for _, k := range keys[1:] {
		result += "," + k
	}
	return result
}

// DetailedSegment represents a detailed segment
type DetailedSegment struct {
	ID              int64    `json:"id"`
	Name            string   `json:"name"`
	ActivityType    string   `json:"activity_type"`
	Distance        float64  `json:"distance"`
	AverageGrade    float64  `json:"average_grade"`
	MaximumGrade    float64  `json:"maximum_grade"`
	ElevationHigh   float64  `json:"elevation_high"`
	ElevationLow    float64  `json:"elevation_low"`
	StartLatLng     []float64 `json:"start_latlng"`
	EndLatLng       []float64 `json:"end_latlng"`
	ClimbCategory   int      `json:"climb_category"`
	City            string   `json:"city"`
	State           string   `json:"state"`
	Country         string   `json:"country"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
}

// DetailedSegmentEffort represents a detailed segment effort
type DetailedSegmentEffort struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	SegmentID       int64    `json:"segment_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Segment         DetailedSegment `json:"segment"`
}

// Activity represents a Strava activity
type Activity struct {
	ID               int64   `json:"id"`
	Name             string  `json:"name"`
	Type             string  `json:"type"`
	SportType        string  `json:"sport_type"`
	StartDate        string  `json:"start_date"`
	StartDateLocal   string  `json:"start_date_local"`
	ElapsedTime      int     `json:"elapsed_time"`
	MovingTime       int     `json:"moving_time"`
	Distance         float64 `json:"distance"`
	ElevationGain    float64 `json:"elevation_gain"`
	Calories         int     `json:"calories"`
	AverageSpeed     float64 `json:"average_speed"`
	MaxSpeed         float64 `json:"max_speed"`
	AverageCadence   float64 `json:"average_cadence"`
	AverageHeartRate int     `json:"average_heart_rate"`
	Temperature      int     `json:"temperature"`
	Commute          int     `json:"commute"`
	Trainer          int     `json:"trainer"`
	Private          bool    `json:"private"`
	Starred          bool    `json:"starred"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	Athlete          SummaryAthlete `json:"athlete"`
	SegmentEfforts   []DetailedSegmentEffort `json:"segment_efforts"`
	Map              PolylineMap  `json:"map"`
}

// ActivitySummary represents a summary of an activity
type ActivitySummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// DetailedActivity represents a detailed activity
type DetailedActivity struct {
	ID               int64   `json:"id"`
	Name             string  `json:"name"`
	Type             string  `json:"type"`
	SportType        string  `json:"sport_type"`
	StartDate        string  `json:"start_date"`
	StartDateLocal   string  `json:"start_date_local"`
	ElapsedTime      int     `json:"elapsed_time"`
	MovingTime       int     `json:"moving_time"`
	Distance         float64 `json:"distance"`
	ElevationGain    float64 `json:"elevation_gain"`
	Calories         int     `json:"calories"`
	AverageSpeed     float64 `json:"average_speed"`
	MaxSpeed         float64 `json:"max_speed"`
	AverageCadence   float64 `json:"average_cadence"`
	AverageHeartRate int     `json:"average_heart_rate"`
	Temperature      int     `json:"temperature"`
	Commute          int     `json:"commute"`
	Trainer          int     `json:"trainer"`
	Private          bool    `json:"private"`
	Starred          bool    `json:"starred"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	Athlete          SummaryAthlete `json:"athlete"`
	SegmentEfforts   []DetailedSegmentEffort `json:"segment_efforts"`
	Map              PolylineMap  `json:"map"`
}

// ActivityZone represents an activity zone
type ActivityZone struct {
	ID              int64    `json:"id"`
	Name            string   `json:"name"`
	ActivityType    string   `json:"activity_type"`
	Distance        float64  `json:"distance"`
	AverageGrade    float64  `json:"average_grade"`
	MaximumGrade    float64  `json:"maximum_grade"`
	ElevationHigh   float64  `json:"elevation_high"`
	ElevationLow    float64  `json:"elevation_low"`
	StartLatLng     []float64 `json:"start_latlng"`
	EndLatLng       []float64 `json:"end_latlng"`
	ClimbCategory   int      `json:"climb_category"`
	City            string   `json:"city"`
	State           string   `json:"state"`
	Country         string   `json:"country"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffort represents an effort in an activity zone
type ActivityZoneEffort struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneSummary represents a summary of an activity zone
type ActivityZoneSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// DetailedActivityZone represents a detailed activity zone
type DetailedActivityZone struct {
	ID               int64   `json:"id"`
	Name             string  `json:"name"`
	Type             string  `json:"type"`
	SportType        string  `json:"sport_type"`
	StartDate        string  `json:"start_date"`
	StartDateLocal   string  `json:"start_date_local"`
	ElapsedTime      int     `json:"elapsed_time"`
	MovingTime       int     `json:"moving_time"`
	Distance         float64 `json:"distance"`
	ElevationGain    float64 `json:"elevation_gain"`
	Calories         int     `json:"calories"`
	AverageSpeed     float64 `json:"average_speed"`
	MaxSpeed         float64 `json:"max_speed"`
	AverageCadence   float64 `json:"average_cadence"`
	AverageHeartRate int     `json:"average_heart_rate"`
	Temperature      int     `json:"temperature"`
	Commute          int     `json:"commute"`
	Trainer          int     `json:"trainer"`
	Private          bool    `json:"private"`
	Starred          bool    `json:"starred"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	Athlete          SummaryAthlete `json:"athlete"`
	SegmentEfforts   []DetailedSegmentEffort `json:"segment_efforts"`
	Map              PolylineMap  `json:"map"`
}

// ActivityZoneEffortSummary represents a summary of an activity zone effort
type ActivityZoneEffortSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// DetailedActivityZoneEffort represents a detailed activity zone effort
type DetailedActivityZoneEffort struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetail represents the detail of an activity zone effort
type ActivityZoneEffortDetail struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSummary represents a summary of an activity zone effort detail
type ActivityZoneEffortDetailSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSet represents a set of activity zone effort details
type ActivityZoneEffortDetailSet struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSetSummary represents a summary of an activity zone effort detail set
type ActivityZoneEffortDetailSetSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSetDetail represents a detailed set of activity zone effort details
type ActivityZoneEffortDetailSetDetail struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSetDetailSummary represents a summary of a detailed set of activity zone effort details
type ActivityZoneEffortDetailSetDetailSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSetDetailSet represents a set of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSet struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSetDetailSetSummary represents a summary of a set of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSetDetailSetDetail represents a detailed set of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetail struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSummary represents a summary of a detailed set of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSet represents a set of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSet struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetSummary represents a summary of a set of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetail represents a detailed set of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetail struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSummary represents a summary of a detailed set of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSet represents a set of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSet struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetSummary represents a summary of a set of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetail represents a detailed set of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetail struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSummary represents a summary of a detailed set of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSet represents a set of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSet struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetSummary represents a summary of a set of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetail represents a detailed set of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetail struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSummary represents a summary of a detailed set of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSet represents a set of detailed sets of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSet struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetSummary represents a summary of a set of detailed sets of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetail represents a detailed set of detailed sets of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetail struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSummary represents a summary of a detailed set of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSet represents a set of detailed sets of detailed sets of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSet struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetSummary represents a summary of a set of detailed sets of detailed sets of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetail represents a detailed set of detailed sets of detailed sets of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetail struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSummary represents a summary of a detailed set of detailed sets of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSet represents a set of detailed sets of detailed sets of detailed sets of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSet struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetSummary represents a summary of a set of detailed sets of detailed sets of detailed sets of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetail represents a detailed set of detailed sets of detailed sets of detailed sets of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetail struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
	ClimbCategory   int      `json:"climb_category"`
	Private         bool     `json:"private"`
	Starred         bool     `json:"starred"`
	Athlete         SummaryAthlete `json:"athlete"`
	Zone             ActivityZone `json:"zone"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSummary represents a summary of a detailed set of detailed sets of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSummary struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Distance   float64 `json:"distance"`
	ElapsedTime int    `json:"elapsed_time"`
	Type       string  `json:"type"`
	SportType  string  `json:"sport_type"`
	StartDate  string  `json:"start_date"`
	StartDateLocal string `json:"start_date_local"`
	Athlete    SummaryAthlete `json:"athlete"`
}

// ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSet represents a set of detailed sets of detailed sets of detailed sets of detailed sets of detailed sets of detailed sets of detailed sets of detailed activity zone effort detail sets
type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSet struct {
	ID              int64    `json:"id"`
	ActivityID      int64    `json:"activity_id"`
	ZoneID          int64    `json:"zone_id"`
	ElapsedTime     int      `json:"elapsed_time"`
	MovingTime      int      `json:"moving_time"`
	StartDate       string   `json:"start_date"`
	StartDateLocal  string   `json:"start_date_local"`
	Distance        float64  `json:"distance"`
	AverageSpeed    float64  `json:"average_speed"`
	MaxSpeed        float64  `json:"max_speed"`
	HeartRate       int      `json:"heart_rate"`
	Cadence         int      `json:"cadence"`
	Watts           int      `json:"watts"`
	Temperature     int      `json:"temperature"`
	Grade           float64  `json:"grade"`
}
