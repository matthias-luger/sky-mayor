package mongo

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client                   *mongo.Client
	electionPeriodCollection *mongo.Collection
	votingCollection         *mongo.Collection
)

func Init() error {

	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URL")))

	if err != nil {
		return err
	}

	electionPeriodCollection = client.Database("sky_mayor").Collection("election_period")
	votingCollection = client.Database("sky_mayor").Collection("voting")

	return nil
}

func Disconnect() {
	ctx := context.Background()
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
