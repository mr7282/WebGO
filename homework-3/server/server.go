package server

import (
	"net/http"
	"html/template"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	RootDir string
	TemplateDir string
	IndexTemplate string

}

func NewServer() *Server {
	return &Server{
		RootDir: "homework-3",
		TemplateDir: "www/templates",
		IndexTemplate: "./index.html"
	}
}

func StartServer() *http.ServeMux{
	Route := http.NewServeMux()
	logrus.Fatal(http.ListenAndServe(":8080", Route))
	return Route

}

var tmpl = template.Must(template.New("myBlog").ParseFiles("./www/templates/index.html"))

func viewBlog(wr http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(wr, "Blog", myBlog); err != nil {
		log.Println(err)
	}
}