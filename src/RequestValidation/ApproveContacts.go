package RequestValidation

import (
	"lib"
	"messages"
	"objects"
)

func ApproveContacts(sessionObject objects.Session, pojoObj objects.ApproveContacts) objects.ErrorResponse{

	defer lib.HandlePanic()

	if pojoObj.Id == nil || *pojoObj.Id == ""{
		return objects.ErrorResponse{
			Status:false,
			Msg:messages.CompanyIdEmpty,
		}
	}

	if pojoObj.Password == nil || *pojoObj.Password == ""{
		return objects.ErrorResponse{
			Status:false,
			Msg:messages.PasswordEmpty,
		}
	}

	if !*sessionObject.SuperAdmin{
		if sessionObject.Permissions == nil || sessionObject.Permissions.KYC == nil || !*sessionObject.Permissions.KYC.ApproveContact{
			return objects.ErrorResponse{
				Status:false,
				Msg:messages.DontHavePermission,
			}
		}
	}

	// default response true
	return objects.ErrorResponse{
		Status: true,
	}
}
