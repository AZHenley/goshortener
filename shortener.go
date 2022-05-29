package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

var urlLength = 3
var rootDomain = "http://azh.lol/"
var links map[string]string

func randomString(length int) string {
	rstr := make([]byte, length)
	chars := []byte("abcdefghijklmnopqrstuvwxyz")
	for i := 0; i < length; i++ {
		rstr[i] = chars[rand.Intn(len(chars))]

	}
	return string(rstr)
}

func newLink(url string) string {
	// URL must start with http.
	if url[:4] != "http" {
		url = "http://" + url
	}

	for true {
		rstr := randomString(urlLength) // Get a random string.
		if links[rstr] == "" {          // If key does not already exist.
			links[rstr] = url
			return rstr
		}
	}
	return "" // Should never happen.
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	nl := newLink(r.FormValue("body"))
	fmt.Fprintf(w, rootDomain+"%s", nl)
}

func main() {
	links = make(map[string]string)
	links["azh"] = "https://austinhenley.com/"

	http.HandleFunc("/shorten/", shortenHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			http.ServeFile(w, r, "index.html") // Show index page.
		} else { // Attempt to serve shortened url.
			if links[r.URL.Path[1:]] != "" { // Does the link exist?
				//fmt.Fprintf(w, "<meta http-equiv=\"Refresh\" content=\"0; url='%s'\" />", links[r.URL.Path[1:]])
				fmt.Fprintf(w, "<html><meta http-equiv=\"Refresh\" content=\"0; url='%s'\" /></html>", links[r.URL.Path[1:]])
			} else {
				fmt.Fprintf(w, "Link not found.")
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
