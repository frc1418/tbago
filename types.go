package tba

type Team struct {
	Website     string `json:"website"`      // Official website associated with the team
	Name        string `json:"name"`         // Official long form name registered with FIRST
	Locality    string `json:"locality"`     // City of team derived from parsing the address registered with FIRST
	RookieYear  int    `json:"rookie_year"`  // Team's rookie year
	Region      string `json:"region"`       // State of team derived from parsing the address registered with FIRST
	TeamNumber  int    `json:"team_number"`  // Official team number issued by FIRST
	Location    string `json:"location"`     // Long form address that includes city, state, and country provided by FIRST
	Key         string `json:"key"`          // TBA team key with the format frcyyyy
	CountryName string `json:"country_name"` // Country of team derived from parsing the address registered with FIRST
	Motto       string `json:"motto"`        // Team's Motto
	Nickname    string `json:"nickname"`     // Team nickname provided by FIRST
}

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

type Media []struct {
	Type    string `json:"type"`
	Details struct {
		ImagePartial string `json:"image_partial"`
	} `json:"details"`
	ForeignKey string `json:"foreign_key"`
}
