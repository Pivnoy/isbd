package models

type ReptiloidsResponseAll struct {
	Tv         float32  `json:"tv"`
	Business   float32  `json:"business"`
	Science    float32  `json:"science"`
	Countries  []string `json:"countries"`
	CountriesP float32  `json:"countries_p"`
}

type RepriloidResponseCountry struct {
	Tv       float32 `json:"tv"`
	Business float32 `json:"business"`
	Science  float32 `json:"science"`
}
