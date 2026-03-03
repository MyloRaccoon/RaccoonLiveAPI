package github

type GithubProfile struct {
	Name string `json:"login"`
	Avatar string `json:"avatar_url"`
	ID int `json:"id"`
	URL string `json:"html_url"`
}

type GithubRepo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Language string `json:"language"`
	URL string `json:"html_url"`
}