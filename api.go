package tba

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	TBA_API_URL = "https://www.thebluealliance.com/api/v2/" // Base URL for The Blue Alliance's API
)

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
