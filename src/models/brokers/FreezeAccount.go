package brokers

import (
	"MongoConnection"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lib"
	"objects"
	"time"
)

func FreezeAccount(sessionObject objects.Session, pojoObj objects.FreezeAccount) error{

	defer lib.HandlePanic()

	ctx := context.Background()

	brokerId, objectErr := primitive.ObjectIDFromHex(*pojoObj.Id)

	if objectErr != nil{
		return objectErr
	}

	err := MongoConnection.Database().Client().UseSession(ctx, func(sessionContext mongo.SessionContext) error {

		defer lib.HandlePanic()

		err := sessionContext.StartTransaction()

		if err != nil {
			return err
		}

		opts := options.Update().SetUpsert(false)

		filter := bson.D{{
			"_id", brokerId,
		}}

		update := bson.D{{"$set", bson.D{
			{
				"freezeAccount", true,
			},
			{
				"LUT", time.Now(),
			},
		}}}

		// updating customer approval
		updateRes, updateErr := MongoConnection.Database().Collection("companies").UpdateOne(ctx, filter, update, opts)

		if updateErr != nil{
			return err
		}

		if updateRes.ModifiedCount == 0 && updateRes.UpsertedCount == 0{
			return errors.New("failed to freeze account of the broker")
		}

		_, err = MongoConnection.Database().Collection("admin_logs").InsertOne(sessionContext, bson.M{
			"adminEmail": *sessionObject.Email,
			"LUT": time.Now(),
			"eventType": "broker",
			"event": "Freezing account",
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
