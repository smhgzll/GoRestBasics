package Endpoints

import (
	"GoBasics/Services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllArticles")
	var articles = Services.GetArticles()
	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	id, err := strconv.Atoi(key)
	if err == nil {
		fmt.Println(id)
	}

	var articles = Services.GetArticle(id)
	json.NewEncoder(w).Encode(articles)
}