package models

type Article struct {
	Type         string
	HarvesterID  string
	CerebroScore float64 `json:"cerebro-score"`
	URL          string
	Title        string
	CleanImage   string
}
