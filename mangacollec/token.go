package mangacollec

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

const tokenUrl = "https://api.mangacollec.com/oauth/token"
const bootstrapMode = false

type TokenResponse struct {
    AccessToken  string `json:"access_token"`
    TokenType    string `json:"token_type"`
    ExpiresIn    int    `json:"expires_in"`
    RefreshToken string `json:"refresh_token"`
    CreatedAt    int64  `json:"created_at"`
}

type TokenStore struct {
    AccessToken  string    `json:"access_token"`
    RefreshToken string    `json:"refresh_token"`
    ExpiresAt    time.Time `json:"expires_at"`
    ClientID     string    `json:"client_id"`
    ClientSecret string    `json:"client_secret"`
}

func refreshAccessToken(store *TokenStore)  error {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", store.RefreshToken)
	data.Set("client_id", store.ClientID)
	data.Set("client_secret", store.ClientSecret)

	resp, err := http.PostForm(tokenUrl, data)
	if err != nil {
		return fmt.Errorf("Request failed: %w", err)
	}
	defer resp.Body.Close()

	var token TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return fmt.Errorf("Couldn't decode token response: %w", err)
	}

	store.AccessToken = token.AccessToken
	store.RefreshToken = token.RefreshToken
	store.ExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)

	return nil
}

func getValidToken(store *TokenStore) (string, error) {
	if time.Now().Add(5 * time.Minute).After(store.ExpiresAt) {
		if err := refreshAccessToken(store); err != nil {
			return "", nil
		}
	}
	return store.AccessToken, nil
}

func Bootstrap() *TokenStore {
    store := &TokenStore{
        RefreshToken: os.Getenv("MC_REFRESH_TOKEN"),
        ClientID:     os.Getenv("MC_CLIENT_ID"),
        ClientSecret: os.Getenv("MC_CLIENT_SECRET"),
        ExpiresAt:    time.Now(),
    }
    return store
}