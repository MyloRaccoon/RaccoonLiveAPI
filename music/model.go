package music

import (
	"encoding/json"
	"fmt"
	"os"
)

func saveMusic(musics []Music) error {
	data, err := json.Marshal(musics)
	if err != nil {
		return err
	}

	err = os.WriteFile("./music/musics.json", data, 0644)
	if err != nil {
		return err
	}
	return err
}

func getMusics() ([]Music, error) {

	data, err := os.ReadFile("./music/musics.json")
	if err != nil {
		return []Music{}, err
	}

	var musics []Music
	err = json.Unmarshal(data, &musics)
	if err != nil {
		return []Music{}, err
	}

	return musics, nil
}

func getMusicById(id string) (Music, error) {
	musics, err := getMusics()
	if err != nil {
		return Music{}, err
	}
	for _, music := range musics {
		if music.ID == id {
			return music, nil
		}
	}
	return Music{}, fmt.Errorf("Error getting music: id '%s' doesn't exists.", id)
}

func postMusic(music Music) error {
	if _, err := getMusicById(music.ID); err == nil {
		return fmt.Errorf("Error posting music: id '%s' already exists.", music.ID)
	}

	musics, err := getMusics()
	if err != nil {
		return err
	}

	musics = append(musics, music)

	return  saveMusic(musics)
}

func deleteMusicById(id string) (Music, error) {
	musics, err := getMusics()
	if err != nil {
		return Music{}, err
	}

	new_musics := []Music{}
	removed_music := Music{}

	for _, music := range musics {
		if music.ID != id {
			new_musics = append(new_musics, music)
		} else {
			removed_music = music
		}
	}

	err = saveMusic(new_musics)
	if err != nil {
		return Music{}, err
	}

	if removed_music.ID == "" {
		return Music{}, fmt.Errorf("Error deleting music: music of id '%s' does not exists.", id)
	}

	return removed_music, nil
}