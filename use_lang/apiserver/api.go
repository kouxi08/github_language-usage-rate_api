package apiserver

import (
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	Lang string `json:"langage"`
}

func HandlerRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/artivles", returnArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnArticles(w http.ResponseWriter, r *http.Request) {
	var task []Task

}
