package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func main() {
	mux := http.NewServeMux()
	sm := NewSessionManage()

	mux.HandleFunc("/api/chunk-upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			cookie, err := r.Cookie("session_upload")
			// check session cookie uploads
			if err == http.ErrNoCookie {
				// decode or parse body request
				var payload PayloadGET
				err := json.NewDecoder(r.Body).Decode(&payload)
				if err != nil {
					http.Error(w, "Failed to parse JSON: "+err.Error(), http.StatusBadRequest)
					return
				}
				// generate the session id
				data := map[string]interface{}{
					"chunk_size":     payload.ChunkSize,
					"file_size":      payload.FileSize,
					"chunk_num":      payload.ChunkNum,
					"chunk_required": map[string]interface{}{},
				}
				sessionID, err := sm.CreateSession(data)
				if err != nil {
					http.Error(w, "Failed to create session: "+err.Error(), http.StatusBadRequest)
					return
				}
				// create cookie
				newCookie := &http.Cookie{
					Name:    "session_upload",
					Value:   sessionID,
					Expires: time.Now().Add(30 * time.Minute),
				}
				// set cookie
				http.SetCookie(w, newCookie)

				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")

				response := Response{
					Status:  "success",
					Message: "session is created",
					Data: map[string]interface{}{
						"session_id": sessionID,
						"chunk_size": payload.ChunkSize,
						"file_size":  payload.FileSize,
						"chunk_num":  payload.ChunkNum,
						"file_name":  payload.FileName,
						// "chunk_required" :
					},
				}

				if err := json.NewEncoder(w).Encode(response); err != nil {
					http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)
				}

				return
			}

			session, exists := sm.getSession(cookie.Value)
			if !exists {
				http.Error(w, fmt.Sprintf("Session not exists"), http.StatusInternalServerError)
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")

			response := Response{
				Status:  "success",
				Message: "session is created",
				Data:    session.Data,
			}

			if err := json.NewEncoder(w).Encode(response); err != nil {
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

type PayloadGET struct {
	FileSize  int    `json:"file_size"`
	ChunkSize int    `json:"chunk_size"`
	ChunkNum  int    `json:"chunk_num"`
	FileName  string `json:"file_name"`
}
