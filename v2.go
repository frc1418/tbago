package tba

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	TBA_API_URL = "https://www.thebluealliance.com/api/v2/" // Base URL for The Blue Alliance's API
)

// Deprecated: APIv2 is deprecated
// Init creates a new client for The Blue Alliance.
func Init(developer, application, version string) (client Client, err error) {
	if len(developer) == 1 || len(application) == 1 || len(version) == 1 {
		err = errors.New("Invalid arguments")
		return client, err
	}
	client = Client{
		key:     developer + ":" + application + ":" + version,
		version: 2,
	}
	os.Stderr.WriteString("The Blue Alliance APIv2 will soon be deprecated\n")
	os.Stderr.WriteString("See https://github.com/CarlColglazier/tba for more information\n\n")
	return client, err
}

func (tba Client) appID() string {
	return tba.key
}

type Team struct {
	Address          string `json:"address"`
	City             string `json:"city"`
	Country          string `json:"country"`
	GmapsPlaceID     string `json:"gmaps_place_id"`
	GmapsURL         string `json:"gmaps_url"`
	HomeChampionship struct {
		Num2017 string `json:"2017"`
		Num2018 string `json:"2018"`
	} `json:"home_championship"`
	Key          string  `json:"key"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	LocationName string  `json:"location_name"`
	Motto        string  `json:"motto"`
	Name         string  `json:"name"`
	Nickname     string  `json:"nickname"`
	PostalCode   string  `json:"postal_code"`
	RookieYear   int     `json:"rookie_year"`
	StateProv    string  `json:"state_prov"`
	TeamNumber   int     `json:"team_number"`
	Website      string  `json:"website"`
	Location     string  // Deprecated
	Locality     string  // Deprecated
	Region       string  // Deprecated
	CountryName  string  // Deprecated
}

// Deprecated: APIv2 is deprecated
// Event represents the event model.
type Event struct {
	Key                 string        `json:"key"`
	Website             string        `json:"website"`
	Official            bool          `json:"official"`
	EndDate             string        `json:"end_date"`
	Name                string        `json:"name"`
	ShortName           string        `json:"short_name"`
	FacebookEid         interface{}   `json:"facebook_eid"`
	EventDistrictString interface{}   `json:"event_district_string"`
	VenueAddress        string        `json:"venue_address"`
	EventDistrict       int           `json:"event_district"`
	Location            string        `json:"location"`
	EventCode           string        `json:"event_code"`
	Year                int           `json:"year"`
	Webcast             []interface{} `json:"webcast"`
	Timezone            string        `json:"timezone"`
	Alliances           []struct {
		Declines []interface{} `json:"declines"`
		Picks    []string      `json:"picks"`
	} `json:"alliances"`
	EventTypeString string `json:"event_type_string"`
	StartDate       string `json:"start_date"`
	EventType       int    `json:"event_type"`
}

// Deprecated: APIv2 is deprecated
// Match represents the match model.
type Match struct {
	CompLevel   string `json:"comp_level"`
	MatchNumber int    `json:"match_number"`
	Videos      []struct {
		Type string `json:"type"`
		Key  string `json:"key"`
	} `json:"videos"`
	TimeString     string      `json:"time_string"`
	SetNumber      int         `json:"set_number"`
	Key            string      `json:"key"`
	Time           int         `json:"time"`
	ScoreBreakdown interface{} `json:"score_breakdown"`
	Alliances      struct {
		Blue struct {
			Score int      `json:"score"`
			Teams []string `json:"teams"`
		} `json:"blue"`
		Red struct {
			Score int      `json:"score"`
			Teams []string `json:"teams"`
		} `json:"red"`
	} `json:"alliances"`
	EventKey string `json:"event_key"`
}

// Deprecated: APIv2 is deprecated
// Award represents the award model.
type Award struct {
	EventKey      string `json:"event_key"`
	AwardType     int    `json:"award_type"`
	Name          string `json:"name"`
	RecipientList []struct {
		TeamNumber int         `json:"team_number"`
		Awardee    interface{} `json:"awardee"`
	} `json:"recipient_list"`
	Year int `json:"year"`
}

// Deprecated: APIv2 is deprecated
// Media represents the media model.
type Media []struct {
	Type    string `json:"type"`
	Details struct {
		ImagePartial string `json:"image_partial"`
	} `json:"details"`
	ForeignKey string `json:"foreign_key"`
}

// Deprecated: APIv2 is deprecated
// Robot represents the robot model.
type Robot struct {
	TeamKey string `json:"team_key"`
	Name    string `json:"name"`
	Key     string `json:"key"`
	Year    int    `json:"year"`
}

// Deprecated: APIv2 is deprecated
// Stat represents the stat model.
type Stats struct {
	OPRs         interface{} `json:"oprs"`
	YearSpecific struct {
		Qual    interface{}
		Playoff interface{}
	}
	CCWMs map[string]float64 `json:"ccwms"`
}

// Deprecated: APIv2 is deprecated
type Ranking struct {
	PointTotal  int         `json:"point_total"`
	TeamKey     string      `json:"team_key"`
	eventPoints interface{} `json:"event_points"`
	Rank        int         `json:"rank"`
	RookieBonus int         `json:"rookie_bonus"`
}

// Deprecated: APIv2 is deprecated
type DistrictPoints struct {
	Points      interface{}
	Tiebreakers interface{}
}

// sendRequest sends a HTTP GET request to The Blue Alliance API.
// An HTTP response is returned.
func (tba Client) sendRequest(url string) (*http.Response, error) {
	// Use a custom client in order to the set 'X-TBA-App-Id' header.
	requestUrl := TBA_API_URL + url
	client := &http.Client{}
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-TBA-App-Id", tba.appID())
	return client.Do(req)
}

func (tba Client) get(url string) ([]byte, error) {
	resp, err := tba.sendRequest(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Response code was not 200")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func (tba Client) jsonToStruct(url string, v interface{}) error {
	resp, err := tba.get(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, v)
	return err
}

// Deprecated: APIv2 is deprecated
// Team list request
func (tba Client) GetTeams(pageNum int) ([]Team, error) {
	url := fmt.Sprintf("teams/%d", pageNum)
	var t []Team
	err := tba.jsonToStruct(url, &t)
	return t, err
}

// Deprecated: APIv2 is deprecated
// Team request
func (tba Client) GetTeam(team int) (Team, error) {
	url := fmt.Sprintf("team/frc%d", team)
	var t Team
	err := tba.jsonToStruct(url, &t)
	return t, err
}

// Deprecated: APIv2 is deprecated
// Event request
func (tba Client) GetEvent(key string) (Event, error) {
	url := fmt.Sprintf("event/%s", key)
	var e Event
	err := tba.jsonToStruct(url, &e)
	return e, err
}

// Deprecated: APIv2 is deprecated
// Team events request
func (tba Client) GetTeamEvents(team int, year int) ([]Event, error) {
	url := fmt.Sprintf("team/frc%d/%d/events", team, year)
	var e []Event
	err := tba.jsonToStruct(url, &e)
	return e, err
}

// Deprecated: APIv2 is deprecated
// Team awards request
func (tba Client) GetTeamAwards(team int) ([]Award, error) {
	url := fmt.Sprintf("team/frc%d/awards", team)
	var a []Award
	err := tba.jsonToStruct(url, &a)
	return a, err
}

// Deprecated: APIv2 is deprecated
// Team event awards request
func (tba Client) GetTeamEventAwards(team int, event string) ([]Award, error) {
	url := fmt.Sprintf("team/frc%d/event/%s/awards", team, event)
	var a []Award
	err := tba.jsonToStruct(url, &a)
	return a, err
}

// Deprecated: APIv2 is deprecated
// Team years active request
func (tba Client) GetTeamYears(team int) ([]int, error) {
	url := fmt.Sprintf("team/frc%d/years_participated", team)
	var y []int
	err := tba.jsonToStruct(url, &y)
	return y, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetMatch(match string) (Match, error) {
	url := fmt.Sprintf("match/%s", match)
	var m Match
	err := tba.jsonToStruct(url, &m)
	return m, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetEventMatches(event string) ([]Match, error) {
	url := fmt.Sprintf("event/%s/matches", event)
	var m []Match
	err := tba.jsonToStruct(url, &m)
	return m, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetTeamEventMatches(team int, event string) ([]Match, error) {
	url := fmt.Sprintf("team/frc%d/event/%s/matches", team, event)
	var m []Match
	err := tba.jsonToStruct(url, &m)
	return m, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetEventTeams(event string) ([]Team, error) {
	url := fmt.Sprintf("event/%s/teams", event)
	var t []Team
	err := tba.jsonToStruct(url, &t)
	return t, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetEventStats(event string) ([]Stats, error) {
	url := fmt.Sprintf("event/%s/stats", event)
	var s []Stats
	err := tba.jsonToStruct(url, &s)
	return s, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetEventAwards(event string) ([]Award, error) {
	url := fmt.Sprintf("event/%s/awards", event)
	var a []Award
	err := tba.jsonToStruct(url, &a)
	return a, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetEventRankings(event string) ([][]string, error) {
	url := fmt.Sprintf("event/%s/rankings", event)
	var r [][]string
	err := tba.jsonToStruct(url, &r)
	return r, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetEvents(year int) ([]Event, error) {
	url := fmt.Sprintf("events/%d", year)
	var m []Event
	err := tba.jsonToStruct(url, &m)
	return m, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetTeamMedia(team int) ([]Media, error) {
	url := fmt.Sprintf("team/frc%d/media", team)
	var m []Media
	err := tba.jsonToStruct(url, &m)
	return m, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetTeamYearMedia(team int, year int) ([]Media, error) {
	url := fmt.Sprintf("team/frc%d/%d/media", team, year)
	var m []Media
	err := tba.jsonToStruct(url, &m)
	return m, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetTeamHistoryEvents(team int) ([]Event, error) {
	url := fmt.Sprintf("team/frc%d/history/events", team)
	var e []Event
	err := tba.jsonToStruct(url, &e)
	return e, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetTeamHistoryAwards(team int) ([]Award, error) {
	url := fmt.Sprintf("team/frc%d/history/awards", team)
	var a []Award
	err := tba.jsonToStruct(url, &a)
	return a, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetTeamHistoryRobots(team int) ([]Robot, error) {
	url := fmt.Sprintf("team/frc%d/history/robots", team)
	var r []Robot
	err := tba.jsonToStruct(url, &r)
	return r, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetTeamHistoryDistricts(team int) (map[string]string, error) {
	url := fmt.Sprintf("team/frc%d/history/awards", team)
	var d map[string]string
	err := tba.jsonToStruct(url, &d)
	return d, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetDistricts(year int) ([]map[string]string, error) {
	url := fmt.Sprintf("districts/%d", year)
	var d []map[string]string
	err := tba.jsonToStruct(url, &d)
	return d, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetDistrictEvents(district string, year int) ([]Event, error) {
	url := fmt.Sprintf("district/%s/%d/events", district, year)
	var e []Event
	err := tba.jsonToStruct(url, &e)
	return e, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetDistrictRankings(district string, year int) ([]Ranking, error) {
	url := fmt.Sprintf("district/%s/%d/rankings", district, year)
	var r []Ranking
	err := tba.jsonToStruct(url, &r)
	return r, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetDistrictTeams(district string, year int) ([]Team, error) {
	url := fmt.Sprintf("district/%s/%d/teams", district, year)
	var t []Team
	err := tba.jsonToStruct(url, &t)
	return t, err
}

// Deprecated: APIv2 is deprecated
func (tba Client) GetDistrictPoints(event string) (DistrictPoints, error) {
	url := fmt.Sprintf("event/%s/district_points", event)
	var p DistrictPoints
	err := tba.jsonToStruct(url, &p)
	return p, err
}
