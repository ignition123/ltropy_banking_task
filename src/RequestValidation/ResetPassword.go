package RequestValidation

import (
	"lib"
	"messages"
	"objects"
)

// request validation of broker registration
func ResetPassword(pojoObj objects.ResetPassword) objects.ErrorResponse{

	defer lib.HandlePanic()

	// email cannot be empty
	if pojoObj.Email == nil || *pojoObj.Email == ""{

		return objects.ErrorResponse{
			Status:false,
			Msg:messages.EmailEmpty,
		}
	}

	// default response true
	return objects.ErrorResponse{
		Status: true,
	}
}
