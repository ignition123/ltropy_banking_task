package auth

import (
	"MongoConnection"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lib"
	"objects"
	"time"
)

func ResetPassword(pojoObj objects.ResetPassword) (*mongo.UpdateResult, error){

	defer lib.HandlePanic()

	// setting default password for admin
	hash := sha256.New()
	hash.Write([]byte("admin@123456"))

	ctx := context.Background()

	opts := options.Update().SetUpsert(false)

	filter := bson.D{{
		"email", *pojoObj.Email,
	}}

	update := bson.D{{"$set", bson.D{
		{
			"password", hex.EncodeToString(hash.Sum(nil)),
		},
		{
			"LUT", time.Now(),
		},
	}}}

	// updating the offline people in mongodb

	updateRes, err := MongoConnection.Database().Collection("admin").UpdateOne(ctx, filter, update, opts)

	return updateRes, err
}
