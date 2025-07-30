package main

import (
	"devinggo/manage/model/common"
	"fmt"
)

func main() {
	// env := map[string]interface{}{
	// 	"greet":   "Hello, %v!",
	// 	"names":   []string{"world", "you"},
	// 	"sprintf": fmt.Sprintf,
	// }

	// code := `sprintf(greet, names[0])`

	// program, err := expr.Compile(code, expr.Env(env))
	// if err != nil {
	// 	panic(err)
	// }

	// output, err := expr.Run(program, env)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(output)
	c := common.TemplateEnv{}
	c.Value.Value = []uint16{9}
	fmt.Println(c.Value.ToValue())
}
