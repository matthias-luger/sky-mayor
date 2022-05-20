package model

type ApiCandidates struct {
	Key   string  `json:"key"`
	Name  string  `json:"name"`
	Perks []*Perk `json:"perks"`
	Votes int     `json:"votes"`
}

type ApiElectionData struct {
	Year       int              `json:"year"`
	Candidates []*ApiCandidates `json:"candidates"`
}

type ApiMayorData struct {
	Key      string          `json:"key"`
	Name     string          `json:"name"`
	Perks    []*Perk         `json:"perks"`
	Election ApiElectionData `json:"election"`
}

type ApiElectionResponse struct {
	Success     bool            `json:"success"`
	LastUpdated int64           `json:"lastUpdated"`
	Mayor       ApiMayorData    `json:"mayor"`
	Current     ApiElectionData `json:"current"`
}
