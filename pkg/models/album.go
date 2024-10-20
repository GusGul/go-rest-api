package models

type Album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Genre  string  `json:"genre"`
	Price  float64 `json:"price"`
}
