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

func Home(w http.ResponseWriter, r *http.Request, Variable *Variable, DetailSimple *groupie.DetailSimple) { //page d'acceuil
	template, err := template.ParseFiles("./index.html", "templates/forms.html")
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
	template, err := template.ParseFiles("./pages/AllArtiste.html", "templates/test.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, Detail)
}

func User(w http.ResponseWriter, r *http.Request, Variable *Variable, Detail *[]groupie.Detail, DetailLocation *DetailLocation, Date *Date, DetailSimple groupie.DetailSimple) (groupie.DetailSimple, int) { // page de l'entrée utilisateur
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

		fmt.Println(DetailLocation.Locations[0])

		tempDate := groupie.ApiDate(DetailSimple.Id)

		Date.Id = tempDate.Id
		Date.Dates = tempDate.Dates

		fmt.Println(Date.Dates[0])

	}
	tmpl.Execute(w, struct{ Success bool }{true})

	return DetailSimple, choix

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

	if member > 0 {

		for i := 0; i < len(Detail); i++ {
			if len((Detail)[i].Members) == member {
				fmt.Println((Detail)[i].Name)
				DetailTest = append(DetailTest, (Detail)[i])

			}
		}
		//fmt.Println(Detail)
		fmt.Println(len(DetailTest))
		DetailFull = DetailTest
	}

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

		fmt.Println(DetailLocation.Locations[0])

		tempDate := groupie.ApiDate(DetailSimple.Id)

		Date.Id = tempDate.Id
		Date.Dates = tempDate.Dates

		fmt.Println(Date.Dates[0])

	}
	tmpl.Execute(w, struct{ Success bool }{true})

	return DetailFull, DetailSimple, choix
}

func main() { // fonction main
	var Variable *Variable = new(Variable)
	var DetailLocation *DetailLocation = new(DetailLocation)
	var Date *Date = new(Date)

	Detail := groupie.AllARtiste()
	DetailSimple := groupie.Api(Variable.Entrer, 0, 0)
	choix := 0
	choix2 := 0
	lenDetail := len(Detail)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //page d'acceuil
		DetailSimple, choix = User(w, r, Variable, &Detail, DetailLocation, Date, DetailSimple)
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

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fi := http.FileServer(http.Dir("./hangman/assets/"))
	http.Handle("/hangman/assets/", http.StripPrefix("/hangman/assets/", fi))
	http.ListenAndServe(":8000", nil)
}
