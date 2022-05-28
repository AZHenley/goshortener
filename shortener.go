package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

var links map[string]string

func randomString(length int) string {
	rstr := make([]byte, length)
	chars := []byte("abcdefghijklmnopqrstuvwxyz")
	for i := 0; i < length; i++ {
		rstr[i] = chars[rand.Intn(len(chars))]

	}
	return string(rstr)
}

func newLink(url string) {
	for true {
		rstr := randomString(4) // Get a random string.
		if links[rstr] == "" {  // If key does not already exist.
			links[rstr] = url
			break
		}
	}
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	links = make(map[string]string)
	links["aaaa"] = "https://austinhenley.com/"

	http.HandleFunc("/shorten/", shortenHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			http.ServeFile(w, r, "index.html") // Show index page.
		} else { // Attempt to serve shortened url.
			if links[r.URL.Path[1:]] != "" { // Does the link exist?
				//fmt.Fprintf(w, "<meta http-equiv=\"Refresh\" content=\"0; url='%s'\" />", links[r.URL.Path[1:]])
				fmt.Fprintf(w, "<html><meta http-equiv=\"Refresh\" content=\"0; url='%s'\" /></html>", links[r.URL.Path[1:]])
			} else {
				fmt.Fprintf(w, "Hmmm.")
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
