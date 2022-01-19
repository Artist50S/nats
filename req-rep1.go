package main

import (
	"encoding/json"

	"log"

	nats "github.com/nats-io/nats.go"
)

type User struct {
	Mail string `json:"mail"`
	Code string `json:"code"`
}

func main() {
	autorization := &User{
		Mail: "qwerty@ru",
		Code: "123456",
	}
	server := "nats://localhost:4222"
	natsConnection, _ := nats.Connect(server)
	log.Println("connect to", server)
	data, err := json.Marshal(&autorization)
	if err != nil {
		log.Fatal(err)
	}
	natsConnection.Publish("Discavery.OrderService2", data)
	log.Println("msg shipped: %v", string(data))

}
