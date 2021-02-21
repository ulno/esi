// Example based on Rest and unit testing in https://golangdocs.com/

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAddExtNewArticle(t *testing.T) {
	testArticle := Article{ID: "42",
		Title:  "ulno.net",
		Author: "Ulno",
		Link:   "http://ulno.net"}
	articleJSON, _ := json.Marshal(testArticle)
	resp, err := http.Post("http://localhost:8000/article", "", bytes.NewBuffer(articleJSON))
	if err != nil {
		t.Error("Problem adding new article via REST:", err)
		return
	}
	resp, err = http.Get("http://localhost:8000/article/42")
	if err != nil {
		t.Error("Problem reading article via REST.")
		return
	}
	findArticleJSON, _ := ioutil.ReadAll(resp.Body)
	var findArticle Article
	json.Unmarshal(findArticleJSON, &findArticle)
	if findArticle != testArticle {
		t.Error("Couldn't find or parse article after adding via REST.")
		return
	}
}
