package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/RoyerHernandez/socialNetWork.git/models"
)

func AlreadyExistUser(email string) (models.User, bool, string) {
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("socialNetwork")
	col := db.Collection("users")

	filter := bson.M{"email": email}

	err := col.FindOne(ctx, filter).Decode(&user)
	Id := user.ID
	if err != nil {
		return user, false, Id
	}
	return user, true, Id
}
