package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

type Categories struct {
	XMLName xml.Name `xml:"categories"`
	category []Category
}

type Category struct {
	XMLName xml.Name `xml:"category"`
	Id int8 `xml:"id,attr"`
	Desc string `xml:"category"`
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
	data, err := ioutil.ReadFile(filename);
	if err != nil {
		return nil, err
	}

	return data, err
}

func (cat *Categories) getCategories(data []byte) error {
	err := xml.NewDecoder(strings.NewReader(string(data))).Decode(&cat)
	if err != nil {
		return err
	}

	return nil
}
