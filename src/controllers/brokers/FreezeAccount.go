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

func FreezeAccount(req *http.Request, res http.ResponseWriter){

	defer lib.HandleHttpPanic(res)

	sessionObject := objects.Session{}

	// creating object from redis session
	if err := json.Unmarshal([]byte(req.Header.Get("Redissession")), &sessionObject); err != nil {

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.ErrorMessage,
		}

		json.NewEncoder(res).Encode(responseObject)

		return

	}

	// calling the objects of FreezeAccount
	pojoObj := objects.FreezeAccount{}

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
	requestVal := RequestValidation.FreezeAccount(sessionObject, pojoObj)

	if !requestVal.Status{

		json.NewEncoder(res).Encode(requestVal)

		return
	}

	// validating password of the admin
	_, matchErr := brokers.MatchPassword(sessionObject, *pojoObj.Password)

	if matchErr != nil{
		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.InvalidPassword,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	// freezing account of the brokers
	tranErr := brokers.FreezeAccount(sessionObject, pojoObj)

	if tranErr != nil{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.FailedToFrozeAccount,
		}

		json.NewEncoder(res).Encode(responseObject)

		return

	}

	// broker contacts verified and froze account
	responseObject := &objects.SuccessResponse{
		Status:true,
		Msg:messages.BrokerAccountFroze,
	}

	json.NewEncoder(res).Encode(responseObject)
}