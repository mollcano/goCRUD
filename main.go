package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

type Article struct {
	ID      string `json:"id,omitempty"`
	Title   string `json:"title, omitempty"`
	Desc    string `json:"desc, omitempty"`
	Content string `json:"content, omitempty"`
}

var articles []Article

func GetArticlesEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range articles {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Article{})
}

func GetArticleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(articles)
}

func CreateArticleEndpoint(w http.ResponseWriter, req *http.Request) {
	var article Article
	_ = json.NewDecoder(req.Body).Decode(&article)
	articles = append(articles, article)
	json.NewEncoder(w).Encode(articles)
}

func DeleteArticleEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(articles)
}

func UpdateArticleEndpoint(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	var article Article
	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			_ = json.NewDecoder(req.Body).Decode(&article)
			articles = append(articles, article)
			break
		}
	}
	json.NewEncoder(w).Encode(articles)
}

func main() {
	router := mux.NewRouter()
	articles = append(articles, Article{ID: "1", Title: "Hello", Desc: "Description", Content: "Content"})
	articles = append(articles, Article{ID: "2", Title: "Hello2", Desc: "Description", Content: "Content"})
	router.HandleFunc("/articles", GetArticleEndpoint).Methods("GET")
	router.HandleFunc("/articles/{id}", GetArticlesEndpoint).Methods("GET")
	router.HandleFunc("/articles", CreateArticleEndpoint).Methods("POST")
	router.HandleFunc("/articles/{id}", DeleteArticleEndpoint).Methods("DELETE")
	router.HandleFunc("/articles/{id}", UpdateArticleEndpoint).Methods("PUT")
	log.Fatal(http.ListenAndServe(":12345", router))
}
