package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/feed", feed)
	http.ListenAndServe(":4000", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	responseBody := //JSON to a
		"<h1>Neighbour.ly</h1><a href='/feed'>Click for feed</a>"
	fmt.Fprint(w, responseBody)
}

func feed(w http.ResponseWriter, r *http.Request) {
	responseBody := "<h1>Neighbourhood feed</h1><ul><li>Post 1</li><li>Post 2</li></ul><a href='/'>Back we go</a>"
	fmt.Fprint(w, responseBody)
}
