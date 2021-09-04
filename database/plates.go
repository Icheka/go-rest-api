package database

import (
	"strconv"
)

// struct

// ID
// color
// price
// dimensions
// shape
// type
// material

type Plate struct {
	Id         string  `json:"id"`
	Color      string  `json:"color"`
	Price      float64 `json:"price"`
	Dimensions string  `json:"dimensions"`
	Shape      string  `json:"shape"`
	Type       string  `json:"type"`
	Material   string  `json:"material"`
}

var Plates []Plate // slice

func PopulatePlates() {
	for i := 1; i <= 10; i++ {
		Plates = append(Plates, Plate{
			Id:         strconv.Itoa(i),
			Color:      "red",
			Price:      90.4,
			Dimensions: "40x40in",
			Shape:      "round",
			Type:       "Bowl",
			Material:   "ceramic",
		})
	}
}
