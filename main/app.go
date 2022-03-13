package main

import (
	"fmt"
	"github.com/Pivnoy/isbd/server"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
)

func BasicHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Some hadler")
	if err != nil {
		return
	}
}

func prepareArgs(arg []string) {
	for i := 0; i < len(arg); i++ {
		arg[i] = strings.Trim(arg[i], " ")
	}
}

func main() {
	initArgs := os.Args[1:]
	prepareArgs(initArgs)
	server.Init(initArgs)
	router := mux.NewRouter()
	router.HandleFunc("/", BasicHandler)
	http.Handle("/", router)
	err := http.ListenAndServe(":7000", router)
	if err != nil {
		return
	}
}
