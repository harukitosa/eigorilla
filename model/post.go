package model

//Post is someone's post date.
type Post struct {
	ID string `json:"id"`
	Sentence string `json:"sentence"`
	Date string `json:"date"`
	UserID string `json:"userID"`
    UserPhotoURL string `json:"userPhotoURL"`
    UserName string `json:"userName"`
}
