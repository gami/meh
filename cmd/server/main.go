package main

import (
	"fmt"
	"log"
	"net/http"

	openapi "meh/api/openapi"
	"meh/controller"
	"meh/di"
)

func main() {
	port := 8080
	s := Server()
	s.Addr = fmt.Sprintf("0.0.0.0:%d", port)
	log.Printf("running meh addr=%v", s.Addr)
	log.Fatal(s.ListenAndServe())
}

func Server() *http.Server {
	r, err := controller.NewRouter()
	if err != nil {
		log.Fatal(err)
	}

	return &http.Server{
		Handler: openapi.HandlerFromMuxWithBaseURL(
			di.InjectController(),
			r,
			"",
		),
	}
}
