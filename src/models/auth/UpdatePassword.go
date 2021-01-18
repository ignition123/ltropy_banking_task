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

func UpdatePassword(sessionObject objects.Session, pojoObj objects.UpdatePassword) (*mongo.UpdateResult, error){

	defer lib.HandlePanic()

	// setting default password for admin
	hashPassword := sha256.New()
	hashPassword.Write([]byte(*pojoObj.Password))

	ctx := context.Background()

	opts := options.Update().SetUpsert(false)

	filter := bson.D{{
		"email", *sessionObject.Email,
	}}

	update := bson.D{{"$set", bson.D{
		{
			"password", hex.EncodeToString(hashPassword.Sum(nil)),
		},
		{
			"LUT", time.Now(),
		},
	}}}

	// updating the offline people in mongodb

	updateRes, err := MongoConnection.Database().Collection("admin").UpdateOne(ctx, filter, update, opts)

	return updateRes, err
}
