package main

import (
	"fmt"
	"github.com/Generalzy/General/General"
	"log"
	"net/http"
)

func main() {
	engine:=General.New()
	engine.Get("/", func(w http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			_,_ = fmt.Fprintf(w, "Header[%q] = %q \n", k, v)
		}
	})
	log.Fatalln(engine.Run(":8080"))
}
