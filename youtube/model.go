package youtube

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const POOLING_MAX = 10

func getLastVideo() (YoutubeVideo, error) {
	err := godotenv.Load()
	if err != nil {
		return YoutubeVideo{}, err
	}

	ctx := context.Background()

	apiKey := os.Getenv("GOOGLE_API_KEY")
	channelID := os.Getenv("YTB_CHANNEL_ID")

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return YoutubeVideo{}, err
	}

	uploadsPlaylistID := "UU" + channelID[2:]
	call := service.PlaylistItems.List([]string{"snippet"}).
		PlaylistId(uploadsPlaylistID).
		MaxResults(POOLING_MAX)

	resp, err := call.Do()
	if err != nil {
		return YoutubeVideo{}, err
	}

	videoIDs := make([]string, 0, len(resp.Items))
	for _, item := range resp.Items {
		videoIDs = append(videoIDs, item.Snippet.ResourceId.VideoId)
	}

	videosCall := service.Videos.List([]string{"snippet", "liveStreamingDetails"}).
		Id(strings.Join(videoIDs, ","))
	videosResponse, err := videosCall.Do()

	for _, videoData := range videosResponse.Items {
		if videoData.LiveStreamingDetails == nil {
			return YoutubeVideo {
				Title: videoData.Snippet.Title,
				Date: videoData.Snippet.PublishedAt,
				Description: videoData.Snippet.Description,
				Thumbnail: videoData.Snippet.Thumbnails.High.Url,
				ID: videoData.Id,
				URL: fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoData.Id),
			}, nil
		}
	}

	return YoutubeVideo{}, nil
}