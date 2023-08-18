//go:build !grpc

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"packer/config"
	"packer/packer"
)

// HTTPServer handles HTTP requests for the packer API.
type HTTPServer struct {
	calculator packer.Packer
}

// NewHttpServer initialize new HTTPServer
func NewHttpServer(calculator packer.Packer) *HTTPServer {
	return &HTTPServer{calculator: calculator}
}

type updatePackSizeRequest struct {
	Sizes []uint `json:"sizes"`
}

type calculatePacksRequest struct {
	Quantity uint `json:"quantity"`
}

type calculatePacksResponse struct {
	Packs map[uint]uint `json:"packs"`
}

// UpdatePackSizesHandler handles the HTTP request to update pack sizes.
func (s *HTTPServer) UpdatePackSizesHandler(w http.ResponseWriter, r *http.Request) {
	var sizesRequest updatePackSizeRequest
	err := json.NewDecoder(r.Body).Decode(&sizesRequest)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	s.calculator.UpdatePackSizes(sizesRequest.Sizes)
	w.WriteHeader(http.StatusOK)
}

// CalculatePacksHandler handles the HTTP request to calculate packs.
func (s *HTTPServer) CalculatePacksHandler(w http.ResponseWriter, r *http.Request) {
	var calculateRequest calculatePacksRequest
	err := json.NewDecoder(r.Body).Decode(&calculateRequest)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	result := s.calculator.CalculatePacks(calculateRequest.Quantity)
	response := calculatePacksResponse{Packs: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response) // TODO handle error
}

// StartServer starts new http server
func StartServer(cfg *config.Config) {
	server := NewHttpServer(packer.NewDefaultPackageCalculator(cfg.PacketSizes))
	http.HandleFunc("/calculate-packs", server.CalculatePacksHandler)
	http.HandleFunc("/update-pack-sizes", server.UpdatePackSizesHandler)
	fmt.Printf("Server listening on port %d...\n", cfg.HTTPPort)
	panic(http.ListenAndServe(fmt.Sprintf(":%d", cfg.HTTPPort), nil))
}
