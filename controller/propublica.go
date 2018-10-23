package controller

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func GetLegislators() {
	response, _ := http.Get("http://pokeapi.co/api/v2/pokedex/kanto")
	responseData, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(responseData)) 
}

func PropublicaTest() {
	fmt.Println("hello from PropublicaTest")
}