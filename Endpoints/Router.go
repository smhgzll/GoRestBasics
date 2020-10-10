package Endpoints

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/articles/all", returnAllArticles)
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}