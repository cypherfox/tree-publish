package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/julienschmidt/sse"
)

type HomePage struct {
	Time string
}

func serveHomepage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writingSync.Lock()
	programIsRunning = true
	writingSync.Unlock()
	var homepage HomePage
	homepage.Time = time.Now().String()
	tmpl := template.Must(template.ParseFiles("frontend/html/homepage.html"))
	_ = tmpl.Execute(writer, homepage)
	writingSync.Lock()
	programIsRunning = false
	writingSync.Unlock()
}

func streamTime(timer *sse.Streamer) {
	fmt.Println("Streaming time started")
	for serviceIsRunning {
		timer.SendString("", "time", time.Now().String())
		time.Sleep(1 * time.Millisecond)
	}
}
