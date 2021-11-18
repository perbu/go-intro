package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Response struct.
type helloResponse struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time""`
}

func (h helloResponse) getTimeFromHello() time.Time {
	return h.Time
}


func Run() {
	http.HandleFunc("/", myHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Listen() failed: %s", err)
	}
}

func myHandler(response http.ResponseWriter, request *http.Request) {

	m := helloResponse{
		Message: "Hello World!",
		Time:    time.Now(),
	}x
	log.Println(m.getTimeFromHello())
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
