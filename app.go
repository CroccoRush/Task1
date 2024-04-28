package main

import (
	"log"
	"net/http"
)

func StartServer() {

	log.Printf("http server started")
	router := NewRouter()
	server := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	log.Println("http server started on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
