package objects

import(
	"github.com/dgrijalva/jwt-go"
)

type SignIn struct{
	Email *string `json:"email"`
	Password *string `json:"password"`
}

type SignInClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Create the JWT key used to create the signature
var SignInjwtKey = []byte("coynce_secret_key_12345567890!@#$%^&*()")

type AuthSuccessDataResponse struct{
	Status bool `json:"status"`
	AccessToken string `json:"access_token"`
}

type JsonLogOut struct {
	Msg    string `json:"msg"`
	Status bool `json:"status"`
	Logout bool `json:"logout"`
}
