package main

import (
	"net/http"
	"encoding/json"
	"log"
)

func respondWithError(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode > 499 {
		log.Println("Responding with 500 level error:", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, statusCode, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshal JSON respose: %v", payload)
		w.WriteHeader(500) // Internal Server error
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

// json.Marshal() is a function in Go's standard library 
// used to convert Go data structures (such as structs, 
// slices, maps, and primitive types) into a JSON-encoded representation

// import (
//     "encoding/json"
//     "fmt"
// )

// type Person struct {
//     FirstName string `json:"first_name"`
//     LastName  string `json:"last_name"`
//     Age       int    `json:"age"`
// }

// func main() {
//     person := Person{
//         FirstName: "John",
//         LastName:  "Doe",
//         Age:       30,
//     }

//     // Convert the Go struct to a JSON-encoded byte slice
//     jsonBytes, err := json.Marshal(person)
//     if err != nil {
//         fmt.Println("Error marshaling JSON:", err)
//         return
//     }

//     // Print the JSON-encoded data as a string
//     fmt.Println(string(jsonBytes))
// }
