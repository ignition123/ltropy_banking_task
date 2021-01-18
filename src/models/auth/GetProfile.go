package auth

import (
	"MongoConnection"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lib"
	"objects"
)

func GetProfile(sessionObject objects.Session) (map[string]interface{}, error){

	defer lib.HandlePanic()

	adminDetails := make(map[string]interface{})

	ctx := context.Background()

	filter := bson.D{{
		"email", *sessionObject.Email,
	}}

	project := options.FindOne().SetProjection(bson.M{"email":1, "name":1, "superAdmin":1, "CreatedDate":1, "permissions":1, "dp":1, "DOB":1, "designation":1, "nickName":1, "_id":0})

	userErr := MongoConnection.Database().Collection("admin").FindOne(ctx, filter, project).Decode(&adminDetails)

	if userErr != nil{

		return nil, userErr
	}

	return adminDetails, nil
}
