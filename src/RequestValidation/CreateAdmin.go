package RequestValidation

import (
	"lib"
	"messages"
	"objects"
)

func CreateAdminValidate(pojoObj objects.CreateSuperAdmin) objects.ErrorResponse{

	defer lib.HandlePanic()

	// validate fields exists in the request

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
