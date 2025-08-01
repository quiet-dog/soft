package main

import (
	"fmt"

	"github.com/expr-lang/expr"
)

func main() {
	env := map[string]interface{}{
		"greet":   "Hello, %v!",
		"names":   []string{"world", "you"},
		"sprintf": fmt.Sprintf,
	}

	code := `names[0] == "world"`

	program, err := expr.Compile(code, expr.Env(env))
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	if v, ok := output.(bool); ok {
		fmt.Println(v)
	}

	// c := common.TemplateEnv{}
	// c.Value.Value = []uint16{9}
	// fmt.Println(c.Value.ToValue())
}
