package tba

import (
	"fmt"
)

// Team list request
func (tba Client) GetTeams(pageNum int) ([]Team, error) {
	url := fmt.Sprintf("teams/%d", pageNum)
	var t []Team
	err := tba.jsonToStruct(url, &t)
	return t, err
}

// Team request
func (tba Client) GetTeam(team int) (Team, error) {
	url := fmt.Sprintf("team/frc%d", team)
	var t Team
	err := tba.jsonToStruct(url, &t)
	return t, err
}

// Event request
func (tba Client) GetEvent(key string) (Event, error) {
	url := fmt.Sprintf("event/%s", key)
	var e Event
	err := tba.jsonToStruct(url, &e)
	return e, err
}

// Team events request
func (tba Client) GetTeamEvents(team int, year int) ([]Event, error) {
	url := fmt.Sprintf("team/frc%d/%d/events", team, year)
	var e []Event
	err := tba.jsonToStruct(url, &e)
	return e, err
}

// Team awards request
func (tba Client) GetTeamAwards(team int) ([]Award, error) {
	url := fmt.Sprintf("team/frc%d/awards", team)
	var a []Award
	err := tba.jsonToStruct(url, &a)
	return a, err
}

// Team event awards request
func (tba Client) GetTeamEventAwards(team int, event string) ([]Award, error) {
	url := fmt.Sprintf("team/frc%d/event/%s/awards", team, event)
	var a []Award
	err := tba.jsonToStruct(url, &a)
	return a, err
}


// Team years active request
func (tba Client) GetTeamYears(team int) ([]int, error) {
	url := fmt.Sprintf("team/frc%d/years_participated", team)
	var y []int
	err := tba.jsonToStruct(url, &y)
	return y, err
}

func (tba Client) GetMatch(match string) (Match, error) {
	url := fmt.Sprintf("match/%s", match)
	var m Match
	err := tba.jsonToStruct(url, &m)
	return m, err
}

func (tba Client) GetEventMatches(event string) ([]Match, error) {
	url := fmt.Sprintf("event/%s/matches", event)
	var m []Match
	err := tba.jsonToStruct(url, &m)
	return m, err
}

func (tba Client) GetTeamEventMatches(team int, event string) ([]Match, error) {
	url := fmt.Sprintf("team/frc%d/event/%s/matches", team, event)
	var m []Match
	err := tba.jsonToStruct(url, &m)
	return m, err
}

func (tba Client) GetEventTeams(event string) ([]Team, error) {
	url := fmt.Sprintf("event/%s/teams", event)
	var t []Team
	err := tba.jsonToStruct(url, &t)
	return t, err
}

func (tba Client) GetEventStats(event string) ([]Stats, error) {
	url := fmt.Sprintf("event/%s/stats", event)
	var s []Stats
	err := tba.jsonToStruct(url, &s)
	return s, err
}

func (tba Client) GetEventAwards(event string) ([]Award, error) {
	url := fmt.Sprintf("event/%s/awards", event)
	var a []Award
	err := tba.jsonToStruct(url, &a)
	return a, err
}

func (tba Client) GetEventRankings(event string) ([][]string, error) {
	url := fmt.Sprintf("event/%s/rankings", event)
	var r [][]string
	err := tba.jsonToStruct(url, &r)
	return r, err
}

func (tba Client) GetEvents(year int) ([]Event, error) {
	url := fmt.Sprintf("events/%d", year)
	var m []Event
	err := tba.jsonToStruct(url, &m)
	return m, err
}

func (tba Client) GetTeamMedia(team int) ([]Media, error) {
	url := fmt.Sprintf("team/frc%d/media", team)
	var m []Media
	err := tba.jsonToStruct(url, &m)
	return m, err
}

func (tba Client) GetTeamYearMedia(team int, year int) ([]Media, error) {
	url := fmt.Sprintf("team/frc%d/%d/media", team, year)
	var m []Media
	err := tba.jsonToStruct(url, &m)
	return m, err
}

func (tba Client) GetTeamHistoryEvents(team int) ([]Event, error) {
	url := fmt.Sprintf("team/frc%d/history/events", team)
	var e []Event
	err := tba.jsonToStruct(url, &e)
	return e, err
}

func (tba Client) GetTeamHistoryAwards(team int) ([]Award, error) {
	url := fmt.Sprintf("team/frc%d/history/awards", team)
	var a []Award
	err := tba.jsonToStruct(url, &a)
	return a, err
}

func (tba Client) GetTeamHistoryRobots(team int) ([]Robot, error) {
	url := fmt.Sprintf("team/frc%d/history/robots", team)
	var r []Robot
	err := tba.jsonToStruct(url, &r)
	return r, err
}

func (tba Client) GetTeamHistoryDistricts(team int) (map[string]string, error) {
	url := fmt.Sprintf("team/frc%d/history/awards", team)
	var d map[string]string
	err := tba.jsonToStruct(url, &d)
	return d, err
}

func (tba Client) GetDistricts(year int) ([]map[string]string, error) {
	url := fmt.Sprintf("districts/%d", year)
	var d []map[string]string
	err := tba.jsonToStruct(url, &d)
	return d, err
}

func (tba Client) GetDistrictEvents(district string, year int) ([]Event, error) {
	url := fmt.Sprintf("district/%s/%d/events", district, year)
	var e []Event
	err := tba.jsonToStruct(url, &e)
	return e, err
}

func (tba Client) GetDistrictRankings(district string, year int) ([]Ranking, error) {
	url := fmt.Sprintf("district/%s/%d/rankings", district, year)
	var r []Ranking
	err := tba.jsonToStruct(url, &r)
	return r, err
}

func (tba Client) GetDistrictTeams(district string, year int) ([]Team, error) {
	url := fmt.Sprintf("district/%s/%d/teams", district, year)
	var t []Team
	err := tba.jsonToStruct(url, &t)
	return t, err
}

func (tba Client) GetDistrictPoints(event string) (DistrictPoints, error) {
	url := fmt.Sprintf("event/%s/district_points", event)
	var p DistrictPoints
	err := tba.jsonToStruct(url, &p)
	return p, err
}
