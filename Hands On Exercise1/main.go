package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)

	fmt.Println("Closing web server")
	http.ListenAndServe(":8080", nil)
}

type person struct {
	Name string `json:"name"`
}

func encode(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		Name: "Ayush",
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Failed to encode json and send to server", err)
	}
}

func decode(w http.ResponseWriter, r *http.Request) {
	var rb person
	if err := json.NewDecoder(r.Body).Decode(&rb); err != nil {
		log.Println("Failed to decode and receive from server", err)
	}
	log.Println(rb)
}
