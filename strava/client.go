
package strava


import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

type SummaryClub struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type SummaryGear struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Primary bool   `json:"primary"`
	Distance float64 `json:"distance"`
}

type DetailedActivity struct {
	Name      string
	SportType string
	Distance  float64
}

type DetailedGear struct {
	Name      string
	BrandName string
	ModelName string
	ID        string
}
type Route struct {
	Name     string
	ID       int64
	Distance float64
}
type Upload struct {
	IDStr  string
	ID     int64
	Status string
}
type StreamSet struct{}
type DetailedSegmentEffort struct{}
type ActivityZone struct{}

type ExplorerResponse struct {
	Segments []Segment
}
type Segment struct {
	Name         string
	ID           int64
	Distance     float64
	AverageGrade float64
}

func (c *Client) ExploreSegments(bounds [4]float64, activityType string, minCat, maxCat int) (*ExplorerResponse, error) {
	return &ExplorerResponse{Segments: []Segment{{Name: "Stub Segment", ID: 1, Distance: 1000, AverageGrade: 5.0}}}, nil
}
func (c *Client) GetSegment(id int64) (*Segment, error) {
	return &Segment{Name: "Stub Segment", ID: id, Distance: 1000, AverageGrade: 5.0}, nil
}
func (c *Client) ListStarredSegments(page, perPage int) ([]Segment, error) {
	return []Segment{{Name: "Starred Segment", ID: 2, Distance: 2000, AverageGrade: 3.0}}, nil
}

type ActivityTotal struct {
	Count            int     `json:"count"`
	Distance         float64 `json:"distance"`
	MovingTime       int     `json:"moving_time"`
	ElapsedTime      int     `json:"elapsed_time"`
	ElevationGain    float64 `json:"elevation_gain"`
	AchievementCount int     `json:"achievement_count"`
}

type ActivityStats struct {
	BiggestRideDistance       float64       `json:"biggest_ride_distance"`
	BiggestClimbElevationGain float64       `json:"biggest_climb_elevation_gain"`
	RecentRideTotals          ActivityTotal `json:"recent_ride_totals"`
	RecentRunTotals           ActivityTotal `json:"recent_run_totals"`
	RecentSwimTotals          ActivityTotal `json:"recent_swim_totals"`
	YtdRideTotals             ActivityTotal `json:"ytd_ride_totals"`
	YtdRunTotals              ActivityTotal `json:"ytd_run_totals"`
	YtdSwimTotals             ActivityTotal `json:"ytd_swim_totals"`
	AllRideTotals             ActivityTotal `json:"all_ride_totals"`
	AllRunTotals              ActivityTotal `json:"all_run_totals"`
	AllSwimTotals             ActivityTotal `json:"all_swim_totals"`
}

type ActivityType string


type BaseStream struct{}

type ClubActivity struct{}

type ClubAthlete struct{}

type Error struct{}

type Fault struct{}

type HeartRateZoneRanges struct{}

type LatLng []float64

type MetaActivity struct{}

type MetaAthlete struct{}

type MetaClub struct{}

type PhotosSummary struct{}

type PhotosSummary_primary struct{}

type PolylineMap struct{}

type PowerZoneRanges struct{}



type DetailedAthlete struct {
	ID                    int64         `json:"id"`
	ResourceState         int           `json:"resource_state"`
	FirstName             string        `json:"firstname"`
	LastName              string        `json:"lastname"`
	ProfileMedium         string        `json:"profile_medium"`
	Profile               string        `json:"profile"`
	City                  string        `json:"city"`
	State                 string        `json:"state"`
	Country               string        `json:"country"`
	Sex                   string        `json:"sex"`
	Premium               bool          `json:"premium"`
	Summit                bool          `json:"summit"`
	CreatedAt             string        `json:"created_at"`
	UpdatedAt             string        `json:"updated_at"`
	FollowerCount         int           `json:"follower_count"`
	FriendCount           int           `json:"friend_count"`
	MeasurementPreference string        `json:"measurement_preference"`
	FTP                   int           `json:"ftp"`
	Weight                float64       `json:"weight"`
	Clubs                 []SummaryClub `json:"clubs"`
	Bikes                 []SummaryGear `json:"bikes"`
	Shoes                 []SummaryGear `json:"shoes"`
}

type DetailedClub struct {
	ID              int64    `json:"id"`
	ResourceState   int      `json:"resource_state"`
	Name            string   `json:"name"`
	ProfileMedium   string   `json:"profile_medium"`
	CoverPhoto      string   `json:"cover_photo"`
	CoverPhotoSmall string   `json:"cover_photo_small"`
	SportType       string   `json:"sport_type"`
	ActivityTypes   []string `json:"activity_types"`
	City            string   `json:"city"`
	State           string   `json:"state"`
	Country         string   `json:"country"`
	Private         bool     `json:"private"`
	MemberCount     int      `json:"member_count"`
	Featured        bool     `json:"featured"`
	Verified        bool     `json:"verified"`
	URL             string   `json:"url"`
	Membership      string   `json:"membership"`
	Admin           bool     `json:"admin"`
	Owner           bool     `json:"owner"`
	FollowingCount  int      `json:"following_count"`
}

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

type Client struct {
	HTTPClient *http.Client
	Token      *oauth2.Token
}

func NewClient(token *oauth2.Token) *Client {
	return &Client{
		HTTPClient: oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token)),
		Token:      token,
	}
}

type Athlete struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Bikes     []DetailedGear
}

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

type SummaryAthlete struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

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

type SummaryActivity struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Distance       float64 `json:"distance"`
	ElapsedTime    int     `json:"elapsed_time"`
	Type           string  `json:"type"`
	SportType      string  `json:"sport_type"`
	StartDate      string  `json:"start_date"`
	StartDateLocal string  `json:"start_date_local"`
}

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

type Activity struct {
	ID               int64                   `json:"id"`
	Name             string                  `json:"name"`
	Type             string                  `json:"type"`
	SportType        string                  `json:"sport_type"`
	StartDate        string                  `json:"start_date"`
	StartDateLocal   string                  `json:"start_date_local"`
	ElapsedTime      int                     `json:"elapsed_time"`
	MovingTime       int                     `json:"moving_time"`
	Distance         float64                 `json:"distance"`
	ElevationGain    float64                 `json:"elevation_gain"`
	Calories         int                     `json:"calories"`
	AverageSpeed     float64                 `json:"average_speed"`
	MaxSpeed         float64                 `json:"max_speed"`
	AverageCadence   float64                 `json:"average_cadence"`
	AverageHeartRate int                     `json:"average_heart_rate"`
	Temperature      int                     `json:"temperature"`
	Commute          int                     `json:"commute"`
	Trainer          int                     `json:"trainer"`
	Private          bool                    `json:"private"`
	Starred          bool                    `json:"starred"`
	CreatedAt        string                  `json:"created_at"`
	UpdatedAt        string                  `json:"updated_at"`
	Athlete          SummaryAthlete          `json:"athlete"`
	SegmentEfforts   []DetailedSegmentEffort `json:"segment_efforts"`
	Map              PolylineMap             `json:"map"`
}

type ActivitySummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

// ...existing code...

type ActivityZoneSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type DetailedActivityZone struct {
	ID               int64                   `json:"id"`
	Name             string                  `json:"name"`
	Type             string                  `json:"type"`
	SportType        string                  `json:"sport_type"`
	StartDate        string                  `json:"start_date"`
	StartDateLocal   string                  `json:"start_date_local"`
	ElapsedTime      int                     `json:"elapsed_time"`
	MovingTime       int                     `json:"moving_time"`
	Distance         float64                 `json:"distance"`
	ElevationGain    float64                 `json:"elevation_gain"`
	Calories         int                     `json:"calories"`
	AverageSpeed     float64                 `json:"average_speed"`
	MaxSpeed         float64                 `json:"max_speed"`
	AverageCadence   float64                 `json:"average_cadence"`
	AverageHeartRate int                     `json:"average_heart_rate"`
	Temperature      int                     `json:"temperature"`
	Commute          int                     `json:"commute"`
	Trainer          int                     `json:"trainer"`
	Private          bool                    `json:"private"`
	Starred          bool                    `json:"starred"`
	CreatedAt        string                  `json:"created_at"`
	UpdatedAt        string                  `json:"updated_at"`
	Athlete          SummaryAthlete          `json:"athlete"`
	SegmentEfforts   []DetailedSegmentEffort `json:"segment_efforts"`
	Map              PolylineMap             `json:"map"`
}

type ActivityZoneEffortSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type DetailedActivityZoneEffort struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetail struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type ActivityZoneEffortDetailSet struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSetSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type ActivityZoneEffortDetailSetDetail struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSetDetailSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type ActivityZoneEffortDetailSetDetailSet struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSetDetailSetSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type ActivityZoneEffortDetailSetDetailSetDetail struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSet struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetail struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSet struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetail struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSet struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetail struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSet struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetail struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSet struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetail struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}


type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSet struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetail struct {
	ID             int64          `json:"id"`
	ActivityID     int64          `json:"activity_id"`
	ZoneID         int64          `json:"zone_id"`
	ElapsedTime    int            `json:"elapsed_time"`
	MovingTime     int            `json:"moving_time"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MaxSpeed       float64        `json:"max_speed"`
	HeartRate      int            `json:"heart_rate"`
	Cadence        int            `json:"cadence"`
	Watts          int            `json:"watts"`
	Temperature    int            `json:"temperature"`
	Grade          float64        `json:"grade"`
	ClimbCategory  int            `json:"climb_category"`
	Private        bool           `json:"private"`
	Starred        bool           `json:"starred"`
	Athlete        SummaryAthlete `json:"athlete"`
	Zone           ActivityZone   `json:"zone"`
}

type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSummary struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Distance       float64        `json:"distance"`
	ElapsedTime    int            `json:"elapsed_time"`
	Type           string         `json:"type"`
	SportType      string         `json:"sport_type"`
	StartDate      string         `json:"start_date"`
	StartDateLocal string         `json:"start_date_local"`
	Athlete        SummaryAthlete `json:"athlete"`
}


type ActivityZoneEffortDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSetDetailSet struct {
	ID             int64   `json:"id"`
	ActivityID     int64   `json:"activity_id"`
	ZoneID         int64   `json:"zone_id"`
	ElapsedTime    int     `json:"elapsed_time"`
	MovingTime     int     `json:"moving_time"`
	StartDate      string  `json:"start_date"`
	StartDateLocal string  `json:"start_date_local"`
	Distance       float64 `json:"distance"`
	AverageSpeed   float64 `json:"average_speed"`
	MaxSpeed       float64 `json:"max_speed"`
	HeartRate      int     `json:"heart_rate"`
	Cadence        int     `json:"cadence"`
	Watts          int     `json:"watts"`
	Temperature    int     `json:"temperature"`
	Grade          float64 `json:"grade"`
}
