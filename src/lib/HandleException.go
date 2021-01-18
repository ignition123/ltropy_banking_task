package lib

import (
	"encoding/json"
	"fmt"
	"messages"
	"net/http"
	"objects"
)

func HandleHttpPanic(res http.ResponseWriter) { 
  
    if a := recover(); a != nil { 

        fmt.Println("RECOVER "+a.(string))

         responseObject := &objects.ErrorResponse{
	    	Status:false,
	    	Msg:messages.ErrorMessage,
	    }

    	json.NewEncoder(res).Encode(responseObject)
        
    } 

} 

func HandlePanic(){

	if a := recover(); a != nil{

		fmt.Println("RECOVER "+a.(string))
	}

}