package git

type Repos struct {
	Name string
}

func RepoSearch(username string) []Repos {
	var repository []Repos

	url := "https://api.github.com/users/" + username + "/repos"

	Http_request(url, &repository)

	// for _, repo := range repository {
	// 	fmt.Printf("- %s :\n", repo.Name)
	// }

	return repository

}
