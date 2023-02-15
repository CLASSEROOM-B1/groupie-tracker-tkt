package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Data struct {
	Results []Results
}

type Results struct {
	Geometry Geometry
}

type Geometry struct {
	Lat float64
	Lng float64
}

func Cord(location []string) string {

	var DataLoc []Data

	var Data *Data = new(Data)

	var locLat []float64
	var locLng []float64

	var loc string

	fmt.Println(location)

	for i := 0; i < len(location); i++ {
		split := strings.Split(location[i], "-")
		fmt.Println(split[0])
		var test string = "https://api.opencagedata.com/geocode/v1/json?q=" + split[0] + "&key=735c71b1c2694ef383ee59af8349e696"
		url, err := http.Get(test)
		if err != nil {
			os.Exit(1)
		}
		Arstiste_json, err := ioutil.ReadAll(url.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(Arstiste_json, Data)
		DataLoc = append(DataLoc, *Data)
	}
	// fmt.Println(DataLoc[4].Results[0].Geometry.Lat)
	// fmt.Println(DataLoc[4].Results[0].Geometry.Lng)

	for i := 0; i < len(DataLoc); i++ {
		locLat = append(locLat, DataLoc[0].Results[i].Geometry.Lat)
		locLng = append(locLng, DataLoc[0].Results[i].Geometry.Lng)
	}

	for i := 0; i < len(locLat); i++ {
		loc += fmt.Sprintf("%f", locLat[i])
		loc += " "
		loc += fmt.Sprintf("%f", locLng[i])
		loc += " "
	}

	return loc
}
