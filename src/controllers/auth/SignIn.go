package auth

import (
	"RedisConnection"
	"RequestValidation"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"lib"
	"messages"
	"models/auth"
	"net/http"
	"objects"
	"time"
)

func SignIn(req *http.Request, res http.ResponseWriter){

	defer lib.HandleHttpPanic(res)

	// calling the objects of SignIn
	pojoObj := objects.SignIn{}

	// decoding the json request
	err := json.NewDecoder(req.Body).Decode(&pojoObj)

	// if error then message
	if err != nil{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.RequestValidateFail,
		}

		json.NewEncoder(res).Encode(responseObject)

		return

	}

	// validating the request fields
	requestVal := RequestValidation.SignInValidate(pojoObj)

	if !requestVal.Status{

		json.NewEncoder(res).Encode(requestVal)

		return
	}

	// converting password to sha256

	// setting default password for admin
	hash := sha256.New()

	hash.Write([]byte(*pojoObj.Password))

	*pojoObj.Password = hex.EncodeToString(hash.Sum(nil))

	// checking if user Exists in database then create session and put the session in caching DB like Redis
	adminDetails, err := auth.SignIn(pojoObj)

	if err != nil{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.InvalidUserIdPassword,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT signInClaims, which includes the username and expiry time
	signInClaims := objects.SignInClaims{
		Email: adminDetails["email"].(string),
		StandardClaims: jwt.StandardClaims{

			// In JWT, the expiry time is expressed as unix milliseconds

			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the signInClaims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, signInClaims)

	// Create the JWT string
	tokenString, err := token.SignedString(objects.SignInjwtKey)

	if err != nil {

		responseObject := objects.ErrorResponse{
			Status:false,
			Msg:messages.ErrorMessage,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	jsonString, err := json.Marshal(adminDetails)

	// if err then error is thrown
	if err != nil{

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.ErrorMessage,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	ctx := context.Background()

	// inserting into redis
	redisErr := RedisConnection.Client().Set(ctx, tokenString, string(jsonString), 0).Err()

	if redisErr != nil {

		responseObject := &objects.ErrorResponse{
			Status:false,
			Msg:messages.ErrorMessage,
		}

		json.NewEncoder(res).Encode(responseObject)

		return
	}

	// setting expire of the token for 24 hours
	RedisConnection.Client().Expire(ctx, tokenString, time.Hour * 24)

	// success response after successfully adding new record to database
	responseObject := &objects.AuthSuccessDataResponse{
		Status:true,
		AccessToken: tokenString,
	}

	json.NewEncoder(res).Encode(responseObject)
}
