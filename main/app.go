package main

import (
	"encoding/json"
	"fmt"
	"github.com/Pivnoy/isbd/server"
	"github.com/Pivnoy/isbd/service"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
)

func BasicHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "I sosal Menya Ebali")
	if err != nil {
		return
	}
}

func prepareArgs(arg []string) {
	for i := 0; i < len(arg); i++ {
		arg[i] = strings.Trim(arg[i], " ")
	}
}

func GetJoblessAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	response, err := service.GetAllJobless()
	if err != nil {
		panic("Error in serv Jobless")
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}

}

func main() {
	initArgs := os.Args[1:]
	prepareArgs(initArgs)
	server.Init(initArgs)

	router := mux.NewRouter()
	router.HandleFunc("/", BasicHandler)
	router.HandleFunc("/stats/job", GetJoblessAllHandler)
	http.Handle("/", router)
	err := http.ListenAndServe(":7000", router)
	if err != nil {
		return
	}
}
