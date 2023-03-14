package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/RoyerHernandez/socialNetWork.git/models"
)

func SearchProfile(Id string) (models.User, error) {
	var profile models.User
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("socialNetwork")
	col := db.Collection("users")

	objId, _ := primitive.ObjectIDFromHex(Id)
	filter := bson.M{"_id": objId}

	err := col.FindOne(ctx, filter).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("register not found ! " + err.Error())
		return profile, err
	}
	return profile, nil
}
