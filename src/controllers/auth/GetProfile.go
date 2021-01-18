package auth

import (
	"encoding/json"
	"lib"
	"messages"
	"models/auth"
	"net/http"
	"objects"
)

func GetProfile(req *http.Request, res http.ResponseWriter){

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

	// get profile of the admin
	findRes, err := auth.GetProfile(sessionObject)

	if err != nil{
		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.ErrorMessage,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	// entire profile response
	responseObject := &objects.ProfileResp{
		Status:true,
		Profile:findRes,
	}

	json.NewEncoder(res).Encode(responseObject)
}
