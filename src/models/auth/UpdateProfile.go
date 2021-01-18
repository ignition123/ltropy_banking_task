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

func UpdateProfile(sessionObject objects.Session, pojoObj *objects.UpdateProfile) (*mongo.UpdateResult, error){

	defer lib.HandlePanic()

	// creating update document
	updateDoc := make(map[string]interface{})

	updateDoc["name"] = *pojoObj.Name
	updateDoc["designation"] = *pojoObj.Designation
	updateDoc["DOB"] = *pojoObj.DOB
	updateDoc["nickName"] = *pojoObj.NickName
	updateDoc["LUT"] = time.Now()

	ctx := context.Background()

	opts := options.Update().SetUpsert(false)

	filter := bson.D{{
		"email", *sessionObject.Email,
	}}

	update := bson.D{{"$set", updateDoc}}

	// updating the offline people in mongodb

	updateRes, err := MongoConnection.Database().Collection("admin").UpdateOne(ctx, filter, update, opts)

	return updateRes, err
}
