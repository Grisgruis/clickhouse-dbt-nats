package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

type Msg struct {
	Source string `json:"source"`
	Value  string `json:"value"`
}

func main() {
	log.Printf("Booting")

	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Printf("%s", err)
	}
	defer nc.Close()

	// Fork connection so we can publish json/structs directly
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	// Simple async subscriber
	sub, err := ec.Subscribe("foo", func(msg *nats.Msg) {
		log.Printf("Received message: %s", string(msg.Data))
	})
	if err != nil {
		log.Printf("%s", err)
	}
	defer sub.Unsubscribe()

	// Create a msg
	msg := Msg{
		Source: "golang",
		Value:  "Value from Golang",
	}

	// Simple publisher
	if err := ec.Publish("foo", msg); err != nil {
		log.Printf("%s", err)
	}

	// Keep alive
	time.Sleep(180 * time.Second)
}
