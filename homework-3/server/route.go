package server

import (
	"net/http"
)

func Router(route *http.ServeMux) {
	route.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	route.HandleFunc("/", viewBlog)
}

