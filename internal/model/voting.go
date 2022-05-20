package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vote struct {
	MayorKey string `bson:"key"`
	Votes    int    `bson:"votes"`
}

type Voting struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Year      int                `bson:"year"`
	Votes     []*Vote            `bson:"votes"`
	Timestamp time.Time          `bson:"timestamp"`
}
