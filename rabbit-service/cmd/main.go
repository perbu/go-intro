package main

import (
	"fmt"
	"github.com/perbu/go-intro/rabbit-service/api"
)


func main() {
	fmt.Println("Silly Rabbit Web Service starting up")
	api.Run()
}
