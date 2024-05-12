package mongo

import (
	"context"
	"errors"
	"feedback-bot/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var errInsertFail = errors.New("failed to insert entry")

type Mongo struct {
	questionsColl *mongo.Collection
	qaColl        *mongo.Collection
}

func (c *Mongo) GetQuestion() *entity.Question {
	cur, err := c.questionsColl.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		ans := entity.Question{}
		cur.Decode(&ans)
		return &ans
	}

	return nil
}

func (c *Mongo) InsertQuestion(question *entity.Question) error {
	_, err := c.questionsColl.InsertOne(context.Background(), question)
	if err != nil {
		return errInsertFail
	}

	return nil
}

func (c *Mongo) InsertQA(qa *entity.QA) error {
	_, err := c.qaColl.InsertOne(context.Background(), qa)
	if err != nil {
		return errInsertFail
	}

	return nil
}
