package RedisConnection

import(
	"cns"
	"github.com/go-redis/redis"
	"log"
	"lib"
	"context"
)

var client *redis.Client

func Connect() bool{

	defer lib.HandlePanic()

	ctx := context.Background()

	client = redis.NewClient(&redis.Options{
		Addr:cns.Config["RedisAddr"],
		Password:cns.Config["RedisAut"],
		DB:0,  
	})

	pong, err := client.Ping(ctx).Result()

	if err != nil{
		log.Println("Redis failed to connect")
		return false
	}

	log.Println(pong)

	return true
}

func Client() *redis.Client{

	defer lib.HandlePanic()

	return client
}