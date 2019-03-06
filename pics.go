package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func fetchPics() {
	fmt.Println("Fetching Pics Please Wait ....")
	file, _ := os.Open("xml/sales.xml")
	defer file.Close()
	br := bufio.NewReader(file)

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

	// start parsing with a go routine
	go parser.Parse()

	// and finally read parsed data
	for part := range resultChannel {

		PARTNO := part.Elements["PARTNO"][0].InnerText

		FetchPic(PARTNO)
	}

}

func FetchPic(mpn string) {
	// Request the HTML page.
	var url = "http://www.walthers.com/exec/productinfo/" + mpn
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("#mainImage").Each(func(i int, element *goquery.Selection) {

		imgSrc, exists := element.Attr("src")
		if exists {
			fetchImage(imgSrc, mpn)
		}
	})
}

func fetchImage(url string, mpn string) {
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create("images/" + mpn + ".jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}
