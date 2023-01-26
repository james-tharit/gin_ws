package models

// album represents data about a record album.
type Album struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Released int32  `json:"released"`
}
