package models

type Country struct {
	EconomicsGep int        `json:"economics_gep"`
	IsCaptured   bool       `json:"is_captured"`
	AvgSalary    float32    `json:"avg_salary"`
	President    *President `json:"president"`
}
