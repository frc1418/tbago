package tbago

type tbaStatusBuilder struct {
	client *Client
}
type TBAStatus struct {
	CurrentSeason    int      `json:"current_season"`
	MaxSeason        int      `json:"max_season"`
	IsDatafeedDown   bool     `json:"is_datafeed_down"`
	ContBuildEnabled bool     `json:"contbuild_enabled"`
	DownEvents       []string `json:"down_events"`
	Android          struct {
		LatestAppVersion int `json:"latest_app_version"`
		MinAppVersion    int `json:"min_app_version"`
	} `json:"android"`
	IOS struct {
		LatestAppVersion int `json:"latest_app_version"`
		MinAppVersion    int `json:"min_app_version"`
	} `json:"ios"`
}

type teamsBuilder struct {
	page     int
	event    string
	district string
	year     int
	simple   bool
	client   *Client
}

type teamBuilder struct {
	number int
	event  string
	simple bool
	client *Client
}
type Team struct {
	Key              string         `json:"key"`
	TeamNumber       int            `json:"team_number"`
	Nickname         string         `json:"nickname"`
	Name             string         `json:"name"`
	City             string         `json:"city"`
	StateProv        string         `json:"state_prov"`
	Country          string         `json:"country"`
	Address          string         `json:"address"`
	PostalCode       string         `json:"postal_code"`
	GMapsPlaceID     string         `json:"gmaps_place_id"`
	GMapsURL         string         `json:"gmaps_url"`
	Latitude         float64        `json:"lat"`
	Longitude        float64        `json:"lng"`
	LocationName     string         `json:"location_name"`
	Website          string         `json:"website"`
	Motto            string         `json:"motto"`
	RookieYear       int            `json:"rookie_year"`
	HomeChampionship map[int]string `json:"home_championship"`
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
	team   int
	simple bool
	client *Client
}

type Event struct {
	Key             string   `json:"key"`
	Name            string   `json:"name"`
	EventCode       string   `json:"event_code"`
	EventType       int      `json:"event_type"`
	District        District `json:"district"`
	City            string   `json:"city"`
	StateProv       string   `json:"state_prov"`
	PostalCode      string   `json:"postal_code"`
	Country         string   `json:"country"`
	Address         string   `json:"address"`
	StartDate       string   `json:"start_date"`
	EndDate         string   `json:"end_date"`
	Year            int      `json:"year"`
	ShortName       string   `json:"short_name"`
	EventTypeString string   `json:"event_type_string"`
	Week            int      `json:"week"`
	GMapsPlaceID    string   `json:"gmaps_place_id"`
	GMapsURL        string   `json:"gmaps_url"`
	Latitude        float64  `json:"lat"`
	Longitude       float64  `json:"lng"`
	LocationName    string   `json:"location_name"`
	Timezone        string   `json:"timezone"`
	DivisionKeys    []string `json:"division_keys"`
	Website         string   `json:"website"`
	FIRSTEventID    string   `json:"first_event_id"`
	FIRSTEventCode  string   `json:"first_event_code"`
	Webcasts        []struct {
		Channel string `json:"channel"`
		Type    string `json:"type"`
	} `json:"webcasts"`
	ParentEventKey    string `json:"parent_event_key"`
	PlayoffType       int    `json:"playoff_type"`
	PlayoffTypeString string `json:"playoff_type_string"`
}

type awardsBuilder struct {
	year   int
	team   int
	event  string
	client *Client
}
type Award struct {
	Name          string `json:"name"`
	AwardType     int    `json:"award_type"`
	EventKey      string `json:"event_key"`
	RecipientList []struct {
		TeamKey string `json:"team_key"`
		Awardee string `json:"awardee"`
	} `json:"recipient_list"`
	Year int `json:"year"`
}

type statusBuilder struct {
	team   int
	event  string
	client *Client
}
type Status struct {
	Alliance struct {
		Backup struct {
			Out string `json:"out"`
			In  string `json:"in"`
		} `json:"backup"`
		Name   string `json:"name"`
		Number int    `json:"number"`
		Pick   int    `json:"pick"`
	} `json:"alliance"`
	AllianceStatusStr string `json:"alliance_status_str"`
	LastMatchKey      string `json:"last_match_key"`
	NextMatchKey      string `json:"next_match_key"`
	OverallStatusStr  string `json:"overall_status_str"`
	Playoff           struct {
		Level              string    `json:"level"`
		CurrentLevelRecord WLTRecord `json:"current_level_record"`
		Record             WLTRecord `json:"record"`
		PlayoffAverage     int       `json:"playoff_average"`
		Status             string    `json:"status"`
	} `json:"playoff"`
	PlayoffStatusStr string `json:"playoff_status_str"`
	Qual             struct {
		NumTeams int `json:"num_teams"`
		Ranking  struct {
			DQ            int       `json:"dq"`
			MatchesPlayed int       `json:"matches_played"`
			QualAverage   int       `json:"qual_average"`
			Rank          int       `json:"rank"`
			Record        WLTRecord `json:"record"`
			SortOrders    []float64 `json:"sort_orders"`
			TeamKey       string    `json:"team_key"`
			SortOrderInfo []struct {
				Name      string `json:"name"`
				Precision string `json:"precision"`
			} `json:"sort_order_info"`
			Status string `json:"status"`
		} `json:"ranking"`
	} `json:"qual"`
}

type matchesBuilder struct {
	team   int
	event  string
	year   int
	simple bool
	client *Client
}
type matchBuilder struct {
	key    string
	simple bool
	client *Client
}
type Match struct {
	Key         string `json:"key"`
	CompLevel   string `json:"comp_level"`
	SetNumber   int    `json:"set_number"`
	MatchNumber int    `json:"match_number"`
	Alliances   struct {
		Blue struct {
			DQTeams        []string `json:"dq_team_keys"`
			Score          int      `json:"score"`
			SurrogateTeams []string `json:"surrogate_team_keys"`
			TeamKeys       []string `json:"team_keys"`
		} `json:"blue"`
		Red struct {
			DQTeams        []string `json:"dq_team_keys"`
			Score          int      `json:"score"`
			SurrogateTeams []string `json:"surrogate_team_keys"`
			TeamKeys       []string `json:"team_keys"`
		} `json:"red"`
	} `json:"alliances"`
	WinningAlliance string `json:"winning_alliance"`
	EventKey        string `json:"event_key"`
	Time            int64  `json:"time"`
	ActualTime      int64  `json:"actual_time"`
	PredictedTime   int64  `json:"predicted_time"`
	PostResultTime  int64  `json:"post_result_time"`
	ScoreBreakdown  struct {
		Red  interface{} `json:"red"`
		Blue interface{} `json:"blue"`
	} `json:"score_breakdown"`
	Videos []struct {
		Type string `json:"type"`
		Key  string `json:"key"`
	} `json:"videos"`
}

type yearsBuilder struct {
	team   int
	client *Client
}
type Year int

type mediaBuilder struct {
	team   int
	year   int
	tag    string
	client *Client
}
type Media struct {
	Type       string                 `json:"type"`
	Key        int                    `json:"key"`
	ForeignKey string                 `json:"foreign_key"`
	Preferred  bool                   `json:"preferrred"`
	Details    map[string]interface{} `json:"details"`
}

type robotsBuilder struct {
	team   int
	client *Client
}
type Robot struct {
	Year    int    `json:"year"`
	Name    string `json:"robot_name"`
	Key     string `json:"key"`
	TeamKey string `json:"team_key"`
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
		CurrentLevelRecord WLTRecord `json:"current_level_record"`
		Level              string    `json:"level"`
		PlayoffAverage     int       `json:"playoff_average"`
		Record             WLTRecord `json:"record"`
		Status             string    `json:"status"`
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
		QualPoints     int `json:"qual_points"`
		ElimPoints     int `json:"elim_points"`
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
	Qual    map[string]interface{} `json:"qual"`
	Playoff map[string]interface{} `json:"playoff"`
}

type oprsBuilder struct {
	event  string
	client *Client
}
type OPRs struct {
	OPRs  map[string]float64 `json:"oprs"`
	DPRs  map[string]float64 `json:"dprs"`
	CCWMs map[string]float64 `json:"ccwms"`
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
		} `json:"qual"`
	} `json:"match_predictions"`
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
		} `json:"qual"`
	} `json:"stat_mean_vars"`
}

type rankingsBuilder struct {
	event  string
	client *Client
}
type Rankings struct {
	Rankings []struct {
		DQ            int       `json:"dq"`
		MatchesPlayed int       `json:"matches_played"`
		QualAverage   int       `json:"qual_average"`
		Rank          int       `json:"rank"`
		Record        WLTRecord `json:"record"`
		ExtraStats    []int     `json:"extra_stats"`
		SortOrders    []float64 `json:"sort_orders"`
		TeamKey       string    `json:"team_key"`
	} `json:"rankings"`
	ExtraStatsInfo []struct {
		Name      string `json:"name"`
		Precision string `json:"precision"`
	} `json:"extra_stats_info"`
	SortOrderInfo []struct {
		Name      string `json:"name"`
		Precision string `json:"precision"`
	} `json:"sort_order_info"`
}

type districtRankingsBuilder struct {
	abbreviation string
	year         int
	client       *Client
}
type DistrictRankings struct {
	TeamKey     string `json:"team_key"`
	Rank        int    `json:"rank"`
	RookieBonus int    `json:"rookie_bonus"`
	PointTotal  int    `json:"point_total"`
	EventPoints []struct {
		AlliancePoints int    `json:"alliance_points"`
		AwardPoints    int    `json:"award_points"`
		DistrictCMP    bool   `json:"district_cmp"`
		ElimPoints     int    `json:"elim_points"`
		EventKey       string `json:"event_key"`
		QualPoints     int    `json:"qual_points"`
		Total          int    `json:"total"`
	} `json:"event_points"`
}

type WLTRecord struct {
	Losses int `json:"losses"`
	Ties   int `json:"ties"`
	Wins   int `json:"wins"`
}
