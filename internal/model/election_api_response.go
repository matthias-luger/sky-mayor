package model

type ApiCandidates struct {
	Key   string  `bson:"key"`
	Name  string  `bson:"key"`
	Perks []*Perk `bson:"perks"`
	Votes int     `bson:"votes"`
}

type ApiElectionData struct {
	Year       int              `bson:"year"`
	Candidates []*ApiCandidates `bson:"candidates"`
}

type ApiMayorData struct {
	Key      string          `bson:"key"`
	Name     string          `bson:"name"`
	Perks    []*Perk         `bson:"perks"`
	Election ApiElectionData `bson:"election"`
}

type ApiElectionResponse struct {
	Success     bool            `bson:"success"`
	LastUpdated int64           `bson:"lastUpdated"`
	Mayor       ApiMayorData    `bson:"mayor"`
	Current     ApiElectionData `bson:"current"`
}
