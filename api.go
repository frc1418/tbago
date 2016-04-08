package tba

import (
	//"encoding/json"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	TBA_API_URL  = "https://thebluealliance.com/api/v2/" // Base URL for The Blue Alliance's API
	APP_ID_QUERY = "?X-TBA-App-Id="                      // Used for constructing the URL string
)

// sendRequest sends a HTTP GET request to The Blue Alliance API.
// An HTTP response is returned.
func (tba Client) sendRequest(url string) (*http.Response, error) {
	// Golang uses headers by default that follow the HTTP spec.
	// The Blue Alliance, on the other hand, is case sensitive when
	// checking headers, which means that in this case the "X-TBA-App-ID"
	// header does not work. Instead, the app ID is sent through the URL.
	requestUrl := TBA_API_URL + url + APP_ID_QUERY + tba.appID()
	return http.Get(requestUrl)
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

func (tba Client) GetTeam(team int) (t Team, e error) {
	url := "team/frc" + strconv.Itoa(team)
	e = tba.jsonToStruct(url, &t)
	return t, e
}

func (tba *Client) GetMatch(match string) (Match, error) {
	url := "match/" + match
	var m Match
	err := tba.jsonToStruct(url, &m)
	return m, err
}

func (tba *Client) GetEventMatches(event string) ([]Match, error) {
	url := "event/" + event + "/matches"
	var m []Match
	err := tba.jsonToStruct(url, &m)
	return m, err
}

func (tba *Client) GetEvents(year int) ([]Event, error) {
	url := "events/" + strconv.Itoa(year)
	var m []Event
	err := tba.jsonToStruct(url, &m)
	return m, err
}
