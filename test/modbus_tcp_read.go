package main

import (
	"fmt"
	"log"
	"time"

	"github.com/simonvetter/modbus"
)

func main() {
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://127.0.0.1:5020",
		Timeout: 3 * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to create Modbus client: %v", err)
	}

	if err = client.Open(); err != nil {
		log.Fatalf("Failed to open Modbus client: %v", err)
	}

	client.SetUnitId(1)

	rs, err := client.ReadRegisters(1, 1, modbus.HOLDING_REGISTER)
	if err != nil {
		log.Fatalf("Failed to read registers: %v", err)
	}

	fmt.Println(rs)

}
