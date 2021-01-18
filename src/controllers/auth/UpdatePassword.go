package auth

import (
	"RequestValidation"
	"encoding/json"
	"lib"
	"messages"
	"models/auth"
	"net/http"
	"objects"
)

func UpdatePassword(req *http.Request, res http.ResponseWriter){

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

	// calling the objects of UpdatePassword
	pojoObj := objects.UpdatePassword{}

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

	// validate the request
	// validating the request fields
	requestVal := RequestValidation.UpdatePassword(pojoObj)

	if !requestVal.Status{

		json.NewEncoder(res).Encode(requestVal)

		return
	}

	// update password to database
	updateRes, err := auth.UpdatePassword(sessionObject, pojoObj)

	if err != nil{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.ErrorMessage,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	// checking update result
	if updateRes.UpsertedCount == 0 && updateRes.ModifiedCount == 0{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.ErrorMessage,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	// password updated successfully
	responseObject := &objects.SuccessResponse{
		Status:true,
		Msg:messages.PasswordUpdatedSuccess,
	}

	json.NewEncoder(res).Encode(responseObject)
}
