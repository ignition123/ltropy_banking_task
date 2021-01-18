package RequestValidation

import (
	"lib"
	"messages"
	"objects"
	"regexp"
)

func UpdateProfile(pojoObj *objects.UpdateProfile) objects.ErrorResponse{

	defer lib.HandlePanic()

	// checking if name empty
	if pojoObj.Name == nil || *pojoObj.Name == ""{
		return objects.ErrorResponse{
			Status:false,
			Msg:messages.NameEmpty,
		}
	}

	// checking if designation empty
	if *pojoObj.Designation == ""{
		pojoObj.Designation = nil
	}

	// checking DOB is empty
	if *pojoObj.DOB == ""{
		pojoObj.DOB = nil
	}

	// checking if nickName is empty
	if *pojoObj.NickName == ""{
		pojoObj.NickName = nil
	}

	if pojoObj.DOB != nil{

		// creating regex to match the pattern for DOB
		re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")

		if !re.MatchString(*pojoObj.DOB){
			return objects.ErrorResponse{
				Status:false,
				Msg:messages.InvalidDOB,
			}
		}
	}

	// default response true
	return objects.ErrorResponse{
		Status: true,
	}
}
