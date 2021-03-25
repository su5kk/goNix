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

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	handleError(err)

	res, getErr := spaceClient.Do(req)
	handleError(getErr)
	posts := []post{}
	jsonFile, readErr := ioutil.ReadAll(res.Body)
	handleError(readErr)
	jsonErr := json.Unmarshal(jsonFile, &posts)
	handleError(jsonErr)
	for _, elem := range posts {
		data, err := json.MarshalIndent(elem, "", "\t")
		handleError(err)
		fmt.Println(string(data))
	}
}
