package main

import (
	"fmt"
	"time"

	"github.com/expr-lang/expr"
)

type Env struct {
	Posts []Post `expr:"posts"`
}

func (Env) Format(t time.Time) string { // Methods defined on the struct become functions.
	return t.Format(time.RFC822)
}

type Post struct {
	Body string
	Date time.Time
}

type RR struct {
	Value []interface{} `expr:"value"`
}

func main() {
	code := `value[0] + 1`
	program, err := expr.Compile(code) // Pass the struct as an environment.
	if err != nil {
		panic(err)
	}

	d := []interface{}{11, 23, 4}
	// 打印d的类型
	r := RR{
		Value: d,
	}
	output, err := expr.Run(program, r)
	if err != nil {
		panic(err)
	}

	fmt.Print(output)
}
