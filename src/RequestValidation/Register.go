package RequestValidation

import (
	"lib"
	"messages"
	"objects"
)

// request validation of broker registration

func RegisterValidate(pojoObj objects.Register) objects.ErrorResponse{

	defer lib.HandlePanic()

	// validate fields exists in the request

	// email cannot be empty
	if pojoObj.Email == nil || *pojoObj.Email == ""{

		return objects.ErrorResponse{
	    	Status:false,
	    	Msg:messages.EmailEmpty,
	    }
	}

	// name cannot be empty
	if pojoObj.Name == nil || *pojoObj.Name == ""{

		return objects.ErrorResponse{
	    	Status:false,
	    	Msg:messages.NameEmpty,
	    }
	}

	// super admin cannot be empty
	if pojoObj.SuperAdmin == nil{

		return objects.ErrorResponse{
	    	Status:false,
	    	Msg:messages.SuperAdminFlagEmpty,
	    }
	}

	// checking permissions if exists
	if pojoObj.Permissions != nil{

		// checking if creating admin rights is present
		if pojoObj.Permissions.CreateAdmin == nil{

			return objects.ErrorResponse{
				Status:false,
				Msg:messages.CreateAdminFlagEmpty,
			}
		}

		// checking values if KYC permissions are available
		if pojoObj.Permissions.KYC != nil{

			// checking contact approval permission
			if pojoObj.Permissions.KYC.ApproveContact == nil{

				return objects.ErrorResponse{
					Status:false,
					Msg:messages.KYCApproveContactEmpty,
				}
			}

			// checking if kyc complete persmission is available
			if pojoObj.Permissions.KYC.KycComplete == nil{

				return objects.ErrorResponse{
					Status:false,
					Msg:messages.KYCApproveEmpty,
				}
			}
		}

		// checking values if broker permissions available
		if pojoObj.Permissions.Broker != nil{

			// checking permissions of blocking broker account
			if pojoObj.Permissions.Broker.AccountBlocked == nil{

				return objects.ErrorResponse{
					Status:false,
					Msg:messages.AccountBlockedEmpty,
				}
			}

			// checking permissions of update broker profile
			if pojoObj.Permissions.Broker.UpdateProfile == nil{

				return objects.ErrorResponse{
					Status:false,
					Msg:messages.UpdateProfileEmpty,
				}
			}

			// checking permissions of freezing broker account
			if pojoObj.Permissions.Broker.FreezeAccount == nil{

				return objects.ErrorResponse{
					Status:false,
					Msg:messages.FreezeAccountEmpty,
				}
			}

			// checking permissions of create new broker id
			if pojoObj.Permissions.Broker.CreateNewTradingId == nil{

				return objects.ErrorResponse{
					Status:false,
					Msg:messages.CreateNewBrokerIdEmpty,
				}
			}

			// checking permissions of resetting broker password
			if pojoObj.Permissions.Broker.ResetBrokerPassword == nil{

				return objects.ErrorResponse{
					Status:false,
					Msg:messages.ResettingPasswordEmpty,
				}
			}
		}

		// checking values if exchange push permission available
		if pojoObj.Permissions.ExchangePush != nil{

			// checking permission of pushing exchange messages
			if pojoObj.Permissions.ExchangePush.Messages == nil{

				return objects.ErrorResponse{
					Status:false,
					Msg:messages.ExchangePushEmpty,
				}
			}
		}
	}

	// default response true
	return objects.ErrorResponse{
		Status: true,
	}
}