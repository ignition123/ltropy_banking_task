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

func UpdatePermission(pojoObj objects.UpdatePermission) (*mongo.UpdateResult, error){

	defer lib.HandlePanic()

	// creating update document
	updateDoc := make(map[string]interface{})

	updateDoc["LUT"] = time.Now()
	updateDoc["superAdmin"] = *pojoObj.SuperAdmin

	// creating permission hashmap
	updatePermissionDoc := make(map[string]map[string]interface{})

	if pojoObj.Permissions != nil{
		updatePermissionDoc["permissions"] = make(map[string]interface{})
	}

	// setting for createAdmin boolean
	if pojoObj.Permissions.CreateAdmin != nil{
		updatePermissionDoc["permissions"]["createAdmin"] = *pojoObj.Permissions.CreateAdmin
	}

	// checking if KYC exists then setting permission for KYC permission
	if pojoObj.Permissions.KYC != nil{

		updatePermissionDoc["permissions"]["kyc"] = make(map[string]interface{})
		kycMap := make(map[string]interface{})

		if pojoObj.Permissions.KYC.KycComplete != nil{
			kycMap["kycComplete"] = *pojoObj.Permissions.KYC.KycComplete
		}

		if pojoObj.Permissions.KYC.ApproveContact != nil{
			kycMap["approveContact"] = *pojoObj.Permissions.KYC.ApproveContact
		}

		updatePermissionDoc["permissions"]["kyc"] = kycMap
	}

	// checking for broker exists then setting permission for Broker permission
	if pojoObj.Permissions.Broker != nil{

		updatePermissionDoc["permissions"]["brokers"] = make(map[string]interface{})
		brokersMap := make(map[string]interface{})

		if pojoObj.Permissions.Broker.AccountBlocked != nil{
			brokersMap["account_blocked"] = *pojoObj.Permissions.Broker.AccountBlocked
		}

		if pojoObj.Permissions.Broker.UpdateProfile != nil{
			brokersMap["update_profile"] = *pojoObj.Permissions.Broker.UpdateProfile
		}

		if pojoObj.Permissions.Broker.FreezeAccount != nil{
			brokersMap["freeze_account"] = *pojoObj.Permissions.Broker.FreezeAccount
		}

		if pojoObj.Permissions.Broker.CreateNewTradingId != nil{
			brokersMap["create_new_trading_id"] = *pojoObj.Permissions.Broker.CreateNewTradingId
		}

		if pojoObj.Permissions.Broker.ResetBrokerPassword != nil{
			brokersMap["reset_broker_password"] = *pojoObj.Permissions.Broker.ResetBrokerPassword
		}

		updatePermissionDoc["permissions"]["brokers"] = brokersMap
	}

	// checking for exchange push and setting permission
	if pojoObj.Permissions.ExchangePush != nil {
		updatePermissionDoc["permissions"]["exchange_push"] = make(map[string]interface{})
		exchangePushMap := make(map[string]interface{})

		if pojoObj.Permissions.ExchangePush.Messages != nil{
			exchangePushMap["messages"] = *pojoObj.Permissions.ExchangePush.Messages
		}

		updatePermissionDoc["permissions"]["exchange_push"] = exchangePushMap
	}

	// setting permission object to null if its superAdmin or permission object is nil
	if *pojoObj.SuperAdmin{
		updateDoc["permissions"] = nil
	}else{
		if pojoObj.Permissions != nil{
			updateDoc["permissions"] = updatePermissionDoc["permissions"]
		}else{
			updateDoc["permissions"] = nil
		}
	}

	ctx := context.Background()

	opts := options.Update().SetUpsert(false)

	filter := bson.D{{
		"email", *pojoObj.Email,
	}}

	update := bson.D{{"$set", updateDoc}}

	// updating the offline people in mongodb

	updateRes, err := MongoConnection.Database().Collection("admin").UpdateOne(ctx, filter, update, opts)

	return updateRes, err
}
