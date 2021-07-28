package xmlencoder

import (
	"encoding/xml"
	"fmt"
)

type Gallery struct {
	XMLName   xml.Name   `xml:"gallery"`
	Id        int        `xml:"id,attr"`
	Name      string     `xml:"name"`
	Paintings []Painting `xml:"paintings"`
}

type Painting struct {
	XMLName     xml.Name `xml:"painting"`
	Id          int      `xml:"id,attr"`
	Name        string   `xml:"name"`
	Description []string `xml:"description"`
}

func Export() {
	gallery := &Gallery{Id: 27, Name: "Test Gallery"}
	gallery.Paintings = append(gallery.Paintings, Painting{Id: 0, Name: "Test Painting"})

	out, _ := xml.MarshalIndent(gallery, " ", "  ")

	var p Gallery
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)
}
