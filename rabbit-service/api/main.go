package api

import (
	"encoding/json"
	"fmt"
	"github.com/perbu/go-intro/rabbit-service/rabbits"
	"log"
	"net/http"
)


// Response struct.
type helloResponse struct {
	Message string `json:"message"`
}

func Run() {
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Listen() failed: %s", err)
	}
}

func helloHandler(response http.ResponseWriter, request *http.Request) {
	var msg string
	var rabbitType rabbits.Rabbit

	rabbitTypeStr, ok := request.URL.Query()["type"] // Use Gorilla Mux, if you'd prefer a router schema.
	if ok {
		// We ignore any subsequent type of rabbit.
		rabbitType, _ = rabbits.StringToRabbit(rabbitTypeStr[0])
		msg = fmt.Sprintf("Hello Rabbit! Your type (%d) is as string: %s", rabbitType, rabbitType)
	} else {
		msg = "Hello World! No rabbits here."
	}
	// and stick it in the message
	m := helloResponse{Message: msg}
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		http.Error(response, "could not marshal json", http.StatusInternalServerError)
		return // Common error is to forget to return here.
	}
	response.Header().Set("content-type", "application/json")
	response.WriteHeader(http.StatusOK)
	_, err = response.Write(jsonBytes)
	if err != nil {
		log.Printf("could not write to socket in helloHandler: %s", err)
	}
	return
}
