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

type Vide struct {
	Index []DetailLocation
}

type DetailLocation struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Continent []int
}

func ApiLocation(a int) *DetailLocation {

	var Detailss *DetailLocation = new(DetailLocation)
	var Vide *Vide = new(Vide)

	url, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		os.Exit(1)
	}
	Arstiste_json, err := ioutil.ReadAll(url.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(Arstiste_json, Vide)

	test(Vide, Detailss, a)

	return Detailss
}

func Tab() *Vide {

	var vide *Vide = new(Vide)

	i := 0

	url, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		os.Exit(1)
	}
	Arstiste_json, err := ioutil.ReadAll(url.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(Arstiste_json, vide)

	fmt.Println(len(vide.Index))

	for i := 0; i < len(vide.Index); i++ {
		vide.Index[i].Continent = Cont(vide.Index[i].Locations)

		fmt.Println(i + 1)
	}

	i = i

	return vide

}

func test(vide *Vide, Detailss *DetailLocation, id int) {

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

	x := Cont(Detailss.Locations)

	fmt.Println(x)

	Detailss.Continent = x

}
