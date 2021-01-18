package RequestValidation

import (
	"lib"
	"messages"
	"objects"
)

func GetAllBrokerIds(sessionObject objects.Session, pojoObj objects.GetAllBrokerIds) objects.ErrorResponse{

	defer lib.HandlePanic()

	if pojoObj.Id == nil || *pojoObj.Id == ""{
		return objects.ErrorResponse{
			Status:false,
			Msg:messages.CompanyIdEmpty,
		}
	}

	if !*sessionObject.SuperAdmin{
		if sessionObject.Permissions == nil || sessionObject.Permissions.Broker == nil || !*sessionObject.Permissions.Broker.CreateNewTradingId{
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
