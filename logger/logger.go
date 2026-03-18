package logger

import (
	"os"
	"time"
)

const LOG_PATH = "raccoon-live-api.log"

func Log(message string) error {
	message = time.Now().String() + " | " + message

	f, err := os.OpenFile(LOG_PATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString("\n" + message)
	return err
}