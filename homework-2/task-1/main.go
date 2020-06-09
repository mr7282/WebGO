package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

// Используя функцию для поиска из прошлого практического задания, постройте сервер, который будет принимать JSON с поисковым запросом в POST-запросе и возвращать ответ в виде массива строк в JSON.

func main() {
	route := http.NewServeMux()
	route.HandleFunc("/", postHendler)
	log.Fatal(http.ListenAndServe(":8080", route))
}

// searchResponse - struct for response description
type searchResponse struct {
	Search string   `json:"search"`
	Sites  []string `json:"sites"`
}

func postHendler(wr http.ResponseWriter, req *http.Request) {
	mySR := searchResponse{}
	readRequest, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(readRequest, &mySR)
	if err != nil {
		log.Fatal(err)
	}

	searchSite := testSearch(mySR.Search, mySR.Sites)
	for _, appendSite := range searchSite {
		var sites []string
		sites = append(sites, appendSite)
		mySR.Sites = sites
	}

	responseSearch, err := json.Marshal(mySR)
	if err != nil {
		log.Fatal(err)
	}

	wr.Write(responseSearch)
}

func testSearch(searchQuery string, whereSearch []string) []string {
	// Массив, который содержит url интеренет страниц, на которых обноружена строка поискового запроса
	var includes []string
	// количество символов в поисковом запросе
	// lenQuery := len(searchQuery)
	for _, url := range whereSearch {
		// Интернет страница из числа переданных как аргумент в testSearch, в виде []byte
		arrByte := openURL(url)
		if strings.Contains(string(arrByte), searchQuery) {
			includes = append(includes, url)
			break
		}

	}
	return includes
}

// Принимает в качестве аргумента url страницы, а возвращает эту страницу в виде []byte
func openURL(url string) []byte {

	getURL, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer getURL.Body.Close()

	openURL, err := ioutil.ReadAll(getURL.Body)
	if err != nil {
		log.Println(err)
	}
	return openURL
}
