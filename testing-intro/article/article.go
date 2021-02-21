package article

import (
	"bytes"
	"encoding/json"
)

// Repository to store articles
type Repository struct {
	articles []*Article
}

// NewArticleRepository returns article repository
func NewArticleRepository(articles []*Article) *Repository {
	return &Repository{
		articles: articles,
	}
}

// Article ...
type Article struct {
	ID     string `json:"Id"`
	Title  string `json:"Title"`
	Author string `json:"author"`
	Link   string `json:"link"`
}

// GenSingleArticle returns all articles matching the given id
func (r *Repository) GenSingleArticle(id string) []byte {
	buf := &bytes.Buffer{}
	for _, article := range r.articles {
		if article.ID == id {
			json.NewEncoder(buf).Encode(article)
		}
	}
	return buf.Bytes()
}

// AddNewArticle add an article to the internal articles list
func (r *Repository) AddNewArticle(article *Article) {
	r.articles = append(r.articles, article)
}

// DeleteArticle deletes all articles that have the given id from teh internal articles list
func (r *Repository) DeleteArticleWithID(id string) {
	for index, article := range r.articles {
		if article.ID == id {
			r.articles = append(r.articles[:index], r.articles[index+1:]...)
		}
	}
}

// genAllArticles returns a json list of all articles in the internal article list
func (r *Repository) GenAllArticles() []byte {
	buf := &bytes.Buffer{}
	json.NewEncoder(buf).Encode(r.articles)
	return buf.Bytes()
}
