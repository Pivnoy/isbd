package models

type Reptiloids struct {
	tv          float32  `json:"tv"`
	business    float32  `json:"business"`
	science     float32  `json:"science"`
	countries   []string `json:"countries"`
	countries_p int      `json:"countries_p"`
}
