package main

import (
	"context"
	"fmt"

	"github.com/InfluxCommunity/influxdb3-go/v2/influxdb3"
)

func main() {
	client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:     "http://localhost:8181",
		Token:    "apiv3_KDfgrll4Hg3VKFpOt5wLOtStjWSmZNeIcW-obG1SYJGc5W2OAZRrH-pXq_5Q-_E7LT0bhKwcMOglg-Ml2J3EJg",
		Database: "DATABASE_NAME",
	})

	defer func(client *influxdb3.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)

	if err != nil {
		panic(err)
	}
	line := "1,sensor=2 value=23.5,current=45i"
	err = client.Write(context.Background(), []byte(line))
	fmt.Println(err)

}
