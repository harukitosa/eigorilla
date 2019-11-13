package handler

import (
	"app/eigorilla/server/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	// postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//sqlite3
	_ "github.com/mattn/go-sqlite3"
)

//curlの操作を覚えてテストできるようにする。

// CreateRoom insert new room date.
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(DatabaseName, DatabaseURL)
	if err != nil {
		panic("We can't open database!（WritePost）")
	}
	defer db.Close()
	var room model.Room
	log.Println("Post: CreateRoom")
	vars := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&room)
	if err != nil {
		w.Write([]byte("json decode error" + error.Error() + "\n"))
	}
	log.Printf("CreateData: %+v", room)
	userID := vars["userID"]
	db.Create(&model.Room{
		ID:      GenerateID(),
		Date:    GenerateDate(),
		Title:   room.Title,
		Explain: room.Explain,
		UserID:  userID,
	})
}

// GetRoomList is return all room list
func GetRoomList(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(DatabaseName, DatabaseURL)
	if err != nil {
		panic("We can't open database!（GetRoomList）")
	}
	defer db.Close()
	log.Println("GET : GetRoomList")
	var rooms []model.Room
	db.Where(&model.Room{}).Find(&rooms)
	json.NewEncoder(w).Encode(rooms)
}

// GetOneRoom is return One room data.
func GetOneRoom(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(DatabaseName, DatabaseURL)
	if err != nil {
		panic("We can't open database!（GetOneRoom）")
	}
	defer db.Close()
	log.Println("GET : GetOneRoom")
	var room model.Room
	vars := mux.Vars(r)
	roomID := vars["roomID"]
	db.Where(&model.Room{ID: roomID}).Find(&room)
	json.NewEncoder(w).Encode(room)
}

