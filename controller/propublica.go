package controller

import (
	"fmt"
	"net/http"
	"encoding/json"
	"political/model"
	"os"
	"log"
)

func GetLegislators() {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.propublica.org/congress/v1/115/senate/members.json", nil)
	if err != nil {
		os.Exit(1)
	}
	req.Header.Add("X-API-KEY", "SpzjlPZlkMlPKKGCLQS1OqZtCN96lPl7sszOTKra")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	var responseObject model.Response
	if err := json.NewDecoder(resp.Body).Decode(&responseObject); err != nil {
		log.Println(err)
	}

	fmt.Println(responseObject.Status)
	fmt.Println(responseObject.Copyright)
}

func PropublicaTest() {
	fmt.Println("hello from PropublicaTest")
}