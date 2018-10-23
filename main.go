package main

import 
( 
	client "political/controller"
	model "political/model"
)

func main() {
	client.PropublicaTest()
	model.PropublicaModel()
	client.GetLegislators()
}