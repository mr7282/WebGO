package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"log"
)

func route(){
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/ws", WebSocketHandler)
	http.ListenAndServe(":8080", nil)
}

func main() {
	go route()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	<-interrupt	
	log.Println("Shutting down")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	indexHTML, _ := os.Open("websocket.html")
	defer indexHTML.Close()
	indexData, _ := ioutil.ReadAll(indexHTML)
	fmt.Fprint(w, string(indexData))
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := websocket.Upgrade(w, r, nil, 1024, 1024)
	defer conn.Close()
	mType, bts, _ := conn.ReadMessage()
	fmt.Println(string(bts))
	conn.WriteMessage(mType, []byte("message from back-end"))
}
