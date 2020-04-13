package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	// "os"
	"log"
)

// Напишите функцию, которая будет получать на вход строку с поисковым запросом (string) и массив ссылок на страницы, по которым стоит
// произвести поиск ([]string). Результатом работы функции должен быть массив строк со ссылками на страницы, на которых обнаружен 
// поисковый запрос. Функция должна искать точное соответствие фразе в тексте ответа от сервера по каждой из ссылок. 

func testSearch(searchQuery string, whereSearch []string) {
	for	_, url := range whereSearch {
		openURL(url)
	}
}

func openURL(url string) []byte {
	getUrl, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer getUrl.Body.Close()

	return getUrl.Body


}


func main() {

	resp, err := http.Get("https://golang.org/")
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	
}
