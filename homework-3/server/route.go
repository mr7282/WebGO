package server

import (
	"net/http"
)

// Router - вызывает обработчики в зависимости от поступившего запроса
func (serv *Server) Router(route *http.ServeMux) {
	route.Handle("/favicon.ico", http.FileServer(http.Dir("./www")))
	route.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	route.HandleFunc("/index", serv.indexTpl)
	route.HandleFunc("/find/post", serv.findPost)
	route.HandleFunc("/editView", serv.editView)
	route.HandleFunc("/editPost", serv.editPost)
}