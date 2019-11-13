package model

//Post is someone's post date.
type Post struct {
	ID       string `json:"id" gorm:"PRIMARY_KEY"`
	Sentence string `json:"sentence"`
	Date     string `json:"date"`
	UserID   string `json:"userID"`
	RoomID   string `json:"roomID"`
}

//User is date
type User struct {
	ID          string `json:"id" gorm:"PRIMARY_KEY"`
	DisplayName string `json:"displayName"`
	PhotoURL    string `json:"photoURL"`
	Profile     string `json:"profile"`
}

//Room is room date
type Room struct {
	ID      string `json:"id" gorm:"PRIMARY_KEY"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Explain string `json:"explain"`
	UserID  string `json:"userID"`
}

//SendPost is sent by json data, not including database.
type SendPost struct {
	PostID       string `json:"postID"`
	Sentence     string `json:"sentence"`
	Date         string `json:"date"`
	UserID       string `json:"userID"`
	RoomID       string `json:"roomID"`
	UserPhotoURL string `json:"userPhotoURL"`
	UserName     string `json:"userName"`
}
