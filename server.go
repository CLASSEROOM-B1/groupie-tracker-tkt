package main

import (
	"bufio"
	"fmt"
	"groupie/groupie"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type Donne struct {
	Detail []groupie.Detail

	DetailLocation *groupie.Vide

	DetailDate *groupie.Index

	DetailSimple *groupie.DetailSimple

	DetailFull []groupie.Detail

	Choix     int
	Choix2    int
	LenDetail int
	ApiOn     int
}

func Home(w http.ResponseWriter, r *http.Request) { //page d'acceuil
	template, err := template.ParseFiles("./pages/index.html", "templates/forms.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func Artiste(w http.ResponseWriter, r *http.Request, Donne *Donne) { //page d'acceuil
	template, err := template.ParseFiles("./pages/artiste.html", "templates/seul.html", "templates/basdepage.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, Donne.DetailSimple)
}

func AllArtise(w http.ResponseWriter, r *http.Request, Donne *Donne) { //page d'acceuil
	template, err := template.ParseFiles("./pages/AllArtiste.html", "./templates/header.html", "templates/test.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, Donne.DetailFull)
}

func Map(w http.ResponseWriter, r *http.Request, loc string) { //page d'acceuil

	templ, err := template.ParseFiles("./pages/map.html", "./templates/header.html")
	if err != nil {
		fmt.Println("Error = ", err)
	}
	err = templ.Execute(w, loc)
	if err != nil {
		fmt.Println("Error = ", err)
	}
}

func User(w http.ResponseWriter, r *http.Request, Donne *Donne) { // page de l'entrée utilisateur
	tmpl := template.Must(template.ParseFiles("./templates/forms.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
	}

	userInput := r.FormValue("entrer")

	if len(userInput) > 0 {

		id := 0

		userInput = strings.ToLower(userInput)

		for _, val := range Donne.Detail {
			if strings.ToLower(val.Name) == userInput {
				id = val.Id
			}
		}

		userInput = ""

		for x := 0; x < len(Donne.Detail); x++ {
			if id == Donne.Detail[x].Id {
				Donne.DetailSimple.Id = Donne.Detail[x].Id
				Donne.DetailSimple.Image = Donne.Detail[x].Image
				Donne.DetailSimple.Name = Donne.Detail[x].Name
				Donne.DetailSimple.Members = Donne.Detail[x].Members
				Donne.DetailSimple.CreationDate = Donne.Detail[x].CreationDate
				Donne.DetailSimple.FirstAlbum = Donne.Detail[x].FirstAlbum
				if Donne.ApiOn == 1 {
					Donne.DetailSimple.Locations = Donne.DetailLocation.Index[x].Locations
					Donne.DetailSimple.Dates = Donne.DetailDate.Index[x].Dates
				}

				x = len(Donne.Detail)

				Donne.Choix = 1
			}
		}

	}
	tmpl.Execute(w, struct{ Success bool }{true})
}

func ImageArtiste(w http.ResponseWriter, r *http.Request, Donne *Donne) { // page de l'entrée utilisateur
	tmpl := template.Must(template.ParseFiles("./templates/test.html"))

	id := 0

	Donne.Choix2 = 0

	Donne.DetailFull = Donne.Detail
	var DetailTest []groupie.Detail

	memberString := r.FormValue("member")
	member, _ := strconv.Atoi(memberString)

	dateString := r.FormValue("date")
	date, _ := strconv.Atoi(dateString)

	albumString := r.FormValue("album")
	album, _ := strconv.Atoi(albumString)

	idStr := r.FormValue("artisteImage")
	id, _ = strconv.Atoi(idStr)

	europe := r.FormValue("europe")

	amerique := r.FormValue("amerique")

	afrique := r.FormValue("afrique")

	asie := r.FormValue("asie")

	oceanie := r.FormValue("oceanie")

	Donne.DetailFull = groupie.Filtre(Donne.Detail, Donne.DetailLocation, Donne.DetailFull, DetailTest, member, date, album, europe, amerique, afrique, asie, oceanie)

	if id > 0 {

		for x := 0; x < len(Donne.Detail); x++ {
			if id == Donne.Detail[x].Id {
				Donne.DetailSimple.Id = Donne.Detail[x].Id
				Donne.DetailSimple.Image = Donne.Detail[x].Image
				Donne.DetailSimple.Name = Donne.Detail[x].Name
				Donne.DetailSimple.Members = Donne.Detail[x].Members
				Donne.DetailSimple.CreationDate = Donne.Detail[x].CreationDate
				Donne.DetailSimple.FirstAlbum = Donne.Detail[x].FirstAlbum
				if Donne.ApiOn == 1 {
					Donne.DetailSimple.Locations = Donne.DetailLocation.Index[x].Locations
					Donne.DetailSimple.Dates = Donne.DetailDate.Index[x].Dates
				}
				Donne.Choix2 = 1

				x = len(Donne.Detail)
			}
		}

		id = 0

	}
	tmpl.Execute(w, struct{ Success bool }{true})
}

func main() { // fonction main
	var Donne *Donne = new(Donne)

	Start(Donne)

	var loc string

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //page d'acceuil
		Donne.Choix = 0
		User(w, r, Donne)
		if Donne.Choix == 1 {
			http.Redirect(w, r, "/artiste", http.StatusSeeOther)
		}
		Home(w, r)
	})

	http.HandleFunc("/AllArtiste", func(w http.ResponseWriter, r *http.Request) { //page d'acceuil
		Donne.Choix2 = 0
		ImageArtiste(w, r, Donne)
		if Donne.Choix2 == 1 {
			http.Redirect(w, r, "/artiste", http.StatusSeeOther)
		}
		AllArtise(w, r, Donne)
	})

	http.HandleFunc("/artiste", func(w http.ResponseWriter, r *http.Request) { //page d'acceu
		Artiste(w, r, Donne)

	})

	http.HandleFunc("/map", func(w http.ResponseWriter, r *http.Request) { //page d'acceu
		loc = groupie.Cord(Donne.DetailSimple)
		Map(w, r, loc)

	})

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8000", nil)
}

func Start(Donne *Donne) {

	Donne.Detail = groupie.AllARtiste()
	Donne.DetailSimple = groupie.Setup()
	Donne.DetailDate = groupie.ApiDate()

	Donne.Choix = 0
	Donne.Choix2 = 0
	Donne.LenDetail = len(Donne.Detail)

	x := 0
	var entree string
	for x < 1 {
		fmt.Print("0 sans api, 1 avec api, 2 api lente : ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		entree = scanner.Text()
		if len(entree) < 1 {
			fmt.Println("ecrit quelque chose")
		} else {
			a, _ := strconv.Atoi(entree)
			if a == 1 {
				Donne.ApiOn = 1
				Donne.DetailLocation = groupie.ApiLocation(Donne.ApiOn)
				x = 1
			} else if a == 2 {
				Donne.ApiOn = 2
				Donne.DetailLocation = groupie.ApiLocation(Donne.ApiOn)
				x = 1
			} else {
				Donne.ApiOn = 0
				x = 1
			}
		}
	}

	fmt.Println("Pret")
}
