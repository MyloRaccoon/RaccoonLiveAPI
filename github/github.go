package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

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

func GetUser(login string) (GithubProfile, error) {
	token := os.Getenv("GITHUB_TOKEN")

	req, _ := http.NewRequest("GET", "https://api.github.com/users/" + login, nil)
	req.Header.Set("Authorization", "Bearer " + token)
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error while getting profile")
		return GithubProfile{}, err
	}
	defer resp.Body.Close()

	var result GithubProfile
	json.NewDecoder(resp.Body).Decode(&result)
	
	return result, nil
}

func GetLastRepo(login string) (GithubRepo, error) {
	token := os.Getenv("GITHUB_TOKEN")

	url := "https://api.github.com/users/"  + login + "/repos?sort=created&direction=desc&per_page=1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GithubRepo{}, err
	}
	req.Header.Set("Authorization", "Bearer " + token)
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error while getting repo")
		return GithubRepo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return GithubRepo{}, fmt.Errorf("github API error: %s", body)
	}

	var repos []GithubRepo
	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err != nil {
		return GithubRepo{}, err
	}
	if len(repos) == 0 {
		return GithubRepo{}, fmt.Errorf("No repo found...")
	}

	return repos[0], nil
}