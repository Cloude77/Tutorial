package main

import (
	"fmt"
)

var score = 99.5 // shared access

func main() {

	sayHello("Sergey")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}
