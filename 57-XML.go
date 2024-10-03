package main

import (
	"encoding/xml"
	"fmt"
)

/*
---> XML <---
Go offers built-in support for XML and
XML-like formats with the "encoding/xml" package.
*/

/*
Plant will be mapped to XML.

Similarly to the JSON examples, field tags contain directives
for the encoder and decoder.

Here we use some special features of the XML package:
the XMLName field name dictates the name of the
XML element representing
this struct; id,attr means that the Id field is
an XML attribute rather than a nested element.
*/
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id, attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v,  name=%v,  origin=%v", p.Id, p.Name, p.Origin)
}

func learn_XML() {
	fmt.Println("\n------------ learn_XML --------------")
	coffee := &Plant{Id: 17, Name: "coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", " ")
	fmt.Println(out)
	fmt.Println(string(out))

	fmt.Println(xml.Header, string(out))
}

func main() {
	learn_XML()
	println("\n-------------------------------")
}
