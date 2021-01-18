package middlewares

import(
	"RedisConnection"
	"context"
	"encoding/json"
	"net/http"
	"objects"
	"strings"
)

func ValidateSession(req *http.Request,res http.ResponseWriter, done chan bool){

	// checking for authorization in headers

	token := strings.Replace(req.Header.Get("Authorization"), "Bearer ", "", 1)

	// removing the Bearer text and trimming the string

	token = strings.TrimSpace(token)

	// creating context
	ctx := context.Background()

	// checking if the session key token is valid or not

	val, err := RedisConnection.Client().Get(ctx, token).Result()

	// if nil value is returned then logout = true and status = false

	if err != nil{
		respStruct := objects.JsonLogOut{}
		respStruct.Msg = "Invalid session id"
		respStruct.Status = false
		respStruct.Logout = true
		json.NewEncoder(res).Encode(respStruct)
		done <- false
		return
	}

	req.Header.Set("RedisSession", val)

	req.Header.Set("Token", token)

	// returning true

	done <- true
}