package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello  {

	return &Hello {l}
	
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    h.l.Println("hello world")

    data, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "ooops", http.StatusBadRequest)
        return
    }

    log.Printf("data received from body is %s: \n", string(data))

    fmt.Fprintf(w, "hello %s from go server ", string(data))
}
