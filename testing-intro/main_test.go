// Example based on Rest and unit testing in https://golangdocs.com/

package main

import (
	"encoding/json"
	"github.com/ulno/esi/testing-intro/article"
	"testing"
)

func TestAddNewArticle(t *testing.T) {
	testArticle := article.Article{
		ID:     "42",
		Title:  "ulno.net",
		Author: "Ulno",
		Link:   "http://ulno.net",
	}
	articleRepository.AddNewArticle(&testArticle)
	findArticleJSON := articleRepository.GenSingleArticle("42")
	findArticle := article.Article{}
	json.Unmarshal(findArticleJSON, &findArticle)
	if findArticle != testArticle {
		t.Error("Couldn't find article after adding.")
		return
	}
}
