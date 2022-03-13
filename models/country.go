package models

type Country struct {
	economics_gep int        `json:"economics_gep"`
	is_captured   bool       `json:"is_captured"`
	avg_salary    bool       `json:"avg_salary"`
	president     *President `json:"president"`
}
