package main

import (
	"fmt"
	"groupie/groupie"
	"log"
	"net/http"
	"strconv"
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
	image  int
	choix  int
}

type Liste struct {
	Simple groupie.DetailSimple
	Detail groupie.Detail
}

func Home(w http.ResponseWriter, r *http.Request, Variable *Variable, DetailSimple *groupie.DetailSimple) { //page d'acceuil
	template, err := template.ParseFiles("./index.html", "./templates/header.html", "templates/forms.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, DetailSimple)
}

func Artiste(w http.ResponseWriter, r *http.Request, Variable *Variable, DetailSimple *groupie.DetailSimple) { //page d'acceuil
	template, err := template.ParseFiles("./pages/artiste.html", "templates/seul.html", "templates/basdepage.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, DetailSimple)
}

func AllArtise(w http.ResponseWriter, r *http.Request, Variable *Variable, Detail *[]groupie.Detail) { //page d'acceuil
	template, err := template.ParseFiles("./pages/AllArtiste.html", "./templates/header.html", "templates/test.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, Detail)
}

func Map(w http.ResponseWriter, r *http.Request, Variable *Variable, loc string) { //page d'acceuil

	templ, err := template.ParseFiles("./pages/map.html", "./templates/header.html")
	if err != nil {
		fmt.Println("Error = ", err)
	}
	err = templ.Execute(w, loc)
	if err != nil {
		fmt.Println("Error = ", err)
	}
}

func User(w http.ResponseWriter, r *http.Request, Variable *Variable, Detail *[]groupie.Detail, DetailLocation *DetailLocation, Date *Date, DetailSimple groupie.DetailSimple) (groupie.DetailSimple, *DetailLocation, int) { // page de l'entrée utilisateur
	tmpl := template.Must(template.ParseFiles("./templates/forms.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
	}
	choix := 0
	a := r.FormValue("entrer")

	if len(a) > 0 {
		Variable.Entrer = a
		DetailSimple = groupie.Api(a, 0, 0)
		choix = 1

		tempLocation := groupie.ApiLocation(DetailSimple.Id)

		DetailLocation.Id = tempLocation.Id
		DetailLocation.Locations = tempLocation.Locations

		tempDate := groupie.ApiDate(DetailSimple.Id)

		Date.Id = tempDate.Id
		Date.Dates = tempDate.Dates

	}
	tmpl.Execute(w, struct{ Success bool }{true})

	return DetailSimple, DetailLocation, choix

}

func ImageArtiste(w http.ResponseWriter, r *http.Request, Variable *Variable, Detail []groupie.Detail, DetailLocation *DetailLocation, Date *Date, DetailSimple groupie.DetailSimple, lenDetail int) ([]groupie.Detail, groupie.DetailSimple, int) { // page de l'entrée utilisateur
	tmpl := template.Must(template.ParseFiles("./templates/test.html"))
	b := 0
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
	}

	Detail = groupie.AllARtiste()

	DetailFull := Detail
	var DetailTest []groupie.Detail

	memberString := r.FormValue("member")
	member, _ := strconv.Atoi(memberString)

	dateString := r.FormValue("date")
	date, _ := strconv.Atoi(dateString)

	albumString := r.FormValue("album")
	album, _ := strconv.Atoi(albumString)

	DetailFull = groupie.Filtre(Detail, DetailFull, DetailTest, member, date, album)

	choix := 0
	a := r.FormValue("artisteImage")
	b, _ = strconv.Atoi(a)

	if b > 0 {
		Variable.image = b
		DetailSimple = groupie.Api(a, 1, b)
		choix = 1
		b = 0

		tempLocation := groupie.ApiLocation(DetailSimple.Id)

		DetailLocation.Id = tempLocation.Id
		DetailLocation.Locations = tempLocation.Locations

		tempDate := groupie.ApiDate(DetailSimple.Id)

		Date.Id = tempDate.Id
		Date.Dates = tempDate.Dates

	}
	tmpl.Execute(w, struct{ Success bool }{true})

	return DetailFull, DetailSimple, choix
}

func main() { // fonction main
	var Variable *Variable = new(Variable)
	var DetailLocation *DetailLocation = new(DetailLocation)
	var Date *Date = new(Date)
	// var Liste *Liste = new(Liste)

	Detail := groupie.AllARtiste()
	DetailSimple := groupie.Api(Variable.Entrer, 0, 0)
	choix := 0
	choix2 := 0
	lenDetail := len(Detail)

	var loc string

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //page d'acceuil
		DetailSimple, DetailLocation, choix = User(w, r, Variable, &Detail, DetailLocation, Date, DetailSimple)
		// DetailSimple := groupie.Api(Variable.Entrer, 0, 0)
		// Detail, DetailSimple, choix2 = ImageArtiste(w, r, Variable, Detail, DetailLocation, Date, DetailSimple, lenDetail)

		if choix == 1 {
			http.Redirect(w, r, "/artiste", http.StatusSeeOther)
		}
		Home(w, r, Variable, &DetailSimple)
	})

	http.HandleFunc("/AllArtiste", func(w http.ResponseWriter, r *http.Request) { //page d'acceuil
		Detail, DetailSimple, choix2 = ImageArtiste(w, r, Variable, Detail, DetailLocation, Date, DetailSimple, lenDetail)
		if choix2 == 1 {
			http.Redirect(w, r, "/artiste", http.StatusSeeOther)
		}
		AllArtise(w, r, Variable, &Detail)
	})

	http.HandleFunc("/artiste", func(w http.ResponseWriter, r *http.Request) { //page d'acceu
		Artiste(w, r, Variable, &DetailSimple)

	})

	http.HandleFunc("/map", func(w http.ResponseWriter, r *http.Request) { //page d'acceu
		loc = groupie.Cord(DetailLocation.Locations, Date.Dates)
		Map(w, r, Variable, loc)

	})

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8000", nil)
}
