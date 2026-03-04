package music

type Date struct {
	Hour int
	Minute int
	Day int
	Month int
	Year int
}

type Music struct {
	ID string
	Title string
	Artist string
	Cover string
	URL string
	ListenDate Date
}