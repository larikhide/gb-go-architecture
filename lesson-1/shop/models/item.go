package models

import "time"

type Item struct {
<<<<<<< HEAD
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
=======
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Price int32  `json:"price"`
>>>>>>> add GetItem test and change CreateItem test
}
