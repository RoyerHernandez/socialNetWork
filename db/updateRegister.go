package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/RoyerHernandez/socialNetWork.git/models"
)

func UpdateRegister(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoCN.Database("socialNetwork")
	col := bd.Collection("users")

	register := make(map[string]interface{})
	if len(user.Name) > 0 {
		register["name"] = user.Name
	}
	if len(user.LastName) > 0 {
		register["lastName"] = user.LastName
	}
	if len(user.Email) > 0 {
		register["email"] = user.Email
	}
	register["birthDate"] = user.BirthDate
	if len(user.Location) > 0 {
		register["location"] = user.Location
	}
	if len(user.WebSite) > 0 {
		register["webSite"] = user.WebSite
	}
	if len(user.Biography) > 0 {
		register["biography"] = user.Biography
	}
	if len(user.Password) > 0 {
		register["password"] = user.Password
	}
	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		register["banner"] = user.Banner
	}
	updateString := bson.M{
		"$set": register,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
