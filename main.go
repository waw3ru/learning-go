package main

import (
	"fmt"

	"github.com/waw3ru/learning-go/utils"
)

func main() {
	s := utils.RunServer(utils.WithPort(2000), utils.WithHost("localhost"))
	p := utils.RunServer(utils.WithPort(2000))
	t := utils.RunServer(utils.WithHost("0.0.0.0"))
	fmt.Println(s)
	fmt.Println(p)
	fmt.Println(t)
}
