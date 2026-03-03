package anilist

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const URL = "https://graphql.anilist.co"

func getUserID(username string) (int, error) {
	query := getUserIDQuery
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

func getLastActivity(userID int) (AnilistActivity, error) {
	query := getLastActivityQuery
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
		SiteURL: siteUrl,
	}, nil
}

func getFavoritesAnime(username string) ([]Anime, error) {
	query := getFavoritesAnimeQuery
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
							SiteURL string `json:"siteUrl"`
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
			SiteURL: animeData.SiteURL,
		})
	}

	return favs, nil
}

func getFavoritesManga(username string) ([]Manga, error) {
	query := getFavoritesMangaQuery
	variables := map[string]any{
		"username": username,
	}
	body, _ := json.Marshal(map[string]any{
		"query": query,
		"variables": variables,
	})

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return []Manga{}, err
	}
	defer resp.Body.Close()

	var data struct {
		Data struct {
			User struct {
				Favourites struct {
					Manga struct {
						Nodes []struct {
							ID int `json:"id"`
							Title struct {
								Romaji string `json:"romaji"`
							} `json:"title"`
							CoverImage struct {
								Large string `json:"large"`
							} `json:"coverImage"`
							Genres []string `json:"genres"`
							SiteURL string `json:"siteUrl"`
						} `json:"nodes"`
					} `json:"manga"`
				} `json:"favourites"`
			} `json:"user"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return []Manga{}, err
	}

	favs := []Manga{}
	for _, mangaData := range data.Data.User.Favourites.Manga.Nodes {
		favs = append(favs, Manga{
			ID: mangaData.ID,
			Title: mangaData.Title.Romaji,
			Cover: mangaData.CoverImage.Large,
			Genres: mangaData.Genres,
			SiteURL: mangaData.SiteURL,
		})
	}

	return favs, nil
}

func getFavoritesCharacters(username string) ([]Character, error) {
	query := getFavoritesCharactersQuery
	variables := map[string]any{
		"username": username,
	}
	body, _ := json.Marshal(map[string]any{
		"query": query,
		"variables": variables,
	})

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return []Character{}, err
	}
	defer resp.Body.Close()

	var data struct {
		Data struct {
			User struct {
				Favourites struct {
					Characters struct {
						Nodes []struct {
							ID int `json:"id"`
							Name struct {
								Full string `json:"full"`
							} `json:"name"`
							Description string `json:"Description"`
							Gender string `json:"gender"`
							DateOfBirth struct {
								Year int `json:"year"`
								Month int `json:"month"`
								Day int `json:"day"`
							} `json:"dateOfBirth"`
							Age string `json:"age"`
							BloodType string `json:"bloodType"`
							Image struct {
								Large string `json:"large"`
							} `json:"image"`
							SiteURL string `json:"siteUrl"`
							Media struct {
								Nodes []struct {
									ID int `json:"id"`
									Title struct {
										Romaji string `json:"romaji"`
									} `json:"title"`
									SiteURL string `json:"siteUrl"`
								} `json:"nodes"`
							} `json:"media"`
						} `json:"nodes"`
					} `json:"characters"`
				} `json:"favourites"`
			} `json:"user"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return []Character{}, err
	}

	favs := []Character{}
	for _, charData := range data.Data.User.Favourites.Characters.Nodes {
		mediaData := charData.Media.Nodes[0]
		favs = append(favs, Character{
			ID: charData.ID,
			Name: charData.Name.Full,
			Age: charData.Age,
			Gender: charData.Gender,
			BloodType: charData.BloodType,
			Description: charData.Description,
			Image: charData.Image.Large,
			SiteURL: charData.SiteURL,
			MediaID: mediaData.ID,
			MediaTitle: mediaData.Title.Romaji,
			MediaURL: mediaData.SiteURL,
		})
	}

	return favs, nil
}

func getFavoritesStaff(username string) ([]Staff, error) {
	query := getFavoritesStaffQuery
	variables := map[string]any{
		"username": username,
	}
	body, _ := json.Marshal(map[string]any{
		"query": query,
		"variables": variables,
	})

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return []Staff{}, err
	}
	defer resp.Body.Close()

	var data struct {
		Data struct {
			User struct {
				Favourites struct {
					Staff struct {
						Nodes []struct {
							ID int `json:"id"`
							Name struct {
								Full string `json:"full"`
							} `json:"name"`
							Description string `json:"Description"`
							SiteURL string `json:"siteUrl"`
						} `json:"nodes"`
					} `json:"staff"`
				} `json:"favourites"`
			} `json:"user"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return []Staff{}, err
	}

	favs := []Staff{}
	for _, staffData := range data.Data.User.Favourites.Staff.Nodes {
		favs = append(favs, Staff{
			ID: staffData.ID,
			Name: staffData.Name.Full,
			Description: staffData.Description,
			SiteURL: staffData.SiteURL,
		})
	}

	return favs, nil
}

func getFavoritesStudio(username string) ([]Studio, error) {
	query := getFavoritesStudiosQuery
	variables := map[string]any{
		"username": username,
	}
	body, _ := json.Marshal(map[string]any{
		"query": query,
		"variables": variables,
	})

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return []Studio{}, err
	}
	defer resp.Body.Close()

	var data struct {
		Data struct {
			User struct {
				Favourites struct {
					Studios struct {
						Nodes []struct {
							ID int `json:"id"`
							Name string `json:"name"`
							SiteURL string `json:"siteUrl"`
						} `json:"nodes"`
					} `json:"studios"`
				} `json:"favourites"`
			} `json:"user"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return []Studio{}, err
	}

	favs := []Studio{}
	for _, staffData := range data.Data.User.Favourites.Studios.Nodes {
		favs = append(favs, Studio{
			ID: staffData.ID,
			Name: staffData.Name,
			SiteURL: staffData.SiteURL,
		})
	}

	return favs, nil
}