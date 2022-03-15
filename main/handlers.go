package main

import (
	"encoding/json"
	"github.com/Pivnoy/isbd/models"
	"net/http"
)

//func GetJoblessCountryHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-type", "application/json")
//	//params := r.URL.Query()
//	//country_name := params.Get("country_name")
//	var jobless models.Human
//
//	// do smth
//
//	err := json.NewEncoder(w).Encode(jobless)
//	if err != nil {
//		return
//	}
//}

//func GetJoblessAllHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-type", "application/json")
//	result, err := service.GetAllJobless()
//	if err != nil {
//		panic("Error in serv Jobless")
//	}
//	err = json.NewEncoder(w).Encode(result)
//	if err != nil {
//		return
//	}
//}

func GetReptiloidsCountryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	//params := r.URL.Query()
	//country_name := params.Get("country_name")
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
	//params := r.URL.Query()
	//country_name := params.Get("country_name")
	var country models.Country

	// do smth

	err := json.NewEncoder(w).Encode(country)
	if err != nil {
		return
	}
}

//func GetCrazyCountryHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-type", "application/json")
//	//params := r.URL.Query()
//	//country_name := params.Get("country_name")
//	var crazy models.Crazy
//
//	// do smth
//
//	err := json.NewEncoder(w).Encode(crazy)
//	if err != nil {
//		return
//	}
//}

//func GetCrazyAllHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-type", "application/json")
//	var crazy models.Crazy
//
//	// do smth
//
//	err := json.NewEncoder(w).Encode(crazy)
//	if err != nil {
//		return
//	}
//}

//func GetPunitiveHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-type", "application/json")
//	var punitive models.Punitive
//
//	// do smth
//
//	err := json.NewEncoder(w).Encode(punitive)
//	if err != nil {
//		return
//	}
//}

//func GetSectsHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-type", "application/json")
//	var sects models.Sects
//
//	// do smth
//
//	err := json.NewEncoder(w).Encode(sects)
//	if err != nil {
//		return
//	}
//}

// DoAttackSectHandler тут лёха в респонс ничего не ждет
func DoAttackSectHandler(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query()
	//sect_name := params.Get("sect_name")
	//punitive_name := params.Get("punitive_name")

	w.WriteHeader(http.StatusOK)
}

// DoMorphHandler тут тоже лёха в респонс ничего не ждет
func DoMorphHandler(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query()
	//human_id := params.Get("human_id")
	//rep_id := params.Get("rep_id")

	w.WriteHeader(http.StatusOK)
}
