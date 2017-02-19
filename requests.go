package tba

import (
	"fmt"
)

// Team list request
func (tba Client) GetTeams(pageNum int) (t []Team, e error) {
	url := fmt.Sprintf("teams/%d", pageNum)
	e = tba.jsonToStruct(url, &t)
	return t, e
}

// Team request
func (tba Client) GetTeam(team int) (t Team, e error) {
	url := fmt.Sprintf("team/frc%d", team)
	e = tba.jsonToStruct(url, &t)
	return t, e
}

// Event request
func (tba Client) GetEvent(key string) (ev Event, e error) {
	url := fmt.Sprintf("event/%s", key)
	e = tba.jsonToStruct(url, &ev)
	return ev, e
}

// Team events request
func (tba Client) GetTeamEvents(team int, year int) (events []Event, e error) {
	url := fmt.Sprintf("team/frc%d/%d/events", team, year)
	e = tba.jsonToStruct(url, &events)
	return events, e
}

// Team awards request
func (tba Client) GetTeamAwards(team int) (awards []Award, e error) {
	url := fmt.Sprintf("team/frc%d/awards", team)
	e = tba.jsonToStruct(url, &events)
	return events, e
}

// Team event awards request
func (tba Client) GetTeamEventAwards(team int, event string) (awards []Award, e error) {
	url := fmt.Sprintf("team/frc%d/event/%s/awards", team, event)
	e = tba.jsonToStruct(url, &events)
	return events, e
}

// Team years active request
func (tba Client) GetTeamYears(team int) (years []int, e error) {
	url := fmt.Sprintf("team/frc%d/years_participated", team)
	e = tba.jsonToStruct(url, &years)
	return years, e
}

func (tba Client) GetMatch(match string) (Match, error) {
	url := fmt.Sprintf("match/%s", match)
	var m Match
	e := tba.jsonToStruct(url, &m)
	return m, e
}

func (tba Client) GetEventMatches(event string) ([]Match, error) {
	url := fmt.Sprintf("event/%s/matches", event)
	var m []Match
	e := tba.jsonToStruct(url, &m)
	return m, e
}

func (tba Client) GetTeamEventMatches(team int, event string) ([]Match, error) {
	url := fmt.Sprintf("team/frc%d/event/%s/matches", team, event)
	var m []Match
	e := tba.jsonToStruct(url, &m)
	return m, e
}

func (tba Client) GetEventTeams(event string) ([]Team, error) {
	url := fmt.Sprintf("event/%s/teams", event)
	var t []Team
	e := tba.jsonToStruct(url, &t)
	return t, e
}

func (tba Client) GetEventStats(event string) ([]Stats, error) {
	url := fmt.Sprintf("event/%s/stats", event)
	var s []Stats
	e := tba.jsonToStruct(url, &s)
	return s, e
}

func (tba Client) GetEventAwards(event string) ([]Award, error) {
	url := fmt.Sprintf("event/%s/awards", event)
	var a []Awards
	e := tba.jsonToStruct(url, &a)
	return a, e
}
func (tba Client) GetEventRankings(event string) ([][]string, error) {
	url := fmt.Sprintf("event/%s/rankings", event)
	var r [][]string
	e := tba.jsonToStruct(url, &r)
	return r, e
}

func (tba Client) GetEvents(year int) ([]Event, error) {
	url := fmt.Sprintf("events/%d", year)
	var m []Event
	e := tba.jsonToStruct(url, &m)
	return m, e
}

func (tba Client) GetTeamMedia(team int) ([]Media, error) {
	url := fmt.Sprintf("team/frc%d/media", team)
	var m []Media
	e := tba.jsonToStruct(url, &m)
	return m, e
}

func (tba Client) GetTeamYearMedia(team int, year int) ([]Media, error) {
	url := fmt.Sprintf("team/frc%d/%d/media", team, year)
	var m []Media
	e := tba.jsonToStruct(url, &m)
	return m, e
}

func (tba Client) GetTeamHistoryEvents(team int) ([]Event, error) {
	url := fmt.Sprintf("team/frc%d/history/events", team)
	var events []Event
	e := tba.jsonToStruct(url, &events)
	return events, e
}

func (tba Client) GetTeamHistoryAwards(team int) ([]Award, error) {
	url := fmt.Sprintf("team/frc%d/history/awards", team)
	var a []Award
	e := tba.jsonToStruct(url, &a)
	return a, e
}

func (tba Client) GetTeamHistoryRobots(team int) ([]Robot, error) {
	url := fmt.Sprintf("team/frc%d/history/robots", team)
	var r []Robot
	e := tba.jsonToStruct(url, &r)
	return r, e
}

func (tba Client) GetTeamHistoryDistricts(team int) (map[string]string, error) {
	url := fmt.Sprintf("team/frc%d/history/awards", team)
	var d map[string]string
	e := tba.jsonToStruct(url, &d)
}

func (tba Client) GetDistricts(year int) ([]map[string]string, error) {
	url := fmt.Sprintf("districts/%d", year)
	var d []map[string]string
	e := tba.jsonToStruct(url, &d)
	return d, e
}

func (tba Client) GetDistrictEvents(district string, year int) ([]Event, error) {
	url := fmt.Sprintf("district/%s/%d/events", district, year)
	var events []Event
	e := tba.jsonToStruct(url, &events)
	return events, e
}

func (tba Client) GetDistrictRankings(district string, year int) ([]Ranking, error) {
	url := fmt.Sprintf("district/%s/%d/rankings", district, year)
	var r []Ranking
	e := tba.jsonToStruct(url, &r)
	return r, e
}

func (tba Client) GetDistrictTeams(district string, year int) ([]Team, error) {
	url := fmt.Sprintf("district/%s/%d/teams", district, year)
	var t []Team
	e := tba.jsonToStruct(url, &t)
	return t, e
}

func (tba Client) GetDistrictPoints(event string) (DistrictPoints, error) {
	url := fmt.Sprintf("event/%s/district_points", event)
	var p DistrictPoints
	e := tba.jsonToStruct(url, &p)
	return p, e
}
