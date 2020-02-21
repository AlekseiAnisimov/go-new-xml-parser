package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	//"encoding/xml"
)

type Categories struct {
	XMLName xml.Name `xml:"categorties"`
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

	cat.getCategories(data)

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
	err := xml.Unmarshal(data, cat)
	if err != nil {
		return err
	}

	return nil
}
