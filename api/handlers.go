package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sais-gallego/shoplite-backend-api/cart"
)

type Server struct {
	Cart *cart.Cart
}

func NewServer() *Server {
	return &Server{
		Cart: cart.NewCart(),
	}
}

func (s *Server) HandleGetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.Cart)
}

func (s *Server) HandleAddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var item cart.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.Cart.AddItem(item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s.Cart)
}

func (s *Server) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	discount, _ := strconv.ParseFloat(r.URL.Query().Get("discount"), 64)
	tax, _ := strconv.ParseFloat(r.URL.Query().Get("tax"), 64)

	summary := s.Cart.GetSummary(discount, tax)
	json.NewEncoder(w).Encode(summary)
}
