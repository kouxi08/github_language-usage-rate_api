package git

import (
	// "main-module/server"

	"math"
)

// func Main() {
// 	username := "kouxi08"
// 	processIfNotDoneRecently(username)
// }

// func processIfNotDoneRecently(username string) {
// 	// データベースなどから直近の処理日時を取得するロジックを実装してください

// 	previous := ""

// 	// 現在時刻と1時間前の時刻を取得
// 	currentTime := time.Now()
// 	oneHourAgo := previous.Add(time.Hour)

// 	// 直近の処理日時が1時間前よりも前の場合にのみ処理を実行
// 	if currentTime > oneHourAgo {
// 		// 処理を実行するロジックをここに書いてください

// 	} else {
// 		// データベースなどに処理日時を保存するロジックを実装してください

// 	}
// }

func User_lang() map[string]float64 {
	totalCount := 0
	username := "kouxi08"
	reponaum := 0
	stocks := map[string]float64{}

	//repos.goからユーザのリポジトリを取得
	re := RepoSearch(username)
	// fmt.Println(re)
	// fmt.Print(re)
	for _, repo := range re {
		url := "https://api.github.com/repos/" + username + "/" + repo.Name + "/languages"
		reponaum += 1

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

	// fmt.Print(reponaum)
	// // fmt.Println(totalCount)
	for languages, count := range stocks {
		percentage := math.Round((float64(count)/float64(totalCount)*100)*10) / 10
		stocks[languages] = percentage
	}
	// stocks["count"] = float64(reponaum)

	// server.Insert(username, stocks)

	return stocks
}
