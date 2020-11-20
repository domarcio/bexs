package main

import (
	"fmt"

	"github.com/domarcio/bexs/config"
)

func main() {
	fmt.Printf("Hello, cmd! Running on `%s` env\n", config.Env)
}
