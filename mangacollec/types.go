package mangacollec

type MangaCollecData struct{
	Editions []struct {
		ID string `json:"id"`
		Title string `json:"title"`
		SeriesID string `json:"series_id"`
		PublisherID string `json:"publisher_id"`
		NotFinished bool `json:"not_finished"`
	} `json:"editions"`
	Series []struct {
		ID string `json:"id"`
		Title string `json:"title"`
		AdultContent bool `json:"adult_content"`
	} `json:"series"`
	Volumes []struct {
		ID string `json:"id"`
		Title string `json:"title"`
		Number int `json:"number"`
		ReleaseDate string `json:"release_date"`
		EditionID string `json:"edition_id"`
		ImageURL string `json:"image_url"`
	} `json:"volumes"`
	BoxEditions []struct {
		ID string `json:"id"`
		Title string `json:"title"`
		PublisherID string `json:"publisher_id"`
		AdultContent bool `json:"adult_content"`	
	} `json:"box_editions"`
	Boxes []struct {
		ID string `json:"id"`
		Title string `json:"title"`
		ReleaseDate string `json:"release_date"`
		CommercialStop bool `json:"commercial_stop"`
		BoxEditionID string `json:"box_edition_id"`
		ImageURL string `json:"iamge_url"`
	} `json:"boxes"`
	BoxVolumes []struct {
		ID string `json:"id"`
		BoxID string `json:"box_id"`
		VolumeID string `json:"volume_id"`
	} `json:"box_volumes"`
	Possessions []struct {
		ID string `json:"id"`
		VolumeID string `json:"volume_id"`
		CreatedAt string `json:"created_at"`
	} `json:"possessions"`
	BoxPossessions []struct {
		ID string `json:"id"`
		BoxID string `json:"box_id"`
		CreatedAt string `json:"created_at"`
	} `json:"box_possessions"`
}

type MangaCollec struct {
	Editions []Edition
	Series []Serie
	Volumes []Volume
	BoxEditions []BoxEdition
	Boxes []Box
	BoxVolumes []BoxVolume
	Possessions []Possession
	BoxPossessions []BoxPossession
}

type Serie struct {
	ID string
	Title string
	AdultContent bool
}

type Edition struct {
	ID string
	Title string
	SerieID string
	Publisher string
	NotFinished bool
}

type Volume struct {
	ID string
	Title string
	Number int
	ReleaseDate string
	EditionID string
	ImageURL string
}

type BoxEdition struct {
	ID string
	Title string
	Publisher string
	AdultContent bool
}

type Box struct {
	ID string
	Title string
	ReleaseDate string
	CommercialStop bool
	BoxEdition BoxEdition
	ImageURL string
}

type BoxVolume struct {
	ID string
	BoxID string
	VolumeID string
}

type Possession struct {
	ID string
	VolumeID string
	Date string
}

type BoxPossession struct {
	ID string
	BoxID string
	Date string
}
