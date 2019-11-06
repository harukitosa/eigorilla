package handler

import (
	"app/eigorilla/server/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
    "os"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	// postgres
   	_ "github.com/jinzhu/gorm/dialects/postgres"

)
//	_ "github.com/mattn/go-sqlite3


// DatabaseName has database name
var DatabaseName string

// DatabaseURL has database url
var DatabaseURL string


// GenerateID return UUID
func GenerateID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return "err"
	}
	uu := u.String()
	return uu
}

// GenerateDate return Year Month Day Hour Minitues Seconds
func GenerateDate() string {
	const layout = "2006-01-02 15:04:05"
	t := time.Now()
	s := ""
	s = t.Format(layout)
	return s
}


// DBInit initialize your datebase and migrate.
func DBInit() {
	//DatabaseURL = "test.sqlite3"
	//DatabaseName = "sqlite3"
	DatabaseURL = os.Getenv("DATABASE_URL")
	DatabaseName = "postgres"

	db, err := gorm.Open(DatabaseName, DatabaseURL)
	if err != nil {
		panic("We can't open database!（dbInit）")
	}
	//残りのモデルはまだ入れてない。
	db.AutoMigrate(&model.Post{})
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

	userID := vars["userID"]

	db.Create(&model.Post {
		ID: GenerateID(),
		Sentence: post.Sentence,
		Date: GenerateDate(),
		UserID: userID,
        UserPhotoURL: post.UserPhotoURL,
        UserName: post.UserName,
	})
	if err != nil {
		log.Println("Warning Error WritePost!!!!!!!!")
	}
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
	json.NewEncoder(w).Encode(posts)
}

// GetUserPost get all user's post date
func GetUserPost(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(DatabaseName, DatabaseURL)
	if err != nil {
		panic("We can't open database!（GetTimeLine）")
	}
	defer db.Close()

	log.Println("GET : GetTimeLine")

	vars := mux.Vars(r)
	userID := vars["userID"]
	var posts []model.Post
	db.Where(&model.Post{UserID: userID}).Find(&posts)
	json.NewEncoder(w).Encode(posts)
}
