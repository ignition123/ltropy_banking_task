package auth

import (
	"MongoConnection"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"lib"
	"objects"
	"time"
)

/*
	This function checks if the email address already exists in the database
*/

func CheckProfile(pojoObj objects.Register) (map[string]interface{}, error){

	defer lib.HandlePanic()

	userDetails := make(map[string]interface{})

	ctx := context.Background()

	filter := bson.D{{
		"email", *pojoObj.Email,
	}}

	userErr := MongoConnection.Database().Collection("admin").FindOne(ctx, filter).Decode(&userDetails)

	if userErr != nil{

		return nil, userErr
	}

	return userDetails, nil
}

/*
	Creating new profile for admin
*/

func CreateAdmin(pojoObj objects.Register) (*mongo.InsertOneResult, error){

	defer lib.HandlePanic()

	// creating insert document
	insertDoc := make(map[string]interface{})

	// setting default password for admin
	hash := sha256.New()
	hash.Write([]byte("admin@123456"))

	// setting values
	insertDoc["email"] = *pojoObj.Email
	insertDoc["name"] = *pojoObj.Name
	insertDoc["password"] = hex.EncodeToString(hash.Sum(nil))
	insertDoc["dp"] = nil
	insertDoc["LUT"] = time.Now()
	insertDoc["CreatedDate"] = time.Now()
	insertDoc["superAdmin"] = *pojoObj.SuperAdmin

	// creating permission hashmap
	insertPermissionDoc := make(map[string]map[string]interface{})

	if pojoObj.Permissions != nil{
		insertPermissionDoc["permissions"] = make(map[string]interface{})
	}

	// setting for createAdmin boolean
	if pojoObj.Permissions.CreateAdmin != nil{
		insertPermissionDoc["permissions"]["createAdmin"] = *pojoObj.Permissions.CreateAdmin
	}

	// checking if KYC exists then setting permission for KYC permission
	if pojoObj.Permissions.KYC != nil{

		insertPermissionDoc["permissions"]["kyc"] = make(map[string]interface{})
		kycMap := make(map[string]interface{})

		if pojoObj.Permissions.KYC.KycComplete != nil{
			kycMap["kycComplete"] = *pojoObj.Permissions.KYC.KycComplete
		}

		if pojoObj.Permissions.KYC.ApproveContact != nil{
			kycMap["approveContact"] = *pojoObj.Permissions.KYC.ApproveContact
		}

		insertPermissionDoc["permissions"]["kyc"] = kycMap
	}

	// checking for broker exists then setting permission for Broker permission
	if pojoObj.Permissions.Broker != nil{

		insertPermissionDoc["permissions"]["brokers"] = make(map[string]interface{})
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

		insertPermissionDoc["permissions"]["brokers"] = brokersMap
	}

	// checking for exchange push and setting permission
	if pojoObj.Permissions.ExchangePush != nil {
		insertPermissionDoc["permissions"]["exchange_push"] = make(map[string]interface{})
		exchangePushMap := make(map[string]interface{})

		if pojoObj.Permissions.ExchangePush.Messages != nil{
			exchangePushMap["messages"] = *pojoObj.Permissions.ExchangePush.Messages
		}

		insertPermissionDoc["permissions"]["exchange_push"] = exchangePushMap
	}

	// setting permission object to null if its superAdmin or permission object is nil
	if *pojoObj.SuperAdmin{
		insertDoc["permissions"] = nil
	}else{
		if pojoObj.Permissions != nil{
			insertDoc["permissions"] = insertPermissionDoc["permissions"]
		}else{
			insertDoc["permissions"] = nil
		}
	}

	insertResult, insertErr := MongoConnection.Database().Collection("admin").InsertOne(context.TODO(), insertDoc)

	return insertResult, insertErr
}