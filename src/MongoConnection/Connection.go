package MongoConnection

import(
	"cns"
	"context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "lib"
)

var db *mongo.Database

func Connect() bool{

	defer lib.HandlePanic()
	/*
		Setting URL for mongodb connections
	*/

	clientOptions := options.Client().ApplyURI(cns.Config["Mongodb"])

	// Connecting to Mongodb with server host localhost and port 27017

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil{
		log.Println("Failed to connect mongodb")
	    log.Println(err)
	    return false
	}

	/*
		Check the connection
	*/ 

	/*
		Checling for ping server if connected
	*/ 

	err = client.Ping(context.TODO(), nil)

	if err != nil{
	    log.Println(err)
	    return false
	}

	db = client.Database(cns.Config["MongodbDatabase"])

	log.Println("Connected to MongoDB! on : ",cns.Config["Mongodb"])

	return true
}

func Database() *mongo.Database{

	defer lib.HandlePanic()
	
	return db
}