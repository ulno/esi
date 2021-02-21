// Example based on Rest and unit testing in https://golangdocs.com/

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article ...
type Article struct {
	ID     string `json:"Id"`
	Title  string `json:"Title"`
	Author string `json:"author"`
	Link   string `json:"link"`
}

const endPointHit = "Endpoint Hit:"

// Articles ...
var Articles []Article

// GenHomePage returns the content of the home page
func GenHomePage() []byte {
	return []byte("Welcome to the HomePage!")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println(endPointHit, "home page")
	w.Write(GenHomePage())
}

// GenSingleArticle returns all articles matching the given id
func GenSingleArticle(id string) []byte {
	buf := new(bytes.Buffer)
	for _, article := range Articles {
		if article.ID == id {
			json.NewEncoder(buf).Encode(article)
		}
	}
	return buf.Bytes()
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	log.Println(endPointHit, "return single article")
	vars := mux.Vars(r)
	key := vars["id"]

	w.Write(GenSingleArticle(key))
}

// AddNewArticle add an article to the internal Articles list
func AddNewArticle(article Article) {
	Articles = append(Articles, article)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	log.Println(endPointHit, "create new article")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	AddNewArticle(article)

	json.NewEncoder(w).Encode(article)
}

// DeleteArticle deletes all articles that have the given id from teh internal Articles list
func DeleteArticle(id string) {
	for index, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	log.Println(endPointHit, "delete article")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

// GenAllArticles returns a json list of all articles in the internal article list
func GenAllArticles() []byte {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(Articles)
	return buf.Bytes()
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	log.Println(endPointHit, "return all articles")
	w.Write(GenAllArticles())
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	Articles = []Article{
		{ID: "1",
			Title:  "Python Intermediate and Advanced 101",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089KVK23P"},
		{ID: "2",
			Title:  "R programming Advanced",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089WH12CR"},
		{ID: "3",
			Title:  "R programming Fundamentals",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089S58WWG"},
	}
	handleRequests()
}
