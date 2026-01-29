package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	// create and open loggers file
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// close loggers file in the end
	defer file.Close()

	// create logger, which we will use
	mylog := log.New(file, "", log.LstdFlags|log.Lshortfile)

	// create server
	s := server.NewServer(mylog)

	// run server
	err = http.ListenAndServe(s.Server.Addr, s.Server.Handler)
	if err != nil {
		mylog.Println(err)
	}
}
