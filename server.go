package main

import (
	"groupie/groupie"
	"log"
	"net/http"
	"text/template"
)

type Detail struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

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

func Home(w http.ResponseWriter, r *http.Request, Variable *Variable, Detail *Detail) { //page d'acceuil
	template, err := template.ParseFiles("./index.html", "templates/test.html", "templates/forms.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, Detail)
}

func User(w http.ResponseWriter, r *http.Request, Variable *Variable, Detail *Detail, DetailLocation *DetailLocation) { // page de l'entrée utilisateur
	tmpl := template.Must(template.ParseFiles("./templates/forms.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	a := r.FormValue("entrer")
	if len(a) > 0 {
		Variable.Entrer = a
		temp := groupie.Api(a)

		Detail.Id = temp.Id
		Detail.Image = temp.Image
		Detail.Name = temp.Name
		Detail.Members = temp.Members
		Detail.CreationDate = temp.CreationDate
		Detail.FirstAlbum = temp.FirstAlbum

		tempLocation := groupie.ApiLocation(Detail.Id)

		DetailLocation.Id = tempLocation.Id
		DetailLocation.Locations = tempLocation.Locations

		tempDate := groupie.ApiDate(Detail.Id)

		DetailLocation.Id = tempDate.Id
		DetailLocation.Locations = tempDate.Dates

	}
	tmpl.Execute(w, struct{ Success bool }{true})

}

func main() { // fonction main
	var Variable *Variable = new(Variable)
	var Detail *Detail = new(Detail)
	var DetailLocation *DetailLocation = new(DetailLocation)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //page d'acceuil
		User(w, r, Variable, Detail, DetailLocation)
		Home(w, r, Variable, Detail)
	})
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fi := http.FileServer(http.Dir("./hangman/assets/"))
	http.Handle("/hangman/assets/", http.StripPrefix("/hangman/assets/", fi))
	http.ListenAndServe(":8000", nil)
}
