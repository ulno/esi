// Example based on Rest and unit testing in https://golangdocs.com/

package main

import (
	"encoding/json"
	"github.com/ulno/esi/testing-intro/article"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var articles = []*article.Article{
	{
		ID:     "1",
		Title:  "Python Intermediate and Advanced 101",
		Author: "Arkaprabha Majumdar",
		Link:   "https://www.amazon.com/dp/B089KVK23P",
	},
	{
		ID:     "2",
		Title:  "R programming Advanced",
		Author: "Arkaprabha Majumdar",
		Link:   "https://www.amazon.com/dp/B089WH12CR",
	},
	{
		ID:     "3",
		Title:  "R programming Fundamentals",
		Author: "Arkaprabha Majumdar",
		Link:   "https://www.amazon.com/dp/B089S58WWG",
	},
}

var articleRepository = article.NewArticleRepository(articles)

const endPointHit = "Endpoint Hit:"

// genHomePage returns the content of the home page
func genHomePage() []byte {
	return []byte("Welcome to the HomePage!")
}

func homePage(w http.ResponseWriter, _ *http.Request) {
	log.Println(endPointHit, "home page")
	w.Write(genHomePage())
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	log.Println(endPointHit, "return single article")
	vars := mux.Vars(r)
	key := vars["id"]

	w.Write(articleRepository.GenSingleArticle(key))
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	log.Println(endPointHit, "create new article")
	reqBody, _ := ioutil.ReadAll(r.Body)
	article := &article.Article{}
	json.Unmarshal(reqBody, article)
	articleRepository.AddNewArticle(article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	log.Println(endPointHit, "delete article")
	vars := mux.Vars(r)
	id := vars["id"]

	articleRepository.DeleteArticleWithID(id)
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	log.Println(endPointHit, "return all articles")
	w.Write(articleRepository.GenAllArticles())
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage).Methods(http.MethodGet)
	myRouter.HandleFunc("/articles", returnAllArticles).Methods(http.MethodGet)
	myRouter.HandleFunc("/article", createNewArticle).Methods(http.MethodPost)
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods(http.MethodDelete)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
