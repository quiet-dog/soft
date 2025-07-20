package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/InfluxCommunity/influxdb3-go/v2/influxdb3"
)

func main() {
	client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:     "http://localhost:8181",
		Token:    "apiv3_m5pZL1Z_fuVx4oEKwkwSiL5qyIYu3CQrih5394FoDuURdYxPqwtWO3IYiZG06-0AXysYINo_f46Pi5-xDQa-pw",
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
	// line := "stat,location=Paris temperature=23.5,humidity=45i"
	// line := "t_56,sensor=s_48 c_49=55.390000,e_48=55.390000"
	// err = client.Write(context.Background(), []byte(line))
	// fmt.Println(err)
	// 查询
	query := "SELECT * FROM t_56 where c_49 is not null"
	res, err := client.Query(context.Background(), query)
	if err != nil {
		fmt.Println("Query error:", err)
		return
	}
	for res.Next() {
		// The query iterator returns each row as a map[string]interface{}.
		// The keys are the column names, allowing you to access the values by column name.
		value := res.Value()
		n, _ := json.Marshal(value)
		fmt.Println("Query result:", string(n))
	}
}
