package broker

import (
	"github.com/nats-io/nats"
	"log"
	"strconv"
)

var (
	Nc *nats.Conn
)

func Connect(host string, port int) error {
	/** Connect to NATS broker */
	var err error
	Nc, err = nats.Connect("nats://" + host + ":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalf("NATS: Can't connect: %v\n", err)
	}

	return err
}
