package model

//Post is someone's post date.
type Post struct {
	ID       string `json:"id" gorm:"PRIMARY_KEY"`
	Sentence string `json:"sentence"`
	Date     string `json:"date"`
	UserID   string `json:"userID"`
}

//User is date
type User struct {
	ID          string `json:"id" gorm:"PRIMARY_KEY"`
	DisplayName string `json:"displayName"`
	PhotoURL    string `json:"photoURL"`
	Profile     string `json:"profile"`
}

//SendPost is sent by json data
type SendPost struct {
	PostID       string `json:"postID"`
	Sentence     string `json:"sentence"`
	Date         string `json:"date"`
	UserID       string `json:"userID"`
	UserPhotoURL string `json:"userPhotoURL"`
	UserName     string `json:"userName"`
}
