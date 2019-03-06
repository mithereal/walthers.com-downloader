/*
This code Do the following stuff :
1) login to a website called : website.com by submitting password and username on the page with url :- http://website.com/login
2) Now after login using the cookies stored by this webiste access user profile page
3) Now using same client which stored the required cookies make another post request to user profile page present at page :-
http://website.com/upser_profile_page .
4) Now get html of this whole page and print it in log as a string .

*/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"

	"golang.org/x/net/publicsuffix"
)

const fp = "xml"

func fetch() {
	fmt.Println("Fetching All Products Please Wait ....")
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{Jar: jar}
	resp, err := client.PostForm("https://dealers.walthers.com/exec/login", url.Values{
		"password": {*password},
		"userid":   {*username},
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err = client.Get("https://dealers.walthers.com/exec/search?manu=&item=&words=restrict&split=30&category=&scale=&instock=Q&price_min=&price_max=&crdate_min=&crdate_max=&exporttype=xml")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	newpath := filepath.Join(".", fp)

	filen := newpath + "/all.xml"

	os.Remove(filen)

	os.MkdirAll(newpath, os.ModePerm)

	// Create the file
	out, err := os.Create(filen)

	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
}
