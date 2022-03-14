package main

import (
	"encoding/json"
	"fmt"
	"github.com/Pivnoy/isbd/service"
	"github.com/gorilla/mux"
	"net/http"
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
	result, err := service.GetAllJobless()
	if err != nil {
		panic("Error in serv Jobless")
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}

func main() {
	//initArgs := os.Args[1:]
	//prepareArgs(initArgs)
	//server.Init(initArgs)
	//rows, err := server.DbInstance.Query("select id,number,name, theme from channel")
	//if err != nil {
	//	panic("")
	//}
	//chanColl := models.CreateChannelCollection(rows)
	//for _, v := range chanColl {
	//	fmt.Println(v)
	//}
	router := mux.NewRouter()
	router.HandleFunc("/", BasicHandler)
	//router.HandleFunc("/stats/job", GetJoblessAllHandler)
	http.Handle("/", router)
	err := http.ListenAndServe(":7000", router)

	if err != nil {
		return
	}
}
