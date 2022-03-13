package main

import (
	"encoding/json"
	"github.com/Pivnoy/isbd/models"
	"net/http"
)

func GetJoblessCountryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := r.URL.Query()
	country_name := params.Get("country_name")
	var jobless models.Jobless

	// do smth

	err := json.NewEncoder(w).Encode(jobless)
	if err != nil {
		return
	}
}

func GetJoblessAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var jobless models.Jobless

	// do smth

	err := json.NewEncoder(w).Encode(jobless)
	if err != nil {
		return
	}

}

func GetReptiloidsCountryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := r.URL.Query()
	country_name := params.Get("country_name")
	var reptiloids models.Reptiloids

	// do smth

	err := json.NewEncoder(w).Encode(reptiloids)
	if err != nil {
		return
	}
}

func GetReptiloidsAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var reptiloids models.Reptiloids

	// do smth

	err := json.NewEncoder(w).Encode(reptiloids)
	if err != nil {
		return
	}
}

func GetCountryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := r.URL.Query()
	country_name := params.Get("country_name")
	var country models.Country

	// do smth

	err := json.NewEncoder(w).Encode(country)
	if err != nil {
		return
	}
}
