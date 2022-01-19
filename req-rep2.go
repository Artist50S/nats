package main

import (
	"encoding/json"

	"log"
	"runtime"

	nats "github.com/nats-io/nats.go"
)

type User struct {
	Mail string `json:"mail"`
	Code string `json:"code"`
}

func main() {
	autorization := &User{
		Mail: "qweqewqrty@ru",
		Code: "1234561223",
	}
	server := "nats://localhost:4222"
	natsConnection, err := nats.Connect(server)
	if err != nil {
		log.Fatal()
	} else {
		log.Println("Connected to", server)
	}
	natsConnection.Subscribe("Discavery.OrderService2", func(m *nats.Msg) {
		log.Println(string(m.Data))
		data, err := json.Marshal(&autorization)
		if err != nil {
			log.Fatal(err)
		}
		if err == nil {
			natsConnection.Publish(m.Reply, data)
			log.Println("msg shipped: %v", string(data))
		}
	})
	runtime.Goexit()
}
