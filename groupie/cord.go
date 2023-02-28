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
	Lat float32
	Lng float32
}

func Cord(location []string, date []string) string {

	var DataLoc []Data

	var Data *Data = new(Data)

	var locLat []float32
	var locLng []float32

	var loc string

	for i := 0; i < len(location); i++ {
		split := strings.Split(location[i], "-")
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
		locLat = append(locLat, DataLoc[i].Results[0].Geometry.Lat)
		locLng = append(locLng, DataLoc[i].Results[0].Geometry.Lng)
	}

	for i := 0; i < len(locLat); i++ {
		loc += fmt.Sprintf("%f", locLat[i])
		loc += " "
		loc += fmt.Sprintf("%f", locLng[i])
		loc += " "
		loc += date[i]
		loc += " "
	}

	return loc
}
