package main
/*
* This is the main function of the project, 
* This project contains all api services related to registration, pricefeed broadcast using websocket, session management
* Order placement apis, order book api, netposition api etc
*/

import(
	"cns"
	"lib"
	"MongoConnection"
	"RedisConnection"
	"fmt"
	"Web"
	"net/http"
)

func main(){

	//mongod --replSet "replicaSet" --dbpath D:\mongo_db\mongodb_new

	// handling panic in the function
	defer lib.HandlePanic()

	httpApp := cns.Http{}

	// Connect Redis
	if(!RedisConnection.Connect()){
		fmt.Println("Failed to connect Redis")
		return
	}

	// Connect Mongodb
	if(!MongoConnection.Connect()){
		fmt.Println("Failed to connect Mongodb")
		return
	}

	// initializing controller routes
	Web.Routes()

	// running http server
	go func(){
		defer cns.CreateHttpServer(":3500")
	}()

	// running https server on seperate port from http
	defer cns.CreateHttpsServer(":3600", "/ssl/pounze.crt","/ssl/pounze.key")

	// default header set for all http response
	defer httpApp.DefaultMethod(func(req *http.Request,res http.ResponseWriter){
		res.Header().Set("Name", "Sudeep Dasgupta")
		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	})
}