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

func Test() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	ctx := context.Background()

	apiKey := os.Getenv("GOOGLE_API_KEY")
	channelID := os.Getenv("YTB_CHANNEL_ID")

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creation service: %v", err)
	}

	// fonctionne probablement mais moins eco (100 unités de cout)
	// call := service.Search.List([]string{"snippet"}).
	// 	ChannelId(channelID).
	// 	MaxResults(1).
	// 	Order("date").
	// 	Type("video")

	uploadsPlaylistID := "UU" + channelID[2:]
	call := service.PlaylistItems.List([]string{"snippet"}).
		PlaylistId(uploadsPlaylistID).
		MaxResults(1)

	resp, err := call.Do()
	if err != nil {
		log.Fatalf("Error calling API: %v", err)
	}

	video := resp.Items[0]
	snippet := video.Snippet

	fmt.Printf("Titre       : %s\n", snippet.Title)
	fmt.Printf("Date        : %s\n", snippet.PublishedAt)
	fmt.Printf("Description : %s\n", snippet.Description)
	fmt.Printf("Thumbnail   : %s\n", snippet.Thumbnails.High.Url)
	fmt.Printf("Video ID    : %s\n", snippet.ResourceId.VideoId)
	fmt.Printf("URL         : https://www.youtube.com/watch?v=%s\n", snippet.ResourceId.VideoId)
}