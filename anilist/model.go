package anilist

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const URL = "https://graphql.anilist.co"

func GetUserID(username string) (int, error) {
	query := `
	query ($name: String!) {
		User(name: $name) {
			id
		}
	}
	`
	variables := map[string]any{
		"name": username,
	}
	body, _ := json.Marshal(map[string]any{
		"query": query,
		"variables": variables,
	})

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	
	var result struct {
		Data struct {
			User struct {
				ID int `json:"id"`
			} `json:"User"`
		} `json:"Data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Data.User.ID, nil
}

func GetLastActivity(userID int) (AnilistActivity, error) {
	query := `
	query ($userId: Int) {
	  Page(page: 1, perPage: 2) {
		activities(userId: $userId, sort: ID_DESC, type: MEDIA_LIST) {
		  ... on ListActivity {
		    id
			status
			progress
			media {
			  title { romaji }
			  siteUrl
			}
		  }
		}
	  }
	}
	`
	variables := map[string]any {
		"userId": userID,
	}
	body, _ := json.Marshal(map[string]any{
		"query": query,
		"variables": variables,
	})

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return AnilistActivity{}, nil
	}
	defer resp.Body.Close()

	var result struct {
		Data struct {
			Page struct {
				Activities []struct {
					ID int `json:"id"`
					Status string `json:"status"`
					Progress string `json:"progress"`
					Media struct {
						Title struct {
							Romaji string `json:"romaji"`
						} `json:"title"`
						SiteURL string `json:"siteUrl"`
					} `json:"media"`
				} `json:"activities"`
			} `json:"Page"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return AnilistActivity{}, nil
	}

	activityData := result.Data.Page.Activities[0]
	status := activityData.Status
	progress := activityData.Progress
	siteUrl := activityData.Media.SiteURL
	title := activityData.Media.Title.Romaji

	return AnilistActivity{
		Title: title,
		Status: status,
		Progress: progress,
		SiteUrl: siteUrl,
	}, nil
}