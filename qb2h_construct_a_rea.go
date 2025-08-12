package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// AppAnalysis struct to hold analysis results
type AppAnalysis struct {
	AppName     string `json:"app_name"`
	Version     string `json:"version"`
	Permissions []string `json:"permissions"`
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	NetworkUsage float64 `json:"network_usage"`
}

// AppAnalyzer struct to hold app analyzer logic
type AppAnalyzer struct{}

// NewAppAnalyzer returns a new instance of AppAnalyzer
func NewAppAnalyzer() *AppAnalyzer {
	return &AppAnalyzer{}
}

// AnalyzeApp analyzes a mobile app and returns the analysis results
func (a *AppAnalyzer) AnalyzeApp(appId string) (*AppAnalysis, error) {
	// TO DO: implement app analysis logic here
	// For demonstration purposes, return a sample analysis result
	return &AppAnalysis{
		AppName:     "QB2H App",
		Version:     "1.0.0",
		Permissions: []string{" INTERNET", " READ_EXTERNAL_STORAGE"},
		CPUUsage:    0.5,
		MemoryUsage: 100,
		NetworkUsage: 50,
	}, nil
}

func main() {
	a := NewAppAnalyzer()

	router := mux.NewRouter()

	router.HandleFunc("/analyze/{appId}", func(w http.ResponseWriter, r *http.Request) {
		appId := mux.Vars(r)["appId"]
		analysis, err := a.AnalyzeApp(appId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(analysis)
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}