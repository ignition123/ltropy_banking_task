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

func ResetPassword(req *http.Request, res http.ResponseWriter){

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

	// checking if the user has rights to create new admin
	if !*sessionObject.SuperAdmin && (sessionObject.Permissions == nil || !*sessionObject.Permissions.CreateAdmin){

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.DontHavePermission,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	// calling the objects of ResetPassword
	pojoObj := objects.ResetPassword{}

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
	requestVal := RequestValidation.ResetPassword(pojoObj)

	if !requestVal.Status{

		json.NewEncoder(res).Encode(requestVal)

		return
	}

	// set password to default password
	updateRes, err := auth.ResetPassword(pojoObj)

	if err != nil{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.ErrorMessage,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	// checking update result
	if updateRes.ModifiedCount == 0 && updateRes.UpsertedCount == 0{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.PasswordAlreadyDefault,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	// success response after successfully
	responseObject := &objects.SuccessResponse{
		Status:true,
		Msg:messages.PasswordResetSuccess,
	}

	json.NewEncoder(res).Encode(responseObject)
}
