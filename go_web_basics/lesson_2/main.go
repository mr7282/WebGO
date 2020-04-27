package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	uuid "github.com/satori/go.uuid"

	"github.com/Sirupsen/logrus"
	"github.com/go-chi/chi"
)

const (
	ARG_KEY    = "key"
	ARG_VALUE  = "value"
	COOKIE_KEY = "cookie"
)

type DB struct {
	data  map[string]string
	mutex *sync.Mutex
}

func (s *DB) Set(key, value string) { // s.Set("msg", "test123")
	s.mutex.Lock()
	s.data[key] = value
	s.mutex.Unlock()
}

func (s *DB) Get(key string) (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	value, exist := s.data[key]
	if !exist {
		return "", errors.New("not exists")
	}
	return value, nil
}

func main() {
	stopchan := make(chan os.Signal)

	inmemoryDB := make(map[string]string)
	protectedDB := &DB{data: inmemoryDB, mutex: &sync.Mutex{}}

	logrus.SetReportCaller(true)

	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		r.Get("/{key}", protectedDB.GetIndexHandler)
		r.Post("/{key}", protectedDB.PostIndexHandler)
	})

	go func() {
		err := http.ListenAndServe(":8080", router)
		log.Fatal(err)
	}()

	signal.Notify(stopchan, os.Interrupt, os.Kill)
	<-stopchan
	log.Print("gracefull shutdown")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := ioutil.ReadFile("./pages/index.html")
	w.Write(file)
}

func (s *DB) GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	userKey := CookieControl(w, r)
	logrus.Info(userKey)

	key := chi.URLParam(r, ARG_KEY) // get localhost:8080/12312312
	value, err := s.Get(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Write([]byte(value))
	}
	logrus.Info(value)
}

func (s *DB) PostIndexHandler(w http.ResponseWriter, r *http.Request) {
	userKey := CookieControl(w, r)
	logrus.Info(userKey)

	key := chi.URLParam(r, ARG_KEY) // post localhost:8080/12312312?value=90
	value := r.FormValue(ARG_VALUE)
	s.data[key] = value

	logVal := map[string]interface{}{
		ARG_KEY:   key,
		ARG_VALUE: value,
	}
	w.Write([]byte("value stored " + key + " " + value))
	logrus.WithFields(logVal).Info("value stored")
}

func CookieControl(w http.ResponseWriter, r *http.Request) string {
	cookie, _ := r.Cookie(COOKIE_KEY)
	if cookie == nil {
		cookie = &http.Cookie{
			Name: COOKIE_KEY,
		}
	}

	userKey := cookie.Value
	if userKey != "" {
		return userKey
	} else {
		cookie.Value = uuid.Must(uuid.NewV4()).String()
		http.SetCookie(w, cookie)
		return cookie.Value
	}

}
