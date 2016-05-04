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

// Team events request
func (tba Client) GetTeamEvents(team, year int) (events []Event, e error) {
	url := fmt.Sprintf("team/frc%d/%d/events", team, year)
	e = tba.jsonToStruct(url, &events)
	return events, e
}

func (tba Client) GetMatch(match string) (Match, error) {
	url := "match/" + match
	var m Match
	err := tba.jsonToStruct(url, &m)
	return m, err
}

func (tba Client) GetEventMatches(event string) ([]Match, error) {
	url := "event/" + event + "/matches"
	var m []Match
	err := tba.jsonToStruct(url, &m)
	return m, err
}

func (tba Client) GetEvents(year int) ([]Event, error) {
	url := fmt.Sprintf("events/%d", year)
	var m []Event
	err := tba.jsonToStruct(url, &m)
	return m, err
}
