package main

import (
	"myBlog/server"
	"os"
	"os/signal"
	"database/sql"
	"github.com/sirupsen/logrus"
	// "fmt"
	// "time"

	_ "github.com/go-sql-driver/mysql"
)

// NewLogger - Создаёт новый логгер
func NewLogger() *logrus.Logger {
	lg := logrus.New()
	lg.SetReportCaller(false)
	lg.SetFormatter(&logrus.TextFormatter{})
	lg.SetLevel(logrus.DebugLevel)
	return lg
}


func main() {
	lg:= NewLogger()

	db,err := sql.Open("mysql","root:1240608@tcp(88.210.21.76:3360)/myBlog?parseTime=true")
	if err != nil {
		lg.WithError(err).Fatal("can`t connect to DataBase")
	}
	defer db.Close()

	serv := server.NewServer(db, lg)
	go func ()  {
		serv.StartServer()
	}()

	stopSig := make(chan os.Signal)
	signal.Notify(stopSig, os.Interrupt, os.Kill)
	<-stopSig
}


