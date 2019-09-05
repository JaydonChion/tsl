package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Eatbookdata struct {
	API_Status     string `json:"API_Status"`
	AdminRemarks   string `json:"AdminRemarks"`
	Category       string `json:"Category"`
	ContactEmail   string `json:"ContactEmail"`
	ContactInfo    string `json:"ContactInfo"`
	ContactNumber  string `json:"ContactNumber"`
	ContactStatus  string `json:"ContactStatus"`
	EatBookArticle string `json:"EatBookArticle"`

	Distance       float32 `json:"Distance"`
	EntryAddress   string  `json:"EntryAddress"`
	EntryID        int     `json:"EntryID"`
	EntryName      string  `json:"EntryName"`
	FacebookPage   string  `json:"FacebookPage"`
	Latitude       float32 `json:"Latitude"`
	Longitude      float32 `json:"Longitude"`
	MasterCategory string  `json:"MasterCategory"`
	OpeningHours   string  `json:"OpeningHours"`
	PostalCode     string  `json:"PostalCode"`
	SubFilter      string  `json:"SubFilter"`
}

// func getEatbookData(longitude float32,latitude float32, distance float32)(bookdatas []Eatbookdata, err error){
func getData(dataType string, latitude float64, longitude float64, distance int) (books []Eatbookdata, err error) {
	sdataType := url.QueryEscape(dataType)

	url := fmt.Sprintf("https://appsbytsl.com/API/V1/Nearby/%s/t5LD3v/%f/%f/%d", sdataType, latitude, longitude, distance)

	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(responseData, &books)
	return

}
