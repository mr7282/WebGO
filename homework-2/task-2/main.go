package main

import (
	"log"
	"net/http"
	
)

// Напишите два роута: один будет записывать информацию в Cookie (например, имя), а второй — получать ее и выводить в ответе на запрос.


func main() {
	route := http.NewServeMux()
	route.HandleFunc("/", setCookie)
	route.HandleFunc("/cookie", getCookie)
	log.Fatal(http.ListenAndServe(":8080", route))
}

func setCookie(wr http.ResponseWriter, req *http.Request) {
	myCookie := http.Cookie{Name:"myCookie", Value:"information"}
	http.SetCookie(wr, &myCookie)
	wr.Write([]byte("set cookie"))
}

func getCookie(wr http.ResponseWriter, req *http.Request) {
	reqCookie, err := req.Cookie("myCookie")
	if err!= nil {
		log.Fatal(err)
	}
	wr.Write([]byte("Cookie:" + reqCookie.Value))
         
}


