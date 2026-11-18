package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sais-gallego/shoplite-backend-api/api"
)

func main() {
	server := api.NewServer()

	http.HandleFunc("/api/v1/cart", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			server.HandleGetCart(w, r)
		} else if r.Method == http.MethodPost {
			server.HandleAddItem(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/v1/cart/checkout", server.HandleCheckout)

	port := ":8080"
	fmt.Printf("🚀 Shoplite Backend API Server listening on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
