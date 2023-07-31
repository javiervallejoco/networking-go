package main

import (
	"fmt"
	"ping"
)

func main() {
	s := ping.Send()
	fmt.Println(s)
}
