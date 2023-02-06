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

type vide struct {
	Index []DetailLocation
}

type DetailLocation struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
}

func ApiLocation(a int) *DetailLocation {

	var Detailss *DetailLocation = new(DetailLocation)
	var vide *vide = new(vide)

	url, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		os.Exit(1)
	}
	Arstiste_json, err := ioutil.ReadAll(url.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(Arstiste_json, vide)

	test(vide, Detailss, a)

	return (Detailss)
}

func test(vide *vide, Detailss *DetailLocation, id int) {

	id_string := strconv.Itoa(id)

	test := "https://groupietrackers.herokuapp.com/api/locations/" + id_string

	url, err := http.Get(test)
	if err != nil {
		os.Exit(1)
	}
	Arstiste_json, err := ioutil.ReadAll(url.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(Arstiste_json, Detailss)

	fmt.Println(Detailss.Locations)
}
