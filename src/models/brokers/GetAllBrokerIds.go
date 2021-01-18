package brokers

import (
	"MongoConnection"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lib"
	"objects"
)

func GetAllBrokerIds(pojoObj objects.GetAllBrokerIds) ([]map[string]interface{}, error){

	defer lib.HandlePanic()

	brokerId, objectErr := primitive.ObjectIDFromHex(*pojoObj.Id)

	if objectErr != nil{
		return nil, objectErr
	}

	brokerIds := make([]map[string]interface{}, 0)

	ctx := context.Background()

	filter := bson.D{{
		"brokerId", brokerId,
	}}

	project := options.Find().SetProjection(bson.M{"LUT":1, "idType":1, "idName":1, "CreatedDate":1, "blockStatus":1, "brokerId":1, "_id":1})

	cursor, brokerErr := MongoConnection.Database().Collection("broker_ids").Find(ctx, filter, project)

	if brokerErr = cursor.All(ctx, &brokerIds); brokerErr != nil {
		return nil, brokerErr
	}

	return brokerIds, nil
}
