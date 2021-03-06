package main

import (
	"encoding/json"
	"fmt"
	"github.com/Pivnoy/isbd/server"
	"github.com/Pivnoy/isbd/service"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
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

/* обработчики пока что тут, в дальнейшем они будут в handlers.go */

// GetSectsHandler  '/sect'
func GetSectsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response, err := service.GetSects()
	if err != nil {
		panic("Error in serv Jobless")
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

// GetPunitiveHandler '/punitive'
func GetPunitiveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response, err := service.GetPunitive()
	if err != nil {
		panic("Error in serv Jobless")
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

// GetJoblessCountryHandler   ('/stats/job?country_name=<name>') если нет параметров, то передает управление GetJoblessAllHandler*/
func GetJoblessCountryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := r.URL.Query()
	countryName := params.Get("country_name")
	if countryName == "" {
		GetJoblessAllHandler(w, r)
	} else {
		result, err := service.GetJoblessCountry(countryName)
		if err != nil {
			panic("Error in serv Jobless Country")
		}
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			return
		}
	}
}

// GetJoblessAllHandler вызывается при отсутсвии параметров у запроса на '/stats/job'
func GetJoblessAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result, err := service.GetAllJobless()
	if err != nil {
		panic("Error in serv Jobless")
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}

func GetCrazyCountryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := r.URL.Query()
	countryName := params.Get("country_name")
	if countryName == "" {
		GetCrazyAllHandler(w, r)
	} else {
		result, err := service.GetCrazyCountry(countryName)
		if err != nil {
			panic("Error in serv Crazy Country")
		}
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			return
		}
	}
}

func GetCrazyAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result, err := service.GetCrazyAll()
	if err != nil {
		panic("Error in serv Crazy")
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}

func GetReptiloidsCountryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := r.URL.Query()
	countryName := params.Get("country_name")
	if countryName == "" {
		GetReptiloidsAllHandler(w, r)
	} else {
		result, err := service.GetCountryReptiloids(countryName)
		if err != nil {
			panic("Error in serv Reptil")
		}
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			return
		}
	}
}

func GetReptiloidsAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result, err := service.GetAllReptiloids()
	if err != nil {
		panic("Error in serv Reptil")
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}

// GetPunitiveHandler '/punitive'
func GetHumanHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response, err := service.GetHuman()
	if err != nil {
		panic("Error in serv Jobless")
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

// GetPunitiveHandler '/punitive'
func GetReptiloids(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response, err := service.GetReptioids()
	if err != nil {
		panic("Error in serv Jobless")
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func DoMorphHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := r.URL.Query()
	humanId, err := strconv.Atoi(params.Get("human_id"))
	if err != nil {
		panic("error of parse to int human id")
	}
	reptiloidId, err := strconv.Atoi(params.Get("reptiloid_id"))
	if err != nil {
		panic("error of parse to int reptiloid id")
	}
	err = service.DoMorth(humanId, reptiloidId)
	if err != nil {
		panic("Error in serv doMorth")
	}
	w.WriteHeader(http.StatusOK)
}

func DoAtackSectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := r.URL.Query()
	sectName := params.Get("sect_name")
	punitiveName := params.Get("punitive_name")
	err := service.DoAttackSect(sectName, punitiveName)
	if err != nil {
		panic("Erroe in serv do Attck")
	}
	w.WriteHeader(http.StatusOK)

}

func GetCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response, err := service.GetCountry()
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
	router.HandleFunc("/sect", GetSectsHandler).Methods("GET")
	router.HandleFunc("/punitive", GetPunitiveHandler).Methods("GET")
	router.HandleFunc("/stats/job", GetJoblessCountryHandler).Methods("GET")
	router.HandleFunc("/stats/crazy", GetCrazyCountryHandler).Methods("GET")
	router.HandleFunc("/stats/reptiloids", GetReptiloidsCountryHandler).Methods("GET")
	router.HandleFunc("/human", GetHumanHandler).Methods("GET")
	router.HandleFunc("/reptiloid", GetReptiloids).Methods("GET")
	router.HandleFunc("/morph", DoMorphHandler).Methods("GET")
	router.HandleFunc("/attack_sect", DoAtackSectHandler).Methods("GET")
	router.HandleFunc("/country", GetCountry).Methods("GET")

	http.Handle("/", router)
	err := http.ListenAndServe(":7000", router)
	if err != nil {
		return
	}
}
