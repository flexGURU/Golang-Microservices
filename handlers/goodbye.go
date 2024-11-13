package handlers

import (
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye  {

	return &Goodbye{l}
	
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	data, _ := io.ReadAll(r.Body)

	log.Println(string(data))


	w.Write([]byte("Byee"))


	
}