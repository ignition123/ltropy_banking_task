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

func CreateSuperAdmin(req *http.Request, res http.ResponseWriter){

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
	if !*sessionObject.SuperAdmin{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.DontHavePermission,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	// calling the objects of CreateSuperAdmin
	pojoObj := objects.CreateSuperAdmin{}

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
	requestVal := RequestValidation.CreateAdminValidate(pojoObj)

	if !requestVal.Status{

		json.NewEncoder(res).Encode(requestVal)

		return
	}

	// update database to create new super admin
	updateRes, err := auth.CreateSuperAdmin(pojoObj)

	if err != nil{
		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.ErrorMessage,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	// checking database update status
	if updateRes.ModifiedCount == 0 && updateRes.UpsertedCount == 0{
		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.ErrorMessage,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	// createAdmin rights updated
	responseObject := &objects.SuccessResponse{
		Status:true,
		Msg:messages.CreatedSuperAdminSuccess,
	}

	json.NewEncoder(res).Encode(responseObject)
}
