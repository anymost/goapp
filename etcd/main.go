package main

import (
	"go.etcd.io/etcd/client"
	"log"
	"time"
)

var (
	client *client.Client
)

func init()  {
	 config := client.Config{
		Endpoints:               []string{"http://127.0.0.1:2379"},
		Transport:               client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	 client, err = client.NewClient(config)
	 if err != nil {
	 	log.Fatal(err)
	 }
}

func main()  {
	kapi := client.NewKeysAPI(client)
	kapi.set
}
