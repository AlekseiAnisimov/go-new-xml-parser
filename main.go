package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
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
	Id    int    `xml:"id,attr"`
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
	//fmt.Println(db)
	cat.saveCategories(db)
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

func (cat *Categories) saveCategories(db *dbx.DB) {
	var batchInsert strings.Builder

	if len(cat.Category) == 0 {
		panic("Category struct is empty")
	}

	for _, value := range cat.Category {
		batchInsert.WriteString(fmt.Sprintf("(%d,%s),", value.Id, strconv.Quote(value.Value)))
	}

	_, err := db.NewQuery("INSERT INTO categories_bck SELECT * FROM categories").Execute()
	if err != nil {
		panic(err)
	}

	_, _ = db.TruncateTable("categories").Execute()

	batchInsertFmt := batchInsert.String()
	batchInsertFmt = batchInsertFmt[:len(batchInsertFmt)-1]

	res, err := db.NewQuery("INSERT INTO categories(id, description) VALUE" + batchInsertFmt).Execute()
	if err != nil {
		_, _ = db.NewQuery("INSERT INTO categories SELECT * FROM categories_bck").Execute()
	}
	_, _ = db.TruncateTable("categories_bck").Execute()
	fmt.Println(res)
}
