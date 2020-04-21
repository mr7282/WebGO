package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// SearchQuery - struct for discript query
type SearchQuery struct {
	Search string   `json:"search"`
	Sites  []string `json:"sites"`
}

func main() {
	mySearchQuery := SearchQuery{"implements some I/O utility functions.", []string{"https://ru.stackoverflow.com/", "https://golang.org/pkg/io/ioutil/", "https://proglib.io/p/learn-regex/"}}
	byteMSQ, err := json.Marshal(mySearchQuery)
	if err != nil {
		log.Fatal(err)
	}
	readMSQ := bytes.NewReader(byteMSQ)
	resp, err := http.Post("http://localhost:8080/", "application/json", readMSQ)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respRead, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(respRead, &mySearchQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Поисковый запрос: \"%v\" \nНайден на следующих сайтах: %v", mySearchQuery.Search, mySearchQuery.Sites)

}
