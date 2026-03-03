package music

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func GetMusics() ([]Music, error) {

	data, err := ioutil.ReadFile("./music/musics.json")
	if err != nil {
		log.Fatal("Error when opening musics file: ", err)
		return []Music{}, err
	}

	var musics []Music
	err = json.Unmarshal(data, &musics)
	if err != nil {
		log.Fatal("Error during unmarshaling: ", err)
		return []Music{}, err
	}

	return musics, nil
}

func GetMusicById(id string) (Music, error) {
	musics, err := GetMusics()
	if err != nil {
		log.Fatal("Error getting music: %s", err)
		return Music{}, err
	}
	for _, music := range musics {
		if music.ID == id {
			return music, nil
		}
	}
	return Music{}, fmt.Errorf("Error getting music: id '%s' doesn't exists.", id)
}

func PostMusic(music Music) error {

	musics, err := GetMusics()
	if err != nil {
		log.Fatal("Error getting musics: ", err)
		return err
	}

	musics = append(musics, music)

	data, err := json.Marshal(musics)
	if err != nil {
		log.Fatal("Error while encoding musics: ", err)
		return err
	}

	err = os.WriteFile("./music/musics.json", data, 0644)
	if err != nil {
		log.Fatal("Error while saving musics: ", err)
	}
	return err
}