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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	links = make(map[string]string)

	http.HandleFunc("/shorten/", handler)
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
