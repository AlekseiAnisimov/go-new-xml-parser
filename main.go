package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

type Categories struct {
	XMLName  xml.Name   `xml:"categories"`
	Category []Category `xml:"category"`
}

type Category struct {
	Id    string `xml:"id,attr"`
	Value string `xml:",chardata"`
}

func main() {
	var xmlFile string = "alidump.yml"
	var data []byte
	var cat Categories
	data, err := readFile(xmlFile)

	if err != nil {
		fmt.Print(err)
	}

	err = cat.getCategories(data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cat)
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
