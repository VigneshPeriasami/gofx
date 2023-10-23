package models

type Company struct {
	Ibans   string `json:"ibans"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Id      int    `json:"id"`
}
