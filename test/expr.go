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

func main() {
	code := `value + 1`
	program, err := expr.Compile(code) // Pass the struct as an environment.
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, map[string]interface{}{"value": 1})
	if err != nil {
		panic(err)
	}

	fmt.Print(output)
}
