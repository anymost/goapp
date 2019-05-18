package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)
var (
	redisClient *redis.Client
)

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func testPing() {
	pong  := redisClient.Ping()
	fmt.Println(pong.Result())
}

func testList() {
	err := redisClient.RPush("hello", 1, 2, 3, 4).Err()
	if err != nil {
		fmt.Println(err)
	}
	for redisClient.LLen("hello").Val() != 0 {
		val := redisClient.LPop("hello").Val()
		fmt.Println(val)
	}
}

func testString() {
	redisClient.Set("hello", "world", time.Minute)
	val := redisClient.Get("hello").Val()
	fmt.Println(val)
}

func testHash() {
	redisClient.HSet("person", "name", "jack")
	val := redisClient.HGet("person", "name").Val()
	fmt.Println(val)
	length := redisClient.HLen("person").Val()
	fmt.Println(length)
}

func testSet()  {
	redisClient.SAdd("set", "hello", "world")
	val := redisClient.SPop("set").Val()
	fmt.Println(val)
}


func main() {
	testSet()
}