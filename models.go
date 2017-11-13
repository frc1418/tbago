package tba

type statusBuilder struct {
	client *Client
}
type Status struct {
	Android struct {
		LatestAppVersion int `json:"latest_app_version"`
		MinAppVersion    int `json:"min_app_version"`
	} `json:"android"`
	iOS struct {
		LatestAppVersion int `json:"latest_app_version"`
		MinAppVersion    int `json:"min_app_version"`
	} `json:"ios"`
	//Backup              map[string]interface{}  `json:"backup"`
	ContBuildEnabled bool     `json:"contbuild_enabled"`
	CurrentSeason    int      `json:"current_season"`
	MaxSeason        int      `json:"max_season"`
	DownEvents       []string `json:"down_events"`
	IsDatafeedDown   bool     `json:"is_datafeed_down"`
}

type teamsBuilder struct {
	page     int
	event    string
	district string
	year     int
	client   *Client
}

type teamBuilder struct {
	number int
	simple bool
	client *Client
}
type Team struct {
	Address          string         `json:"address"`
	City             string         `json:"city"`
	StateProv        string         `json:"state_prov"`
	Country          string         `json:"country"`
	PostalCode       string         `json:"rookie_year"`
	GMapsPlaceID     string         `json:"gmaps_place_id"`
	GMapsURL         string         `json:"gmaps_url"`
	HomeChampionship map[int]string `json:"home_championship"`
	Key              string         `json:"key"`
	Latitute         string         `json:"lat"`
	Longitude        string         `json:"lng"`
	Location         string         `json:"location_name"`
	Motto            string         `json:"motto"`
	Name             string         `json:"name"`
	Nickname         string         `json:"nickname"`
	RookieYear       int            `json:"rookie_year"`
	Number           int            `json:"team_number"`
	Website          string         `json:"website"`
}

type eventsBuilder struct {
	year     int
	team     int
	district string
	simple   bool
	client   *Client
}

type eventBuilder struct {
	id     string
	simple bool
	client *Client
}

type Event struct {
	Address  string `json:"address"`
	City     string `json:"city"`
	Country  string `json:"country"`
	District struct {
		Abbreviation string `json:"abbreviation"`
		DisplayName  string `json:"display_name"`
		Key          string `json:"key"`
		Year         int    `json:"year"`
	} `json:"district"`
	DivisionKeys      []string `json:"division_keys"`
	EventCode         string   `json:"event_code"`
	EventType         int      `json:"event_type"`
	EventTypeString   string   `json:"event_type_string"`
	FIRSTEventCode    string   `json:"first_event_code"`
	FIRSTEventID      string   `json:"first_event_id"`
	GMapsPlaceID      string   `json:"gmaps_place_id"`
	GMapsURL          string   `json:"gmaps_url"`
	Key               string   `json:"key"`
	Latitute          string   `json:"lat"`
	Longitude         string   `json:"lng"`
	LocationName      string   `json:"location_name"`
	Name              string   `json:"name"`
	ParentEventKey    string   `json:"parent_event_key"`
	PlayoffType       int      `json:"playoff_type"`
	PlayoffTypeString int      `json:"playoff_type_string"`
	PostalCode        string   `json:"postal_code"`
	ShortName         string   `json:"short_name"`
	StartDate         string   `json:"start_date"`
	StateProv         string   `json:"state_prov"`
	Timezone          string   `json:"timezone"`
	Webcasts          []struct {
		Channel string `json:"channel"`
		Type    string `json:"type"`
	} `json:"webcasts"`
	Website string `json:"website"`
	Week    int    `json:"week"`
}

type awardsBuilder struct {
	year   int
	team   int
	event  string
	client *Client
}
type Award struct {
	AwardType     int    `json:"award_type"`
	EventKey      string `json:"event_key"`
	Name          string `json:"name"`
	RecipientList []struct {
		TeamNumber int         `json:"team_number"`
		Awardee    interface{} `json:"awardee"`
	} `json:"recipient_list"`
	Year int `json:"year"`
}

type matchesBuilder struct {
	team   int
	event  string
	year   int
	client *Client
}
type matchBuilder struct {
	key    string
	simple bool
	client *Client
}
type Match struct {
	ActualTime int `json:"actual_time"`
	Alliances  struct {
		Blue struct {
			DQTeams        []string `json:"dq_team_keys"`
			Score          int      `json:"score"`
			SurrogateTeams []string `json:"surrogate_team_keys"`
			Teams          []string `json:"teams"`
		} `json:"blue"`
		Red struct {
			DQTeams        []string `json:"dq_team_keys"`
			Score          int      `json:"score"`
			SurrogateTeams []string `json:"surrogate_team_keys"`
			Teams          []string `json:"teams"`
		} `json:"red"`
	} `json:"alliances"`
	CompLevel      string `json:"comp_level"`
	EventKey       string `json:"event_key"`
	Key            string `json:"key"`
	MatchNumber    int    `json:"match_number"`
	PostResultTime int    `json:"post_result_time"`
	PredictedTime  int    `json:"predicted_time"`
	ScoreBreakdown struct {
		Red  interface{} `json:"red"`
		Blue interface{} `json:"blue"`
	} `json:"score_breakdown"`
	SetNumber int `json:"set_number"`
	Time      int `json:"time"`
	Videos    []struct {
		Type string `json:"type"`
		Key  string `json:"key"`
	} `json:"videos"`
	WinningAlliance string `json:"winning_alliance"`
}

type yearsBuilder struct {
	team   int
	client *Client
}
type Year int

type mediaBuilder struct {
	team   int
	year   int
	client *Client
}
type Media struct {
	Type       string `json:"type"`
	ForeignKey string `json:"foreign_key"`
	Preferred  bool   `json:"preferrred"`
	Details    struct {
		// Instagram
		AuthorID        int    `json:"author_id"`
		AuthorName      string `json:"author_name"`
		AuthorURL       string `json:"author_url"`
		Height          int    `json:"height"`
		HTML            string `json:"html"`
		MediaID         string `json:"media_id"`
		ProviderName    string `json:"provider_name"`
		ProviderURL     string `json:"provider_url"`
		ThumbnailHeight string `json:"thumbnail_height"`
		ThumbnailURL    string `json:"thumbnail_url"`
		Title           string `json:"thumbnail_width"`
		Type            string `json:"type"`
		Version         string `json:"version"`
		Width           int    `json:"width"`
	} `json:"details"`
}

type robotsBuilder struct {
	team   int
	client *Client
}
type Robot struct {
	Key     string `json:"key"`
	Name    string `json:"robot_name"`
	TeamKey string `json:"team_key"`
	Year    int    `json:"year"`
}

type districtsBuilder struct {
	team         int
	abbreviation string
	year         int
	client       *Client
}
type districtBuilder struct {
	abbreviation string
	year         int
	client       *Client
}
type District struct {
	Abbreviation string `json:"abbreviation"`
	DisplayName  string `json:"display_name"`
	Key          string `json:"key"`
	Year         int    `json:"year"`
}

type profilesBuilder struct {
	team   int
	client *Client
}
type Profile struct {
	//Details     interface{}   `json:"details"`
	ForeignKey string `json:"foreign_key"`
	Preferred  bool   `json:"preferred"`
	Type       string `json:"type"`
}

type alliancesBuilder struct {
	event  string
	client *Client
}
type Alliance struct {
	Backup   interface{} `json:"backup"`
	Declines []string    `json:"declines"`
	Name     string      `json:"name"`
	Picks    []string    `json:"picks"`
	Status   struct {
		CurrentLevelRecord struct {
			Losses int `json:"losses"`
			Ties   int `json:"ties"`
			Wins   int `json:"wins"`
		} `json:"current_level_record"`
		Level          string `json:"level"`
		PlayoffAverage int    `json:"playoff_average"`
		Record         struct {
			Losses int `json:"losses"`
			Ties   int `json:"ties"`
			Wins   int `json:"wins"`
		} `json:"record"`
		Status string `json:"status"`
	} `json:"status"`
}

type districtPointsBuilder struct {
	event  string
	client *Client
}
type DistrictPoints struct {
	Points map[string]struct {
		AlliancePoints int `json:"alliance_points"`
		AwardPoints    int `json:"award_points"`
		ElimPoints     int `json:"elim_points"`
		QualPoints     int `json:"qual_points"`
		Total          int `json:"total"`
	} `json:"points"`
	Tiebreakers map[string]struct {
		HighestQualScores []int `json:"highest_qual_scores"`
		QualWins          int   `json:"qual_wins"`
	} `json:"tiebreakers"`
}

type insightsBuilder struct {
	event  string
	client *Client
}
type Insights struct {
	// TODO
}

type oprsBuilder struct {
	event  string
	client *Client
}
type OPRs struct {
	CCWMs map[string]float64 `json:"ccwms"`
	DPRs  map[string]float64 `json:"dprs"`
	OPRs  map[string]float64 `json:"oprs"`
}

type predictionsBuilder struct {
	event  string
	client *Client
}
type Predictions struct {
	MatchPredictionStats struct {
		Playoff struct {
			BrierScores  map[string]float64 `json:"brier_scores"`
			ErrMean      float64            `json:"err_mean"`
			ErrVar       float64            `json:"err_var"`
			WLAccuracy   float64            `json:"wl_accuracy"`
			WLAccuracy75 float64            `json:"wl_accuracy_75"`
		} `json:"playoff"`
		Qual struct {
			BrierScores  map[string]float64 `json:"brier_scores"`
			ErrMean      float64            `json:"err_mean"`
			ErrVar       float64            `json:"err_var"`
			WLAccuracy   float64            `json:"wl_accuracy"`
			WLAccuracy75 float64            `json:"wl_accuracy_75"`
		} `json:"qual"`
	} `json:"match_prediction_stats"`
	MatchPredictions struct {
		Playoff map[string]struct {
			Blue            map[string]float64 `json:"blue"`
			Red             map[string]float64 `json:"red"`
			Prob            float64            `json:"prob"`
			WinningAlliance string             `json:"winning_alliance"`
		} `json:"playoff"`
		Qual map[string]struct {
			Blue            map[string]float64 `json:"blue"`
			Red             map[string]float64 `json:"red"`
			Prob            float64            `json:"prob"`
			WinningAlliance string             `json:"winning_alliance"`
		} `json:"playoff"`
	} `json:"match_prediction_stats"`
	RankingPredictionStats struct {
		// TODO: There are probably other fields here.
		LastPlayedMatch string `json:"last_played_match"`
	} `json:"ranking_prediction_stats"`
	// Ranking prediction data is returned as a list of lists containing a team key and a list of mixed ints and floats. TODO: Figure out how to handle this better.
	RankingPredictions [][]interface{} `json:"ranking_predictions"`
	StatMeanVars       struct {
		Playoff map[string]struct {
			Mean map[string]float64 `json:"mean"`
			Var  map[string]float64 `json:"var"`
		} `json:"playoff"`
		Qual map[string]struct {
			Mean map[string]float64 `json:"mean"`
			Var  map[string]float64 `json:"var"`
		} `json:"playoff"`
	} `json:"stat_mean_vars"`
}

type rankingsBuilder struct {
	event  string
	client *Client
}
type Rankings struct {
	ExtraStatsInfo []struct {
		Name      string `json:"name"`
		Precision string `json:"precision"`
	} `json:"extra_stats_info"`
	SortOrderInfo []struct {
		Name      string `json:"name"`
		Precision string `json:"precision"`
	} `json:"sort_order_info"`
	Rankings []struct {
		DQ            int   `json:"dq"`
		ExtraStats    []int `json:"extra_stats"`
		MatchesPlayes int   `json:"matches_played"`
		QualAverage   int   `json:"qual_average"`
		Rank          int   `json:"rank"`
		Record        struct {
			Losses int `json:"losses"`
			Ties   int `json:"ties"`
			Wins   int `json:"wins"`
		} `json:"record"`
		SortOrders []float64 `json:"sort_orders"`
		TeamKey    string    `json:"team_key"`
	} `json:"rankings"`
}

type districtRankingsBuilder struct {
	abbreviation string
	year         int
	client       *Client
}
type DistrictRankings struct {
	EventPoints []struct {
		AlliancePoints int    `json:"alliance_points"`
		AwardPoints    int    `json:"award_points"`
		DistrictCMP    bool   `json:"district_cmp"`
		ElimPoints     int    `json:"elim_points"`
		EventKey       string `json:"event_key"`
		QualPoints     int    `json:"qual_points"`
		Total          int    `json:"total"`
	} `json:"event_points"`
	PointTotal  int    `json:"point_total"`
	Rank        int    `json:"rank"`
	RookieBonus int    `json:"rookie_bonus"`
	TeamKey     string `json:"team_key"`
}
