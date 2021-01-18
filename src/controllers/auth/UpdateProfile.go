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

func UpdateProfile(req *http.Request, res http.ResponseWriter){

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

	// calling the objects of UpdateProfile
	pojoObj := &objects.UpdateProfile{}

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

	// validating the request fields
	requestVal := RequestValidation.UpdateProfile(pojoObj)

	if !requestVal.Status{

		json.NewEncoder(res).Encode(requestVal)

		return
	}

	// update profile of the user
	updateRes, err := auth.UpdateProfile(sessionObject, pojoObj)

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

	// permission updated
	responseObject := &objects.SuccessResponse{
		Status:true,
		Msg:messages.UpdateProfileSuccess,
	}

	json.NewEncoder(res).Encode(responseObject)
}
