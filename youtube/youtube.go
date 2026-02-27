package youtube

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YoutubeVideo struct {
	Title string
	Date string
	Description string
	Thumbnail string
	ID string
	URL string
}

func GetLastVideo() (YoutubeVideo, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return YoutubeVideo{}, err
	}

	ctx := context.Background()

	apiKey := os.Getenv("GOOGLE_API_KEY")
	channelID := os.Getenv("YTB_CHANNEL_ID")

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creation service: %v", err)
		return YoutubeVideo{}, err
	}

	uploadsPlaylistID := "UU" + channelID[2:]
	call := service.PlaylistItems.List([]string{"snippet"}).
		PlaylistId(uploadsPlaylistID).
		MaxResults(1)

	resp, err := call.Do()
	if err != nil {
		log.Fatalf("Error calling API: %v", err)
		return YoutubeVideo{}, err
	}

	videoData := resp.Items[0]
	snippet := videoData.Snippet

	video := YoutubeVideo{
		Title: snippet.Title,
		Date: snippet.PublishedAt,
		Description: snippet.Description,
		Thumbnail: snippet.Thumbnails.High.Url,
		ID: snippet.ResourceId.VideoId,
		URL: fmt.Sprintf("https://www.youtube.com/watch?v=%s", snippet.ResourceId.VideoId),
	}

	return video, nil
}