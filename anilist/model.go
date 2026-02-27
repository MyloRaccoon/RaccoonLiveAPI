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

type Anime struct {
	ID int
	Title string
	Cover string
	Genres []string

}

func GetFavoritesAnime(username string) ([]Anime, error) {
	query := `
	query GetUserFavorites($username: String!) {
		User(name: $username) {
			name
			favourites {
				anime {
					nodes {
						id
						title {
							romaji
							english
							native
						}
						coverImage {
							large
						}
						averageScore
						genres
					}
				}
			}
		}
	}
	`
	variables := map[string]any{
		"username": username,
	}
	body, _ := json.Marshal(map[string]any{
		"query": query,
		"variables": variables,
	})

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return []Anime{}, err
	}
	defer resp.Body.Close()

	var data struct {
		Data struct {
			User struct {
				Favourites struct {
					Anime struct {
						Nodes []struct {
							ID int `json:"id"`
							Title struct {
								Romaji string `json:"romaji"`
							} `json:"title"`
							CoverImage struct {
								Large string `json:"large"`
							} `json:"coverImage"`
							Genres []string `json:"genres"`
						} `json:"nodes"`
					} `json:"anime"`
				} `json:"favourites"`
			} `json:"user"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return []Anime{}, err
	}

	favs := []Anime{}
	for _, animeData := range data.Data.User.Favourites.Anime.Nodes {
		favs = append(favs, Anime{
			ID: animeData.ID,
			Title: animeData.Title.Romaji,
			Cover: animeData.CoverImage.Large,
			Genres: animeData.Genres,
		})
	}

	return favs, nil
}