package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type index struct {
	Index []Date
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

func ApiDate(a int) *Date {

	var Detailss *Date = new(Date)
	var vide *index = new(index)

	url, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		os.Exit(1)
	}
	Arstiste_json, err := ioutil.ReadAll(url.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(Arstiste_json, vide)

	SearchDate(vide, Detailss, a)

	return (Detailss)
}

func SearchDate(vide *index, Detailss *Date, id int) {

	id_string := strconv.Itoa(id)

	test := "https://groupietrackers.herokuapp.com/api/dates/" + id_string

	url, err := http.Get(test)
	if err != nil {
		os.Exit(1)
	}
	Arstiste_json, err := ioutil.ReadAll(url.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(Arstiste_json, Detailss)

	fmt.Println(Detailss.Dates)
}
