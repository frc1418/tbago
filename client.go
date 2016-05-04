// tba is a client for interfacing with the The Blue Alliance.
package tba

import (
	"errors"
)

type Client struct {
	developer   string
	application string
	version     string
}

// Init creates a new client for The Blue Alliance.
func Init(developer, application, version string) (client Client, err error) {
	if len(developer) == 1 || len(application) == 1 || len(version) == 1 {
		err = errors.New("Invalid arguments")
		return client, err
	}
	client = Client{
		developer:   developer,
		application: application,
		version:     version,
	}
	return client, err
}

func (tba Client) appID() string {
	return tba.developer + ":" + tba.application + ":" + tba.version
}
