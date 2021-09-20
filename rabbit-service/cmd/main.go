package main

import (
	_ "embed"
	"fmt"
	"github.com/perbu/go-intro/rabbit-service/api"
)

//go:embed .version
var embeddedVersion string


func main() {
	fmt.Printf("Silly Rabbit Web Service version %s starting up\n", embeddedVersion)
	api.Run()
}
