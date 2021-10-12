package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Note struct {
	logger *log.Logger
}

func NewNote(logger *log.Logger) *Note {
	return &Note(logger)
}

func (note *Note) ServerHTTP(respWri http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(respWri, "Something went wrong !", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(respWri, "Request body has data : %s", data)
}
