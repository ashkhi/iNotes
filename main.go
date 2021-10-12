package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "iNotes", log.LstdFlags)
	getNotesForTodayHandler := handlers.NewNote(logger)

	serveMultiplexer := http.NewServeMux()
	serveMultiplexer.Handle("/getNotesForToday", getNotesForTodayHandler)

	http.ListenAndServe("localhost:31558", serveMultiplexer)
}
