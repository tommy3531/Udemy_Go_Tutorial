package model

import (
	"fmt"
)

type Response struct {
	Status string		`json:"status"`
	Copyright string	`json:"copyright"`
}

func PropublicaModel() {
	fmt.Println("PropublicaModel")
}