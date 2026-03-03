package music

type Music struct {
	ID string
	Title string
	Artist string
	Cover string
	URL string
	ListenDate struct {
		Hour int
		Minute int
		Day int
		Month int
		Year int
	}
}