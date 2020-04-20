package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Напишите функцию, которая будет получать на вход строку с поисковым запросом (string) и массив ссылок на страницы, по которым стоит
// произвести поиск ([]string). Результатом работы функции должен быть массив строк со ссылками на страницы, на которых обнаружен
// поисковый запрос. Функция должна искать точное соответствие фразе в тексте ответа от сервера по каждой из ссылок.

// Данная функция принимает в качестве аргументтов сам поисковый запрос с типом string и url интернет страниц, на которых необходимо искать с типом []string
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

func main() {
	fmt.Println(testSearch("implements some I/O utility functions.", []string{"https://ru.stackoverflow.com/", "https://golang.org/pkg/io/ioutil/", "https://proglib.io/p/learn-regex/"}))
}
