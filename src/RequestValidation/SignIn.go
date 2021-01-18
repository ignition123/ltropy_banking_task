package RequestValidation

import (
	"lib"
	"messages"
	"objects"
)

func SignInValidate(pojoObj objects.SignIn) objects.ErrorResponse{

	defer lib.HandlePanic()

	// validate fields exists in the request

	// email cannot be empty
	if pojoObj.Email == nil || *pojoObj.Email == ""{

		return objects.ErrorResponse{
			Status:false,
			Msg:messages.EmailEmpty,
		}
	}

	// password cannot be empty
	if pojoObj.Password == nil || *pojoObj.Password == ""{

		return objects.ErrorResponse{
			Status:false,
			Msg:messages.PasswordEmpty,
		}
	}

	// default response true
	return objects.ErrorResponse{
		Status: true,
	}
}
