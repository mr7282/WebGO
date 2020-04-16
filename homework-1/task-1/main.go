package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	// "os"
	"log"
	"bytes"
)

// Напишите функцию, которая будет получать на вход строку с поисковым запросом (string) и массив ссылок на страницы, по которым стоит
// произвести поиск ([]string). Результатом работы функции должен быть массив строк со ссылками на страницы, на которых обнаружен 
// поисковый запрос. Функция должна искать точное соответствие фразе в тексте ответа от сервера по каждой из ссылок. 

func testSearch(searchQuery string, whereSearch []string) []string{
	var includes []string
	lenQuery := len(searchQuery)
	// queryByte := []byte(searchQuery)
	for	_, url := range whereSearch {
		arrByte := openURL(url)
		for i := 0; i >= len(arrByte) ; i++ {
			if bytes.Equal(arrByte[i:lenQuery], []byte(searchQuery)) {
				includes = append(includes, url)
				break
			}
		}
	}
	return includes
}

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

func main() {

		
	fmt.Println(testSearch("Golang сравнить", []string{"https://ru.stackoverflow.com/questions/853566/golang-сравнить-слайсы-байт-полученные-из-16-ричных-строк", "https://golang.org/pkg/io/ioutil/"}))
}
