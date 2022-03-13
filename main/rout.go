package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"GetJoblessCountry", "GET", "/stats/job/", GetJoblessCountryHandler},
	Route{"GetJoblessAll", "GET", "/stats/job", GetJoblessAllHandler},
	Route{"GetReptiloidsCountry", "GET", "/stats/reptiloids/", GetReptiloidsCountryHandler},
	Route{"GetReptiloidsAll", "GET", "/stats/reptiloids", GetReptiloidsAllHandler},
	Route{"GetCountry", "GET", "/stats/country/", GetCountryHandler},
}
