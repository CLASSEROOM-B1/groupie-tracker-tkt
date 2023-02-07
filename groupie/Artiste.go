package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Detail struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func AllARtiste() []Detail {

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

	fmt.Println(len(m))

	return m
}
