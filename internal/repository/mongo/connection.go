package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectToMongo(mongoURL string) *Mongo {
	opts := options.Client().ApplyURI(mongoURL)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	var result bson.M
	if err = client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		log.Fatal(err)
	}

	questionsColl := client.Database("admin").Collection("questions")
	qaColl := client.Database("admin").Collection("qa")

	return &Mongo{
		questionsColl: questionsColl,
		qaColl:        qaColl,
	}
}

func initMigrations() {

}
