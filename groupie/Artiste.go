package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Detail struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func Api(a string) *Detail {

	var Details *Detail = new(Detail)

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

	Recherche(a, m, Details)

	return (Details)
}

func Recherche(userInput string, m []Detail, Detail *Detail) {

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

		json.Unmarshal(Arstiste_json, Detail)
		fmt.Println(Detail.Name)

	}

}
