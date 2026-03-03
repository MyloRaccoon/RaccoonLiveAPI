package anilist

type AnilistActivity struct {
	Title string
	Status string
	Progress string
	SiteURL string
}

type Anime struct {
	ID int
	Title string
	Cover string
	Genres []string
	SiteURL string
}

type Manga struct {
	ID int
	Title string
	Cover string
	Genres []string
	SiteURL string
}

type Character struct {
	ID int
	Name string
	Age string
	Gender string
	BloodType string
	Description string
	Image string
	SiteURL string
	MediaID int
	MediaTitle string
	MediaURL string
}

type Staff struct {
	ID int
	Name string
	Description string
	SiteURL string
}

type Studio struct {
	ID int
	Name string
	SiteURL string
}