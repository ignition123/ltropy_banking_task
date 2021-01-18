package auth

import (
	"MongoConnection"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lib"
	"objects"
	"time"
)

func CreateSuperAdmin(pojoObj objects.CreateSuperAdmin) (*mongo.UpdateResult, error){

	defer lib.HandlePanic()

	ctx := context.Background()

	opts := options.Update().SetUpsert(false)

	filter := bson.D{{
		"email", *pojoObj.Email,
	}}

	update := bson.D{{"$set", bson.D{
		{
			"superAdmin", true,
		},
		{
			"permissions", nil,
		},
		{
			"LUT", time.Now(),
		},
	}}}

	// updating the offline people in mongodb

	updateRes, err := MongoConnection.Database().Collection("admin").UpdateOne(ctx, filter, update, opts)

	return updateRes, err
}
