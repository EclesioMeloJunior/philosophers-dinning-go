package main

import (
	"fmt"

	"github.com/EclesioMeloJunior/go-concurrency/philosophers"
)

var n int

func main() {
	fmt.Println("Concurrency Golang")
	philosophers.Problem()
}
