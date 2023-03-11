package db

import (
	"context"
	"github.com/RoyerHernandez/socialNetWork.git/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func InsertRegister(user models.User) (string, bool, error) {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("socialNetwork")
	col := db.Collection("users")

	user.Password, err = PasswordEncrypt(user.Password)
	if err != nil {
		return "", false, err
	}

	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
