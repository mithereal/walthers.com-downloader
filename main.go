package main // package clause

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	Version = "1.0"
)

type PART struct {
	XMLName         xml.Name `xml:"PART"`
	PARTNO          string   `xml:"PARTNO"`
	UPC             string   `xml:"UPC"`
	DESCRIPTION     string   `xml:"DESCRIPTION"`
	LONGDESCRIPTION string   `xml:"LONGDESCRIPTION"`
	CATEGORY        string   `xml:"CATEGORY"`
	MSRP            string   `xml:"MSRP"`
	RETAILSALEPRICE string   `xml:"RETAILSALEPRICE"`
	PRICE           string   `xml:"PRICE"`
	RETAILSALESTART string   `xml:"RETAILSALESTART"`
	RETAILSALEEND   string   `xml:"RETAILSALEEND"`
	DLR_NET         string   `xml:"DLR_NET"`
	DLRSALESTART    string   `xml:"DLRSALESTART"`
	DLRSALEEND      string   `xml:"DLRSALEEND"`
	SCALE           string   `xml:"SCALE"`
	INSTOCK         string   `xml:"INSTOCK"`
	EXPECTED        string   `xml:"EXPECTED"`
	AVAILABILITY    string   `xml:"AVAILABILITY"`
	DISCONTINUED    string   `xml:"DISCONTINUED"`
	ADVRES          string   `xml:"ADVRES"`
}

type PRODLIST struct {
	XMLName xml.Name `xml:"PRODLIST"`
	PARTS   []PART   `xml:"PART"`
}

func main() {
	fmt.Println("Starting the Walthers Fetch Tool")
	createDirs()
	kingpin.Parse()
	fetch()
	marshall()
	if *images != false {
		fetchPics()
	}
	fmt.Println("Success: Your files are in the 'xml' and 'images' directories")
}

func createDirs() {
	var fp = "xml"
	var fp2 = "images"
	newpath := filepath.Join(".", fp)
	newpath2 := filepath.Join(".", fp2)
	os.MkdirAll(newpath, os.ModePerm)
	os.MkdirAll(newpath2, os.ModePerm)
}

func marshall() {
	fmt.Println("Finding Sales Please Wait ....")
	//First open your file and create reader. You can also use gzip file check tests
	file, _ := os.Open("xml/all.xml")
	defer file.Close()
	br := bufio.NewReader(file)

	// then create  following channel to read your parsed data from.
	var resultChannel = make(chan XMLEntry)

	// init parser
	var parser = XMLParser{
		R: br,
		// define tag to loop over
		LoopTag:    "PART",
		OutChannel: &resultChannel,
		// you can skip tags that you are not interested it relatively speeds up the process
		SkipTags: []string{"HAZARDOUS", "ORDMIN", "ORDMULT", "BACKORDERBY", "DISCOUNTCODE", "DISCOUNT", "COUNTRY"},
	}

	v := &PRODLIST{}

	// start parsing with a go routine
	go parser.Parse()

	// and finally read parsed data
	for part := range resultChannel {

		PARTNO := part.Elements["PARTNO"][0].InnerText

		UPC := part.Elements["UPC"][0].InnerText

		DESCRIPTION := part.Elements["DESCRIPTION"][0].InnerText

		LONGDESCRIPTION := part.Elements["LONGDESCRIPTION"][0].InnerText

		CATEGORY := part.Elements["CATEGORY"][0].InnerText

		MSRP := part.Elements["MSRP"][0].InnerText

		RETAILSALEPRICE := part.Elements["RETAILSALEPRICE"][0].InnerText

		PRICE := part.Elements["PRICE"][0].InnerText

		RETAILSALESTART := part.Elements["RETAILSALESTART"][0].InnerText

		RETAILSALEEND := part.Elements["RETAILSALEEND"][0].InnerText

		DLR_NET := part.Elements["DLR_NET"][0].InnerText

		DLRSALE := part.Elements["DLRSALE"][0].InnerText

		DLRSALESTART := part.Elements["DLRSALESTART"][0].InnerText

		DLRSALEEND := part.Elements["DLRSALEEND"][0].InnerText

		SCALE := part.Elements["SCALE"][0].InnerText

		INSTOCK := part.Elements["INSTOCK"][0].InnerText

		EXPECTED := part.Elements["EXPECTED"][0].InnerText

		AVAILABILITY := part.Elements["AVAILABILITY"][0].InnerText

		DISCONTINUED := part.Elements["DISCONTINUED"][0].InnerText

		ADVRES := part.Elements["ADVRES"][0].InnerText

		if DLRSALE != "" {

			v.PARTS = append(v.PARTS, PART{PARTNO: PARTNO, UPC: UPC, DESCRIPTION: DESCRIPTION, LONGDESCRIPTION: LONGDESCRIPTION, CATEGORY: CATEGORY, MSRP: MSRP, RETAILSALEPRICE: RETAILSALEPRICE, PRICE: PRICE, RETAILSALESTART: RETAILSALESTART, RETAILSALEEND: RETAILSALEEND, DLR_NET: DLR_NET, DLRSALESTART: DLRSALESTART, DLRSALEEND: DLRSALEEND, SCALE: SCALE, INSTOCK: INSTOCK, EXPECTED: EXPECTED, AVAILABILITY: AVAILABILITY, DISCONTINUED: DISCONTINUED, ADVRES: ADVRES})
		}
	}

	filen := "xml/sales.xml"
	out, _ := os.Create(filen)

	xmlWriter := io.Writer(out)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

}
