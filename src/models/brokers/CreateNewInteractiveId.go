package brokers

import (
	"MongoConnection"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"lib"
	"objects"
	"time"
)

func CreateNewInteractiveId(sessionObject objects.Session, pojoObj objects.CreateNewInteractiveId) error{

	defer lib.HandlePanic()

	ctx := context.Background()

	brokerId, objectErr := primitive.ObjectIDFromHex(*pojoObj.Id)

	if objectErr != nil{
		return objectErr
	}

	err := MongoConnection.Database().Client().UseSession(ctx, func(sessionContext mongo.SessionContext) error{

		defer lib.HandlePanic()

		err := sessionContext.StartTransaction()

		if err != nil {
			return err
		}

		_, err = MongoConnection.Database().Collection("broker_ids").InsertOne(sessionContext, bson.M{
			"brokerId": brokerId,
			"LUT": time.Now(),
			"idType": 1,
			"idName": "Interactive",
			"CreatedDate": time.Now(),
			"password":"coynce@12345",
			"blockStatus":false,
		})

		if err != nil{
			return err
		}

		_, err = MongoConnection.Database().Collection("admin_logs").InsertOne(sessionContext, bson.M{
			"adminEmail": *sessionObject.Email,
			"LUT": time.Now(),
			"eventType": "broker",
			"event": "Created new interactive id",
			"brokerId": brokerId,
		})

		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			return err
		}

		sessionContext.CommitTransaction(sessionContext)

		return nil
	})

	return err
}
