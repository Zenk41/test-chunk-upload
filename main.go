package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/api/chunk-upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			r.FormFile("chunk")
			w.WriteHeader(http.StatusOK)
			w.Header().Set( "Content-Type", "application/json")
	
			response := Response{
				Status: "success",
				Message: "this path has been accessed",
			}
	
			if err := json.NewEncoder(w).Encode(response);err != nil {
				http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)
			}
		} else if r.Method == "POST" {
			
		}
	})


	log.Printf("Server run on port 4000")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatalf("Cannot server the server on port :3000, err : %v ", err)
	}
 
}