package main

import (
	"log"
	"net/http"
	
)

// Напишите два роута: один будет записывать информацию в Cookie (например, имя), а второй — получать ее и выводить в ответе на запрос.

// https://docs.google.com/document/d/1NdWh7wMM0hrgZbarOwaQsAskCN7AiygE4pRl9WzyD_k/edit#

func main() {
	route := http.NewServeMux()
	route.HandleFunc("/", setCookie)
	route.HandleFunc("/cookie", getCookie)
	log.Fatal(http.ListenAndServe(":8080", route))
}

func setCookie(wr http.ResponseWriter, req *http.Request) {
	req.AddCookie(&http.Cookie{Name: "set my coockie"})
	wr.Write([]byte("set cookie"))
}

func getCookie(wr http.ResponseWriter, req *http.Request) {
	myCookie, err := req.Cookie("Name")
	if err != nil {
		log.Fatal(err)
	}
	wr.Write([]byte(myCookie.Name))
}


