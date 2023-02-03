// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"
// 	"text/template"
// )

// type Album struct {
// 	Title  string
// 	Genres Genres
// }
// type Genres struct {
// 	Data Tabdata
// }
// type Tabdata []struct {
// 	Id int
// }
// type Artiste struct {
// 	Id             int
// 	Name           string
// 	Link           string
// 	Share          string
// 	Picture        string
// 	Picture_small  string
// 	Picture_medium string
// 	Picture_big    string
// 	Picture_xl     string
// 	Nb_album       int
// 	Nb_fan         int
// 	Radio          bool
// 	Tdracklist     string
// }

// type Variable struct { //structure pour le hangman
// 	Entrer string
// }

// func Home(w http.ResponseWriter, r *http.Request, Variable *Variable, Artiste *Artiste) { //page d'acceuil
// 	template, err := template.ParseFiles("./index.html", "templates/test.html", "templates/forms.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	template.Execute(w, Artiste)
// }

// func api(Artiste *Artiste, Variable *Variable) {
// 	url, err := http.Get("https://api.deezer.com/artist/" + Variable.Entrer)
// 	if err != nil {
// 		os.Exit(1)
// 	}
// 	Arstiste_json, err := ioutil.ReadAll(url.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf(string(Arstiste_json))
// 	json.Unmarshal(Arstiste_json, Artiste)
// }

// func User(w http.ResponseWriter, r *http.Request, Variable *Variable, Artiste *Artiste) { // page de l'entrée utilisateur
// 	tmpl := template.Must(template.ParseFiles("./templates/forms.html"))
// 	if r.Method != http.MethodPost {
// 		tmpl.Execute(w, nil)
// 		return
// 	}
// 	a := r.FormValue("entrer")
// 	if len(a) > 0 {
// 		Variable.Entrer = a
// 		api(Artiste, Variable)
// 	}
// 	tmpl.Execute(w, struct{ Success bool }{true})

// }

// func main() { // fonction main

// 	var Artiste *Artiste = &Artiste{}
// 	var Variable *Variable = new(Variable)

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //page d'acceuil
// 		User(w, r, Variable, Artiste)
// 		Home(w, r, Variable, Artiste)
// 	})
// 	fs := http.FileServer(http.Dir("./static/"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))
// 	fi := http.FileServer(http.Dir("./hangman/assets/"))
// 	http.Handle("/hangman/assets/", http.StripPrefix("/hangman/assets/", fi))
// 	http.ListenAndServe(":8000", nil)
// }
