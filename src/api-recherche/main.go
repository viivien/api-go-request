package main

import (
	"log"
	"net/http"
)

type Recherche struct {
	question     string
	propositions map[string]int8
}

type Proposition struct {
	Mot       string
	Occurence int8
}

func main() {
	router := InitializeRouter()

	log.Fatal(http.ListenAndServe("", router))
}
