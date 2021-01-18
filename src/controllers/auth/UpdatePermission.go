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

func UpdatePermission(req *http.Request, res http.ResponseWriter){

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

	// calling the objects of UpdatePermission
	pojoObj := objects.UpdatePermission{}

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
	requestVal := RequestValidation.UpdatePermission(sessionObject, pojoObj)

	if !requestVal.Status{

		json.NewEncoder(res).Encode(requestVal)

		return
	}

	// update permission of the user
	updateRes, err := auth.UpdatePermission(pojoObj)

	if err != nil{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.ErrorMessage,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	// checking update results
	if updateRes.ModifiedCount == 0 && updateRes.UpsertedCount == 0{

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
		Msg:messages.PermissionUpdated,
	}

	json.NewEncoder(res).Encode(responseObject)
}
