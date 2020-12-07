package main

import (
	"fmt"

	"github.com/denilson020898/go_say_hello/v2"
)

func main() {
	name := "Denilson"
	msg := "Denilson"
	fmt.Println(go_say_hello.SayHello(msg, name))
}
