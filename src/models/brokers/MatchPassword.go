package brokers

import (
	"MongoConnection"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lib"
	"objects"
)

func MatchPassword(sessionObject objects.Session, password string) (map[string]interface{}, error){

	defer lib.HandlePanic()

	// setting default password for admin
	hash := sha256.New()
	hash.Write([]byte(password))

	password = hex.EncodeToString(hash.Sum(nil))

	adminDetails := make(map[string]interface{})

	ctx := context.Background()

	filter := bson.D{{
		"email", *sessionObject.Email,
	},
	{
		"password",password,
	}}

	project := options.FindOne().SetProjection(bson.M{"_id":1})

	userErr := MongoConnection.Database().Collection("admin").FindOne(ctx, filter, project).Decode(&adminDetails)

	if userErr != nil{

		return nil, userErr
	}

	return adminDetails, nil
}
