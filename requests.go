package tba

import "fmt"

func (b statusBuilder) Get() (Status, error) {
	url := "status"
	var d Status
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b teamsBuilder) Get() ([]Team, error) {
	var url string
	if b.event != "" {
		url = fmt.Sprintf("event/%s/teams", b.event)
	} else if b.district != "" {
		url = fmt.Sprintf("district/%d%s/teams", b.year, b.district)
	} else {
		fmt.Sprintf("teams/%d", b.page)
	}
	var d []Team
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b teamBuilder) Get() (Team, error) {
	var url string
	if b.simple {
		url = fmt.Sprintf("team/frc%d/simple", b.number)
	} else {
		url = fmt.Sprintf("team/frc%d", b.number)
	}
	var d Team
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b eventsBuilder) Get() ([]Event, error) {
	var url string
	fmt.Println(b.team)
	if b.team != 0 {
		url = fmt.Sprintf("team/frc%d/events", b.team)
	} else if b.district != "" {
		url = fmt.Sprintf("district/%d%s/events", b.year, b.district)
	} else {
		url = fmt.Sprintf("events/%d", b.year)
	}
	var d []Event
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b eventBuilder) Get() (Event, error) {
	url := fmt.Sprintf("event/%s", b.id)
	var d Event
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b awardsBuilder) Get() ([]Award, error) {
	var url string
	if b.team != 0 {
		url = fmt.Sprintf("team/frc%d/awards", b.team)
	} else if b.event != "" {
		url = fmt.Sprintf("event/%s/awards", b.event)
	}
	if b.year != 0 {
		url = fmt.Sprintf("awards/%d", b.year)
	}
	var d []Award
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b matchesBuilder) Get() ([]Match, error) {
	var url string
	if b.team != 0 {
		url = fmt.Sprintf("team/frc%d/matches", b.team, url)
		if b.year != 0 {
			url = fmt.Sprintf("%s/%d", url, b.year)
		}
	} else if b.event != "" {
		url = fmt.Sprintf("event/%s/matches", b.event, url)
	}
	var d []Match
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b matchBuilder) Get() (Match, error) {
	url := fmt.Sprintf("match/%s", b.key)
	if b.simple {
		url += "/simple"
	}
	var d Match
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b yearsBuilder) Get() ([]Year, error) {
	url := fmt.Sprintf("team/frc%d/years_participated", b.team)
	var d []Year
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b mediaBuilder) Get() ([]Media, error) {
	url := fmt.Sprintf("team/frc%d/media", b.team)
	if b.year != 0 {
		url = fmt.Sprintf("%s/%d", url, b.year)
	}
	var d []Media
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b robotsBuilder) Get() ([]Robot, error) {
	url := fmt.Sprintf("team/frc%d/robots", b.team)
	var d []Robot
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b districtsBuilder) Get() ([]District, error) {
	var url string
	if b.team != 0 {
		url = fmt.Sprintf("team/frc%d/districts", b.team, url)
	} else {
		url = fmt.Sprintf("districts/%d", url, b.year)
	}
	var d []District
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b profilesBuilder) Get() ([]Profile, error) {
	url := fmt.Sprintf("team/frc%d/social_media", b.team)
	var d []Profile
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b alliancesBuilder) Get() ([]Alliance, error) {
	url := fmt.Sprintf("event/%s/alliances", b.event)
	var d []Alliance
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b districtPointsBuilder) Get() (DistrictPoints, error) {
	url := fmt.Sprintf("event/%s/district_points", b.event)
	var d DistrictPoints
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b insightsBuilder) Get() (Insights, error) {
	url := fmt.Sprintf("event/%s/insights", b.event)
	var d Insights
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b oprsBuilder) Get() (OPRs, error) {
	url := fmt.Sprintf("event/%s/oprs", b.event)
	var d OPRs
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b predictionsBuilder) Get() (Predictions, error) {
	url := fmt.Sprintf("event/%s/predictions", b.event)
	var d Predictions
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b rankingsBuilder) Get() (Rankings, error) {
	url := fmt.Sprintf("event/%s/rankings", b.event)
	var d Rankings
	err := b.client.URLStruct(url, &d)
	return d, err
}

func (b districtRankingsBuilder) Get() ([]DistrictRankings, error) {
	url := fmt.Sprintf("district/%d%s/rankings", b.year, b.abbreviation)
	var d []DistrictRankings
	err := b.client.URLStruct(url, &d)
	return d, err
}
