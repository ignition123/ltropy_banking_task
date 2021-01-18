package RequestValidation

import (
	"lib"
	"messages"
	"objects"
)

func UpdatePassword(pojoObj objects.UpdatePassword) objects.ErrorResponse{

	defer lib.HandlePanic()

	// password cannot be empty
	if pojoObj.Password == nil || *pojoObj.Password == ""{

		return objects.ErrorResponse{
			Status:false,
			Msg:messages.PasswordEmpty,
		}
	}

	// confirm password cannot be empty
	if pojoObj.ConfirmPassword == nil || *pojoObj.ConfirmPassword == ""{

		return objects.ErrorResponse{
			Status:false,
			Msg:messages.ConfirmPasswordEmpty,
		}
	}

	// matching the password and confirm password
	if *pojoObj.Password != *pojoObj.ConfirmPassword{

		return objects.ErrorResponse{
			Status:false,
			Msg:messages.PasswordDoesNotMatch,
		}
	}

	// default response true
	return objects.ErrorResponse{
		Status: true,
	}

}
