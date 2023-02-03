package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Artiste struct {
	rs []Detail
}

type Detail struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func test() {

	var Details []string

	url, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		os.Exit(1)
	}
	Arstiste_json, err := ioutil.ReadAll(url.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(Arstiste_json, Details)

	var m []Detail
	if err := json.Unmarshal([]byte(Arstiste_json), &m); err != nil {
		panic(err)
	}
	for _, val := range m {
		fmt.Println(val.Id, val.Name)
	}

}

//create a function that take a json array and return a structure

func main() {
	test()
}
