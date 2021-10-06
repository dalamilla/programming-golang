package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
)

type shorter struct {
	PATH string `yaml:"path"`
	URL  string `yaml:"url"`
}

var pathsToUrl = make(map[string]string)

func main() {
	data := []shorter{}

	log.Println("Open the yaml file")
	file, err := ioutil.ReadFile("urls.yaml")
	if err != nil {
		log.Fatal("Failed to open the yaml file")
	}

	log.Println("Unmarshall the yaml file")
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Println("Convert yaml data to map")
	for _, d := range data {
		pathsToUrl[d.PATH] = d.URL
	}

	http.HandleFunc("/", shortenerHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	log.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func shortenerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("receive request with path:", r.URL.Path)
	path := r.URL.Path

	if url, ok := pathsToUrl[path]; ok {
		log.Println("the path exist on the yaml, redirect to existing url")
		http.Redirect(w, r, url, http.StatusFound)
	}

	log.Println("the path does not exist on the yaml")
	fmt.Fprintf(w, "Hello, Good Friend")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {}
