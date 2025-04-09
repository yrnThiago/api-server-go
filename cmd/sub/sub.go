package sub

import (
	"log"

	"github.com/nats-io/nats.go"
)

type Subscriber struct {
	nc *nats.Conn
}

func Connect() *Subscriber {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	return &Subscriber{
		nc,
	}
}

func (s *Subscriber) ReceiveMessage(msgChan chan *nats.Msg, topic string) {
	_, err := s.nc.Subscribe(topic, func(msg *nats.Msg) {
		msgChan <- msg
	})
	if err != nil {
		log.Fatal(err)
	}

	// Keep the subscription active
	select {} // Block forever
}
