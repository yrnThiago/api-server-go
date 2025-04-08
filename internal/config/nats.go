package config

import (
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

var MsgChan chan *nats.Msg

func NatsInit() {
	opts := &server.Options{}
	ns, err := server.NewServer(opts)
	if err != nil {
		panic(err)
	}
	go ns.Start()

	if !ns.ReadyForConnections(4 * time.Second) {
		panic("not ready for connection")
	}

	MsgChan = make(chan *nats.Msg)

	// pub.PublisherInit()
	// sub := sub.Connect()
	//
	// go sub.ReceiveMessage(msgChan, Env.NEW_ORDERS_TOPIC)
	//
	// for msg := range msgChan {
	// 	var order *models.Order
	//
	// 	err = json.Unmarshal(msg.Data, &order)
	// 	if err != nil {
	// 		return
	// 	}
	//
	// 	fmt.Println(order)
	// }
}
