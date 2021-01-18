package cns

var Config = map[string]string{
    "MAXPROCS":"MAX",
    "Mongodb":"mongodb://localhost:27017/?replicaSet=replicaSet",
    "Redis":"localhost:6379",
    "MongodbDatabase":"coynce",
    "RedirectURL": "https://koynce.com",
}

var Websocket = map[string]int{
	"ReadBufferSize": 512,
	"WriteBufferSize": 512,
	"HandShakeDuration": 5,
}