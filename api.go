package tba

import (
    "encoding/json"
    "errors"
    "io/ioutil"
    "net/http"
)

const API_ROOT = "https://www.thebluealliance.com/api/v3/"

func (tba Client) req(method string, path string) (*http.Response, error) {
    url := API_ROOT + path
    client := &http.Client{}
    req, err := http.NewRequest(method, url, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Add("X-TBA-Auth-Key", tba.key)
    return client.Do(req)
}

func (tba Client) get(path string) ([]byte, error) {
    resp, err := tba.req("GET", path)
    if err != nil {
        return nil, err
    }
    if resp.StatusCode != 200 {
        return nil, errors.New("Response code not 200")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    return body, err
}

func (tba Client) URLStruct(url string, v interface{}) error {
    resp, err := tba.get(url)
    if err != nil {
        return err
    }
    err = json.Unmarshal(resp, v)
    return err
}
