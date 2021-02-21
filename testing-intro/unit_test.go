// Example based on Rest and unit testing in https://golangdocs.com/

package main

import (
	"encoding/json"
	"testing"
)

func TestAddNewArticle(t *testing.T) {
	testArticle := Article{
		ID:     "42",
		Title:  "ulno.net",
		Author: "Ulno",
		Link:   "http://ulno.net",
	}
	AddNewArticle(testArticle)
	findArticleJSON := GenSingleArticle("42")
	findArticle := Article{}
	json.Unmarshal(findArticleJSON, &findArticle)
	if findArticle != testArticle {
		t.Error("Couldn't find article after adding.")
		return
	}
}
