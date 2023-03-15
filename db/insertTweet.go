package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/RoyerHernandez/socialNetWork.git/models"
)

func InsertTweet(twt models.RecordTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("socialNetwork")
	col := db.Collection("tweet")

	tweet := bson.M{
		"userID":  twt.UserID,
		"message": twt.Message,
		"date":    twt.Date,
	}
	result, err := col.InsertOne(ctx, tweet)
	if err != nil {
		return "", false, err
	}
	objId, _ := result.InsertedID.(primitive.ObjectID)
	return objId.String(), true, nil
}
