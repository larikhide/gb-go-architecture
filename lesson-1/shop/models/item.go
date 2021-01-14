package models

type Item struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Price int32  `json:"price"`
}
