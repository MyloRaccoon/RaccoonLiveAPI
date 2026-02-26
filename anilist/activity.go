package anilist

import "fmt"

type AnilistActivity struct {
	Title string
	Status string
	Progress string
	SiteUrl string
}

func PrintActivity(activity AnilistActivity) string {
	res := ""
	if activity.Progress != "" {
		res = fmt.Sprintf("%s %s of %s", activity.Status, activity.Progress, activity.Title)
	} else {
		res = fmt.Sprintf("%s %s", activity.Status, activity.Title)
	}
	return res
}