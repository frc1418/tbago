package tba

/*
/api/v3/teams/{page}                                - 2000
/api/v3/team/{team_key}                             - 0100
/api/v3/team/{team_key}/events/{year}               - 0910
/api/v3/team/{team_key}/event/{event_key}/awards    - 0301
/api/v3/team/{team_key}/event/{event_key}/matches   - 0308
/api/v3/team/{team_key}/event/{event_key}/status    - 030F
/api/v3/team/{team_key}/years_participated          - 0109
/api/v3/team/{team_key}/media/{year}                - 0920
/api/v3/team/{team_key}/social_media                - 010A
/api/v3/team/{team_key}/robots                      - 010B
/api/v3/team/{team_key}/districts                   - 010C
/api/v3/events/{year}                               - 0810
/api/v3/event/{event_key}                           - 0200
/api/v3/event/{event_key}/teams                     - 0203
/api/v3/event/{event_key}/matches                   - 0208
/api/v3/event/{event_key}/oprs                      - 020D
/api/v3/event/{event_key}/insights                  - 0206
/api/v3/event/{event_key}/rankings                  - 0202
/api/v3/event/{event_key}/awards                    - 0201
/api/v3/event/{event_key}/alliances                 - 0205
/api/v3/event/{event_key}/district_points           - 020E
/api/v3/match/{match_key}                           - 1000
/api/v3/districts/{year}                            - 0800
/api/v3/district/{district_key}/events              - 0410
/api/v3/district/{district_key}/rankings            - 0402
/api/v3/district/{district_key}/teams               - 0403
*/
import (
	"fmt"
	apiclient "github.com/carlcolglazier/tba/client"
	"github.com/carlcolglazier/tba/client/operations"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"log"
)

type Client struct {
	key     string
	version int
}

const (
	AWARDS             = 0x1
	RANKINGS           = 0x2
	TEAMS              = 0x3
	EVENTS             = 0x4
	ALLIANCES          = 0x5
	INSIGHTS           = 0x6
	MATCHES            = 0x8
	YEARS_PARTICIPATED = 0x9
	SOCIAL_MEDIA       = 0xA
	ROBOTS             = 0xB
	DISTRICTS          = 0xC
	OPRS               = 0xD
	DISTRICT_POINTS    = 0xE
	STATUS             = 0xF
	EVENTS             = 0x10
	MEDIA              = 0x20
	TEAM_KEY           = 0x0100
	EVENT_KEY          = 0x0200
	DISTRICT_KEY       = 0x0400
	YEAR               = 0x0800
	MATCH_KEY          = 0x1000
	PAGE               = 0x2000
)

type RequestBuilder struct {
	client       *Client
	page         int
	team_key     string
	event_key    string
	year         int
	match_key    string
	district_key string
	status       int
}

func New(key string, api_version int) *Client {
	if api_version == 3 {
		return &Client{key, api_version}
	}
}

func (c *Client) Builder() *RequestBuilder {
	return &RequestBuilder{c, 0, "", "", 0, "", "", 0}
}

func (rb *RequestBuilder) Teams(page int) {
	rb.page = page
	rb.status += PAGE
}

func (rb *RequestBuilder) Team(team string) {
	rb.team_key = team
	rb.status += TEAM_KEY
}

func (rb *RequestBuilder) Event(event string) {
	rb.event_key = event
	rb.status += EVENT_KEY
}

func (rb *RequestBuilder) District(district string) {
	rb.district_key = district
	rb.status += DISTRICT_KEY
}

func (rb *RequestBuilder) Year(year int) {
	rb.year = year
	rb.status += YEAR
}
