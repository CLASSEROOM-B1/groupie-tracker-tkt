package main

import (
	"fmt"
	"groupie/groupie"
	"log"
	"net/http"
	"text/template"
)

type DetailLocation struct {
	Id        int      `json:"id"`
	Locations []string `json:"location"`
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Variable struct { //structure pour le hangman
	Entrer string
}

func Home(w http.ResponseWriter, r *http.Request, Variable *Variable, DetailSimple *groupie.DetailSimple) { //page d'acceuil
	template, err := template.ParseFiles("./index.html", "templates/forms.html", "templates/seul.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(DetailSimple)
	template.Execute(w, DetailSimple)
}

func AllArtise(w http.ResponseWriter, r *http.Request, Variable *Variable, Detail *[]groupie.Detail) { //page d'acceuil
	template, err := template.ParseFiles("./pages/AllArtiste.html", "templates/test.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, Detail)
}

func User(w http.ResponseWriter, r *http.Request, Variable *Variable, Detail *[]groupie.Detail, DetailLocation *DetailLocation, DetailSimple groupie.DetailSimple) groupie.DetailSimple { // page de l'entrée utilisateur
	tmpl := template.Must(template.ParseFiles("./templates/forms.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
	}
	a := r.FormValue("entrer")

	if len(a) > 0 {
		Variable.Entrer = a
		DetailSimple = groupie.Api(a)
		fmt.Println(DetailSimple)

		// tempLocation := groupie.ApiLocation(groupie.Id)

		// DetailLocation.Id = tempLocation.Id
		// DetailLocation.Locations = tempLocation.Locations

		// tempDate := groupie.ApiDate(DetailSimple.Id)

		// DetailLocation.Id = tempDate.Id
		// DetailLocation.Locations = tempDate.Dates

	}
	tmpl.Execute(w, struct{ Success bool }{true})

	return DetailSimple

}

func main() { // fonction main
	var Variable *Variable = new(Variable)
	var DetailLocation *DetailLocation = new(DetailLocation)

	Detail := groupie.AllARtiste()
	DetailSimple := groupie.Api(Variable.Entrer)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //page d'acceuil
		DetailSimple = User(w, r, Variable, &Detail, DetailLocation, DetailSimple)
		Home(w, r, Variable, &DetailSimple)
	})

	http.HandleFunc("/AllArtiste", func(w http.ResponseWriter, r *http.Request) { //page d'acceuil
		AllArtise(w, r, Variable, &Detail)
	})

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fi := http.FileServer(http.Dir("./hangman/assets/"))
	http.Handle("/hangman/assets/", http.StripPrefix("/hangman/assets/", fi))
	http.ListenAndServe(":8000", nil)
}
