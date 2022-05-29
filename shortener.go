package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

var urlLength = 3
var maxURLs = 30
var currentURLs = 0
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
	// Delete URLs if too many.
	currentURLs += 1
	if currentURLs > maxURLs {
		links = make(map[string]string)
		currentURLs = 1
	}
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

	http.HandleFunc("/shorten/", shortenHandler)                            // API for creating new links.
	http.HandleFunc("/img/", func(w http.ResponseWriter, r *http.Request) { // Serve images.
		buf, err := ioutil.ReadFile(r.URL.Path[1:])
		if err != nil {
			fmt.Fprintf(w, "Image not found.")
		} else {
			w.Header().Set("Content-Type", "image/png")
			w.Write(buf)
		}
	})
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
