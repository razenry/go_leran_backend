package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Response struct {
	Message string `json:"message"`
}

type TodoResponse struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todo := TodoResponse{
				ID: 1,
				Title: "Belajar Go REST API",
				Description: "Belajar JSON response",
				Completed: false,
			}

			data, err := json.Marshal(todo)

			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)

			// fmt.Fprintln(w, "Todo List")
		case http.MethodPost:
			response := Response{
				Message: "Todo created",
			}

			data, err := json.Marshal(response)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write(data)
			// w.WriteHeader(http.StatusCreated)
			// fmt.Fprintln(w, "Todo Created")
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintln(w, "Method not allowed")
		}
	})

	fmt.Println("API server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}