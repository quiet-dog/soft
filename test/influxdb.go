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
	// line := "stat,location=Paris temperature=23.5,humidity=45i"
	// line := "t_56,sensor=s_48 c_49=55.390000,e_48=55.390000"
	// err = client.Write(context.Background(), []byte(line))
	// fmt.Println(err)
	// 查询
	// query := `SELECT AVG(e_48) FROM t_56 WHERE sensor IN ('s_48') AND 'e_48' IS NOT NULL AND 'c_48' IS NOT NULL GROUP BY time('1h')`
	query := `SELECT 
  TIME_BUCKET('1 hour', time) AS bucket,
  MEAN(field) AS avg_field
FROM t_56
WHERE
  time >= '2022-01-01T00:30:00Z'
  AND time <= '2026-01-01T01:30:00Z'
GROUP BY bucket
ORDER BY bucket`
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
