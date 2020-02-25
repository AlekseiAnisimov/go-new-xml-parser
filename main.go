package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

type Categories struct {
	XMLName  xml.Name   `xml:"categories"`
	Category []Category `xml:"category"`
}

type Category struct {
	Id    string `xml:"id,attr"`
	Value string `xml:",chardata"`
}

type Offers struct {
	XMLName xml.Name `xml:"offers"`
	Offer   []Offer  `xml:"offer"`
}

type Offer struct {
	Id          string  `xml:"id, attr"`
	Available   string  `xml:"available, attr"`
	CategoryId  int     `xml:"categoryId"`
	Category    string  `xml:"category"`
	Name        string  `xml:"name"`
	Description string  `xml:"description"`
	Picture     string  `xml:"picture"`
	Price       float32 `xml:"price"`
	CurrencyId  string  `xml:"currencyId"`
	Url         string  `xml:"url"`
}

type dbDevelopment struct {
	Development struct {
		Dialect    string
		Datasource string
	}
}

func main() {
	var xmlFile string = "alidump.yml"
	var yamlDbFile string = "dbconfig.yml"
	var data []byte
	var cat Categories
	var off Offers
	var dbd dbDevelopment

	data, err := readFile(xmlFile)

	if err != nil {
		fmt.Print(err)
	}

	err = cat.getCategories(data)
	if err != nil {
		fmt.Println(err)
	}

	err = off.getOffers(data)
	if err != nil {
		fmt.Println(err)
	}

	dbParamsFile, err := ioutil.ReadFile(yamlDbFile)
	if err != nil {
		panic(err)
	}

	_ = yaml.Unmarshal(dbParamsFile, &dbd)

	db, _ := dbx.Open(dbd.Development.Dialect, dbd.Development.Datasource)
	fmt.Println(db)
}

func readFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (cat *Categories) getCategories(data []byte) error {
	decoder := xml.NewDecoder(strings.NewReader(string(data)))
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "categories" {
				err := decoder.DecodeElement(&cat, &se)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}

	return nil
}

func (off *Offers) getOffers(data []byte) error {
	decoder := xml.NewDecoder(strings.NewReader(string(data)))
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "offers" {
				err := decoder.DecodeElement(&off, &se)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}

	return nil
}
