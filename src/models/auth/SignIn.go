package auth

import (
	"MongoConnection"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"lib"
	"objects"
)

func SignIn(pojoObj objects.SignIn) (map[string]interface{}, error){

	defer lib.HandlePanic()

	adminDetails := make(map[string]interface{})

	ctx := context.Background()

	filter := bson.D{{
		"email", *pojoObj.Email,
	},
	{
		"password", *pojoObj.Password,
	}}

	userErr := MongoConnection.Database().Collection("admin").FindOne(ctx, filter).Decode(&adminDetails)

	if userErr != nil{

		return nil, userErr
	}

	return adminDetails, nil
}
