package brokers

import (
	"RequestValidation"
	"encoding/json"
	"lib"
	"messages"
	"models/brokers"
	"net/http"
	"objects"
)

func GetAllBrokerIds(req *http.Request, res http.ResponseWriter){

	defer lib.HandleHttpPanic(res)

	sessionObject := objects.Session{}

	// creating object from redis session
	if err := json.Unmarshal([]byte(req.Header.Get("Redissession")), &sessionObject); err != nil{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.ErrorMessage,
		}

		json.NewEncoder(res).Encode(responseObject)

		return

	}

	// calling the objects of GetAllBrokerIds
	pojoObj := objects.GetAllBrokerIds{}

	// decoding the json request
	err := json.NewDecoder(req.Body).Decode(&pojoObj)

	// if error then message
	if err != nil{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.RequestValidateFail,
		}

		json.NewEncoder(res).Encode(responseObject)

		return

	}

	// validating the request
	requestVal := RequestValidation.GetAllBrokerIds(sessionObject, pojoObj)

	if !requestVal.Status{

		json.NewEncoder(res).Encode(requestVal)

		return
	}

	// getting all trading id of the brokers
	result, err := brokers.GetAllBrokerIds(pojoObj)

	if err != nil{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.ErrorMessage,
		}

		json.NewEncoder(res).Encode(responseObject)

		return

	}

	// no broker trading id found
	if len(result) == 0{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.NoTradingIds,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	// broker contacts verified and approved
	responseObject := &objects.TradingIdList{
		Status:true,
		TradingIds:result,
	}

	json.NewEncoder(res).Encode(responseObject)
}
