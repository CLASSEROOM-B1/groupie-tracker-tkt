package groupie

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type DetailSimple struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func Api(a string, choix int, id int) DetailSimple {

	var DetailSimple *DetailSimple = new(DetailSimple)

	url, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		os.Exit(1)
	}
	Arstiste_json, err := ioutil.ReadAll(url.Body)
	if err != nil {
		log.Fatal(err)
	}

	var m []Detail
	if err := json.Unmarshal([]byte(Arstiste_json), &m); err != nil {
		panic(err)
	}

	if choix == 0 {
		Recherche(a, m, DetailSimple)
	} else if choix == 1 {
		RechercheImage(id, m, DetailSimple)
	}

	return *DetailSimple
}

func Recherche(userInput string, m []Detail, DetailSimple *DetailSimple) {

	id := 0
	boucle := false

	userInput = strings.ToLower(userInput)
	userInput = strings.Title(userInput)

	for _, val := range m {
		if val.Name == userInput {
			id = val.Id
			boucle = true
		}
	}

	if boucle == true {

		id_string := strconv.Itoa(id)

		test := "https://groupietrackers.herokuapp.com/api/artists/" + id_string

		url, err := http.Get(test)
		if err != nil {
			os.Exit(1)
		}
		Arstiste_json, err := ioutil.ReadAll(url.Body)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(Arstiste_json, DetailSimple)
	}

}

func RechercheImage(id int, m []Detail, DetailSimple *DetailSimple) {

	id_string := strconv.Itoa(id)

	test := "https://groupietrackers.herokuapp.com/api/artists/" + id_string

	url, err := http.Get(test)
	if err != nil {
		os.Exit(1)
	}
	Arstiste_json, err := ioutil.ReadAll(url.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(Arstiste_json, DetailSimple)

}
