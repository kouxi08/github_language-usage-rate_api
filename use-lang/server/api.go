package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"main-module/git"
)

//	type Task struct {
//		Lang    string  `json:"langage"`
//		Ratio   float64 `json:"ratio"`
//	}
type Task struct {
	Count     int `json:"count"`
	Languages []struct {
		Lang  string  `json:"language"`
		Ratio float64 `json:"ratio"`
	} `json:"lang"`
}

type Count struct {
	Number int `json:"number"`
}

type Tasks []Task
type Counts []Count

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
	langages := git.User_lang()

	var task Task
	// var counts Counts

	for lang, ratio := range langages {
		if lang == "count" {
			task.Count = int(ratio)
		} else {
			task.Languages = append(task.Languages, struct {
				Lang  string  `json:"language"`
				Ratio float64 `json:"ratio"`
			}{
				Lang:  lang,
				Ratio: ratio,
			})
		}
	}
	json.NewEncoder(w).Encode(task)
}
