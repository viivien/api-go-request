package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Status struct {
	status string
}

type Recherche struct {
	question     string
	propositions []string
}

type Proposition struct {
	mot          string
	nbOccurences int8
}

type logWriter struct{}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	statusOK := Status{
		status: "OK",
	}

	json.NewEncoder(w).Encode(statusOK)
}

func RechercheOccurences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Transforme la question
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var recherche Recherche

	err = json.Unmarshal(body, &recherche)

	if err != nil {
		log.Fatal(err)
	}

	// appelle la page google
	resp, err := http.Get("http://google.com" + recherche.question)

	if err != nil {
		fmt.Println("Error:", err)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)

	// parse les reponses possibles

	bodyRecherche, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	// renvoie la reponse
	json.NewEncoder(w).Encode(recherche)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
