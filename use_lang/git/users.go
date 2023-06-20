package git

import (
	"fmt"
)

type Repository struct {
	Name map[string]int
}

func User_lang() {
	totalCount := 0
	username := "Yoshiki217"
	stocks := map[string]float64{}
	//repos.goからユーザのリポジトリを取得
	re := Repo_search(username)
	// fmt.Println(re)

	for _, repo := range re {
		url := "https://api.github.com/repos/" + username + "/" + repo.Name + "/languages"
		// var c []Repository
		//リクエスト作成

		languag := map[string]int{}
		Http_request(url, &languag)

		for languages, count := range languag {
			totalCount += count
			// fmt.Println(count)
			_, ok := stocks[languages]
			if !ok {
				stocks[languages] = float64(count)
			} else {
				stocks[languages] += float64(count)
			}
		}
	}

	// fmt.Println(stocks)
	// fmt.Println(totalCount)
	for languages, count := range stocks {
		percentage := float64(count) / float64(totalCount) * 100
		fmt.Printf("%s, 使用割合 %.1f\n", languages, percentage)
	}
}
