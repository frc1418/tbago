package tbago

// TBA Status
func (tba Client) Status() *tbaStatusBuilder {
	return &tbaStatusBuilder{
		client: &tba,
	}
}

// Team + modifiers
func (tba Client) Teams(page int) *teamsBuilder {
	return &teamsBuilder{
		page:   page,
		client: &tba,
	}
}
func (tba Client) Team(number int) *teamBuilder {
	return &teamBuilder{
		number: number,
		client: &tba,
	}
}
func (b *teamBuilder) Simple() *teamBuilder {
	b.simple = true
	return b
}
func (b *teamBuilder) Event(event string) *teamBuilder {
	b.event = event
	return b
}
func (b *teamBuilder) Events() *eventsBuilder {
	return &eventsBuilder{
		team:   b.number,
		client: b.client,
	}
}
func (b *teamBuilder) Awards() *awardsBuilder {
	return &awardsBuilder{
		team:   b.number,
		event:  b.event,
		client: b.client,
	}
}
func (b *teamBuilder) Matches() *matchesBuilder {
	return &matchesBuilder{
		team:   b.number,
		event:  b.event,
		client: b.client,
	}
}
func (b *teamBuilder) Years() *yearsBuilder {
	return &yearsBuilder{
		team:   b.number,
		client: b.client,
	}
}
func (b *teamBuilder) Media() *mediaBuilder {
	return &mediaBuilder{
		team:   b.number,
		client: b.client,
	}
}
func (b *teamBuilder) Robots() *robotsBuilder {
	return &robotsBuilder{
		team:   b.number,
		client: b.client,
	}
}
func (b *teamBuilder) Districts() *districtsBuilder {
	return &districtsBuilder{
		team:   b.number,
		client: b.client,
	}
}
func (b *teamBuilder) Profiles() *profilesBuilder {
	return &profilesBuilder{
		team:   b.number,
		client: b.client,
	}
}
func (b *teamBuilder) Status() *statusBuilder {
	return &statusBuilder{
		team:   b.number,
		event:  b.event,
		client: b.client,
	}
}

// Event + modifiers
func (tba Client) Events(year int) *eventsBuilder {
	return &eventsBuilder{
		year:   year,
		client: &tba,
	}
}
func (tba Client) Event(id string) *eventBuilder {
	return &eventBuilder{
		id:     id,
		client: &tba,
	}
}
func (b *eventBuilder) Simple() *eventBuilder {
	b.simple = true
	return b
}
func (b *eventBuilder) Team(team int) *eventBuilder {
	b.team = team
	return b
}
func (b *eventBuilder) Alliances() *alliancesBuilder {
	return &alliancesBuilder{
		event:  b.id,
		client: b.client,
	}
}
func (b *eventBuilder) DistrictPoints() *districtPointsBuilder {
	return &districtPointsBuilder{
		event:  b.id,
		client: b.client,
	}
}
func (b *eventBuilder) Insights() *insightsBuilder {
	return &insightsBuilder{
		event:  b.id,
		client: b.client,
	}
}
func (b *eventBuilder) Teams() *teamsBuilder {
	return &teamsBuilder{
		event:  b.id,
		client: b.client,
	}
}
func (b *eventBuilder) Awards() *awardsBuilder {
	return &awardsBuilder{
		event:  b.id,
		client: b.client,
	}
}
func (b *eventBuilder) Matches() *matchesBuilder {
	return &matchesBuilder{
		team:   b.team,
		event:  b.id,
		client: b.client,
	}
}
func (b *eventBuilder) Status() *statusBuilder {
	return &statusBuilder{
		team:   b.team,
		event:  b.id,
		client: b.client,
	}
}

// Awards + modifiers
func (b *awardsBuilder) Year(year int) *awardsBuilder {
	b.year = year
	return b
}
func (b *awardsBuilder) Team(team int) *awardsBuilder {
	b.team = team
	return b
}

// Match + modifiers
func (b *matchesBuilder) Year(year int) *matchesBuilder {
	b.year = year
	return b
}
func (b *matchesBuilder) Simple() *matchesBuilder {
	b.simple = true
	return b
}
func (b *matchesBuilder) Team(team int) *matchesBuilder {
	b.team = team
	return b
}
func (tba Client) Match(key string) *matchBuilder {
	return &matchBuilder{
		key:    key,
		client: &tba,
	}
}
func (b *matchBuilder) Simple() *matchBuilder {
	b.simple = true
	return b
}

// Media + modifiers
func (b *mediaBuilder) Year(year int) *mediaBuilder {
	b.year = year
	return b
}
func (b *mediaBuilder) Tag(tag string) *mediaBuilder {
	b.tag = tag
	return b
}

// District + modifiers
func (tba Client) Districts(year int) *districtsBuilder {
	return &districtsBuilder{
		year:   year,
		client: &tba,
	}
}
func (tba Client) District(abbreviation string, year int) *districtBuilder {
	return &districtBuilder{
		abbreviation: abbreviation,
		year:         year,
		client:       &tba,
	}
}
func (b *districtBuilder) Rankings() *districtRankingsBuilder {
	return &districtRankingsBuilder{
		abbreviation: b.abbreviation,
		year:         b.year,
		client:       b.client,
	}
}
func (b *districtBuilder) Events() *eventsBuilder {
	return &eventsBuilder{
		district: b.abbreviation,
		year:     b.year,
		client:   b.client,
	}
}
func (b *districtBuilder) Teams() *teamsBuilder {
	return &teamsBuilder{
		district: b.abbreviation,
		year:     b.year,
		client:   b.client,
	}
}

// The following models have no modifiers:
// Year, Robot, Social Media, District Points, Insights, OPRs, Predictions, Rankings
