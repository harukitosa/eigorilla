package handler

import (
	"app/eigorilla/server/model"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	// postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//sqlite3
// _ "github.com/mattn/go-sqlite3"

//Mode has two status PRODUCTION or DEPLOY
// var Mode string

// DatabaseName has database name
var DatabaseName string

// DatabaseURL has database url
var DatabaseURL string

// DBInit initialize your datebase and migrate.
func DBInit() {
	// Mode = "PRODUCTION"
	// if Mode == "PRODUCTION" {
	// 	DatabaseURL = "test.sqlite3"
	// 	DatabaseName = "sqlite3"
	// } else if Mode == "DEPLOY" {
	DatabaseURL = os.Getenv("DATABASE_URL")
	DatabaseName = "postgres"
	// }

	db, err := gorm.Open(DatabaseName, DatabaseURL)
	if err != nil {
		panic("We can't open database!（dbInit）")
	}
	//残りのモデルはまだ入れてない。
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Room{})
	defer db.Close()
}

// WritePost insert new post date.
func WritePost(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	db, err := gorm.Open(DatabaseName, DatabaseURL)
	if err != nil {
		panic("We can't open database!（WritePost）")
	}
	defer db.Close()

	log.Println("POST: WritePost")

	vars := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&post)
	if err != nil {
		w.Write([]byte("json decode error" + error.Error() + "\n"))
	}

	log.Printf("postdata:%+v", post)
	userID := vars["userID"]

	db.Create(&model.Post{
		ID:       GenerateID(),
		Sentence: post.Sentence,
		Date:     GenerateDate(),
		UserID:   userID,
		RoomID:   post.RoomID,
	})

	log.Printf("post:%v", post)
}

// GetTimeLine get all post date
func GetTimeLine(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(DatabaseName, DatabaseURL)
	if err != nil {
		panic("We can't open database!（GetTimeLine）")
	}
	defer db.Close()

	log.Println("GET : GetTimeLine")

	var posts []model.Post
	db.Where(&model.Post{}).Find(&posts)

	//postのデータをすべて取った後userデータ総当り
	var sendPosts []model.SendPost
	for i := 0; i < len(posts); i++ {
		var sendPost model.SendPost
		sendPost.PostID = posts[i].ID
		sendPost.Sentence = posts[i].Sentence
		sendPost.Date = posts[i].Date
		sendPost.UserID = posts[i].UserID
		sendPost.RoomID = posts[i].RoomID
		var user model.User
		db.Where(&model.User{ID: posts[i].UserID}).First(&user)
		sendPost.UserPhotoURL = user.PhotoURL
		sendPost.UserName = user.DisplayName
		sendPosts = append(sendPosts, sendPost)
	}
	json.NewEncoder(w).Encode(sendPosts)
}

// GetRoomPost return all post data in the room
func GetRoomPost(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(DatabaseName, DatabaseURL)
	if err != nil {
		panic("We can't open database!（GetRoomPost）")
	}
	defer db.Close()
	log.Println("GET : GetRoomPost")
	vars := mux.Vars(r)
	roomID := vars["roomID"]
	var posts []model.Post
	db.Where(&model.Post{RoomID: roomID}).Find(&posts)
	//その後データベースからユーザーデータを取ってくる
	var user model.User
	//postのデータをすべて取った後userデータ総当り
	var sendPosts []model.SendPost
	if len(posts) != 0 {
		db.Where(&model.User{ID: posts[0].UserID}).First(&user)
		for i := 0; i < len(posts); i++ {
			var sendPost model.SendPost
			sendPost.PostID = posts[i].ID
			sendPost.Sentence = posts[i].Sentence
			sendPost.Date = posts[i].Date
			sendPost.UserID = posts[i].UserID
			sendPost.RoomID = posts[i].RoomID
			sendPost.UserPhotoURL = user.PhotoURL
			sendPost.UserName = user.DisplayName
			sendPosts = append(sendPosts, sendPost)
		}
	}
	json.NewEncoder(w).Encode(sendPosts)
}

// GetUserPost get all user's post date
func GetUserPost(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(DatabaseName, DatabaseURL)
	if err != nil {
		panic("We can't open database!（GetTimeLine）")
	}
	defer db.Close()

	log.Println("GET : GetUserPost")

	vars := mux.Vars(r)
	userID := vars["userID"]
	var posts []model.Post
	db.Where(&model.Post{UserID: userID}).Find(&posts)
	//その後データベースからユーザーデータを取ってくる
	var user model.User
	//postのデータをすべて取った後userデータ総当り
	var sendPosts []model.SendPost
	if len(posts) != 0 {
		db.Where(&model.User{ID: posts[0].UserID}).First(&user)
		for i := 0; i < len(posts); i++ {
			var sendPost model.SendPost
			sendPost.PostID = posts[i].ID
			sendPost.Sentence = posts[i].Sentence
			sendPost.Date = posts[i].Date
			sendPost.UserID = posts[i].UserID
			sendPost.RoomID = posts[i].RoomID
			sendPost.UserPhotoURL = user.PhotoURL
			sendPost.UserName = user.DisplayName
			sendPosts = append(sendPosts, sendPost)
		}
	}
	json.NewEncoder(w).Encode(sendPosts)
}

//UserCheck はユーザーが新規かそうでないかを調べる。
func UserCheck(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(DatabaseName, DatabaseURL)
	if err != nil {
		panic("We can't open database!（UserCheck）")
	}
	defer db.Close()
	log.Println("----------------------------------")
	log.Println("POST :UserCheck")
	vars := mux.Vars(r)
	userID := vars["userID"]
	var user model.User
	db.Where(&model.User{ID: userID}).First(&user)
	if user.ID == "" {
		var newUser model.User
		decoder := json.NewDecoder(r.Body)
		error := decoder.Decode(&newUser)
		if err != nil {
			w.Write([]byte("json decode error" + error.Error() + "\n"))
		}
		log.Printf("%+v", newUser)
		db.Create(&model.User{
			ID:          userID,
			DisplayName: newUser.DisplayName,
			PhotoURL:    newUser.PhotoURL,
			Profile:     newUser.Profile,
		})

		log.Println("NEW USER")
		log.Println(userID)
	}
	log.Println("END CHECK USER")
	log.Println("----------------------------------")
}
