package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Perk struct {
	Name        string `bson:"name"`
	Description string `bson:"description"`
}

type Candidate struct {
	Name  string  `bson:"name"`
	Key   string  `bson:"key"`
	Perks []*Perk `bson:"perks"`
}

type ElectionPeriod struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Year       int                `bson:"year"`
	Start      time.Time          `bson:"start"`
	End        time.Time          `bson:"end"`
	Candidates []*Candidate       `bson:"candidates"`
	Winner     *Candidate         `bson:"winner"`
}
