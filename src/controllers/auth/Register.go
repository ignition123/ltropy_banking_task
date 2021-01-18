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

func Register(req *http.Request, res http.ResponseWriter){

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

	// calling the objects of Registration
	pojoObj := objects.Register{}

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
	requestVal := RequestValidation.RegisterValidate(pojoObj)

	if !requestVal.Status{

		json.NewEncoder(res).Encode(requestVal)

    	return
	}

	// checking if account already exists
	_, accountErr := auth.CheckProfile(pojoObj)

	if accountErr == nil{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.UserAlreadyExists,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	// creating new admin
	insertRes, err := auth.CreateAdmin(pojoObj)

	if err != nil{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.ErrorMessage,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	if insertRes.InsertedID == nil{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.ErrorMessage,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	// success response after successfully adding new record to database
	responseObject := &objects.SuccessResponse{
    	Status:true,
    	Msg:messages.RegisterSuccess,
    }

	json.NewEncoder(res).Encode(responseObject)
}