package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/go-chi/chi"

	"github.com/Sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()
	lg := logrus.New()
	// router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("www/static"))))

	serv := Server{
		lg:    lg,
		Title: "ToDo list",
		Tasks: TaskItems{
			{Text: "Дедлайн завтра, Горим!", Completed: false, Labels: []string{"срочно"}},
			{Text: "Проснулся", Completed: true},
			{Text: "Поел", Completed: true},
			{Text: "Поработал", Completed: false},
			{Text: "Уснул", Completed: false},
		},
	}

	r.Route("/", func(r chi.Router) {
		r.Get("/", serv.HandleGetIndex)
		r.Post("/{taskID}/{status}", serv.HandlePostTaskStatus)
	})

	lg.Info("server is running")
	http.ListenAndServe(":8080", r)
}

type Server struct {
	lg    *logrus.Logger
	Title string
	Tasks TaskItems
}

type TaskItems []TaskItem
type TaskItem struct {
	Text      string // проснулся, встал, поел, поработал, поспал...
	Completed bool
	Labels    []string
}

func (s *Server) HandleGetIndex(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./www/static/index.html")
	data, _ := ioutil.ReadAll(file)

	templ := template.Must(template.New("page").Parse(string(data)))
	err := templ.ExecuteTemplate(w, "page", s)
	if err != nil {
		s.lg.WithError(err).Error("template")
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (s *Server) HandlePostTaskStatus(w http.ResponseWriter, r *http.Request) {
	taskIDStr := chi.URLParam(r, "taskID")
	taskStatusStr := chi.URLParam(r, "status")

	taskID, _ := strconv.ParseInt(taskIDStr, 10, 64)
	taskStatus, _ := strconv.ParseBool(taskStatusStr)

	s.Tasks[taskID].Completed = taskStatus

	data, _ := json.Marshal(s.Tasks[taskID])
	w.Write(data)
	s.lg.WithField("tasks", s.Tasks).Info("status changed")
}

func (tasks TaskItems) TasksWithStatus(completed bool) int {
	total := 0
	for _, tasks := range tasks {
		if tasks.Completed == completed {
			total++
		}
	}
	return total
}

func (tasks TaskItems) CompletePercent() float64 {
	prc := float64(tasks.TasksWithStatus(true)) / float64(len(tasks)) * 100
	return prc
}
