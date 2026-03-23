package mangacollec

import (
	"encoding/json"
	"net/http"
	"os"
)

const URL = "https://api.mangacollec.com"

func fetchData() (MangaCollec, error) {
	token := os.Getenv("MANGACOLLEC_TOKEN")

	req, _ := http.NewRequest("GET", URL + "/v2/user/mylo/collection", nil)
	req.Header.Set("Authorization", "Bearer " + token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return MangaCollec{}, err
	}
	defer resp.Body.Close()

	var data MangaCollecData
	json.NewDecoder(resp.Body).Decode(&data)

	mangaCollec := MangaCollec{}
	
	for _, serie := range data.Series {
		mangaCollec.Series = append(mangaCollec.Series, Serie{
			ID: serie.ID,
			Title: serie.Title,
			AdultContent: serie.AdultContent,
		})
	}

	for _, edition := range data.Editions {
		publisher, _ := getPublisherName(edition.PublisherID)
		mangaCollec.Editions = append(mangaCollec.Editions, Edition{
			ID: edition.ID,
			Title: edition.Title,
			SerieID: edition.SeriesID,
			Publisher: publisher,
			NotFinished: edition.NotFinished,
		})
	}

	for _, volume := range data.Volumes {
		mangaCollec.Volumes = append(mangaCollec.Volumes, Volume{
			ID: volume.ID,
			Title: volume.Title,
			Number: volume.Number,
			ReleaseDate: volume.ReleaseDate,
			EditionID: volume.EditionID,
			ImageURL: volume.ImageURL,
		})
	}

	for _, boxEdition := range data.BoxEditions {
		publisher, _ := getPublisherName(boxEdition.PublisherID)
		mangaCollec.BoxEditions = append(mangaCollec.BoxEditions, BoxEdition{
			ID: boxEdition.ID,
			Title: boxEdition.Title,
			Publisher: publisher,
			AdultContent: boxEdition.AdultContent,
		})
	}

	for _, box := range data.Boxes {
		mangaCollec.Boxes = append(mangaCollec.Boxes, Box{
			ID: box.ID,
			Title: box.Title,
			ReleaseDate: box.ReleaseDate,
			CommercialStop: box.CommercialStop,
			ImageURL: box.ImageURL,
		})
	}

	for _, boxVolume := range data.BoxVolumes {
		mangaCollec.BoxVolumes = append(mangaCollec.BoxVolumes, BoxVolume{
			ID: boxVolume.ID,
			BoxID: boxVolume.BoxID,
			VolumeID: boxVolume.VolumeID,
		})
	}

	for _, possession := range data.Possessions {
		mangaCollec.Possessions = append(mangaCollec.Possessions, Possession{
			ID: possession.ID,
			VolumeID: possession.VolumeID,
			Date: possession.CreatedAt,
		})
	}

	for _, boxPossession := range data.BoxPossessions {
		mangaCollec.BoxPossessions = append(mangaCollec.BoxPossessions, BoxPossession{
			ID: boxPossession.ID,
			BoxID: boxPossession.BoxID,
			Date: boxPossession.CreatedAt,
		})
	}

	return mangaCollec, err
}

func getPublisherName(id string) (string, error) {
	token := os.Getenv("MANGACOLLEC_TOKEN")

	req, _ := http.NewRequest("GET", URL + "v2/publishers/" + id, nil)
	req.Header.Set("Authorization", "Bearer " + token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data struct {
		Publishers []struct {
			Title string `json:"title"`
		}
	}
	json.NewDecoder(resp.Body).Decode(&data)

	return data.Publishers[0].Title, nil
}

func getVolumesInPossession() ([]Volume, error) {
	data, err := fetchData()
	if err != nil {
		return []Volume{}, err
	}

	volumes := []Volume{}
	for _, possession := range data.Possessions {
		volumes = append(volumes, getVolume(data.Volumes, possession.VolumeID))
	}

	return volumes, nil
}

func getVolume(volumes []Volume, id string) Volume {
	result := Volume{}
	for _, volume := range volumes {
		if volume.ID == id {
			result = volume
		}
	}
	return result
}

// func getSeries() ([]Serie, error) {
// 	data, err := fetchData()
// 	if err != nil {
// 		return []Serie{}, err
// 	}

// 	return data.Series, nil
// }

// func getEditions() ([]Edition, error) {
// 	data, err := fetchData()
// 	if err != nil {
// 		return []Edition{}, err
// 	}

// 	return data.Editions, err
// }

// func getVolumes() ([]Volume, error) {
// 	data, err := fetchData()
// 	if err != nil {
// 		return []Volume{}, err
// 	}
// 	return data.Volumes, err
// }

// func getBoxEditions() ([]BoxEdition, error) {
// 	data, err :=
// }

// func getPossessions() ([]Possessions, error) {
// 	data, err := fetchData()
// 	if err != nil {
// 		return []Possessions{}, err
// 	}

// 	result := []Possessions{}
// 	for _, possession := range data.Possessions {
// 		result = append(result, Possessions{
// 			ID: possession.ID,
// 			Volume: Vo,
// 		})
// 	}
// }