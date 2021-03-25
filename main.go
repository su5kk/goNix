package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	concurrentGet()
}

func concurrentGet() {
	url := "https://jsonplaceholder.typicode.com/posts"
	ch := make(chan post)
	posts := []post{}
	for i := 1; i <= 100; i++ {
		go makeRequest(url, ch, i)
		posts = append(posts, <-ch)
	}

	for _, pst := range posts {
		fmt.Println(pst)
	}
}

func makeRequest(url string, ch chan post, i int) {
	url += fmt.Sprintf("/%v", i)
	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	response, _ := spaceClient.Do(request)
	jsonFile, _ := ioutil.ReadAll(response.Body)
	pst := post{}
	jsonErr := json.Unmarshal(jsonFile, &pst)
	handleError(jsonErr)
	ch <- pst
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
