package main

import(
	"net/http"
	"log"
)

var tmpl = template.Must

func main() {
	route := http.NewServeMux()
	route.HandleFunc("/", viewBlog)
	log.Fatal(http.ListenAndServe(":8080", route))
}

func viewBlog(wr http.ResponseWriter, r *http.Request) {
	if err := 
}