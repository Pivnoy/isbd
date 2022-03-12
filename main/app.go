package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func BasicHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Some hadler")
	if err != nil {
		return
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", BasicHandler)
	http.Handle("/", router)

	err := http.ListenAndServe(":7000", router)
	if err != nil {
		return
	}
}
