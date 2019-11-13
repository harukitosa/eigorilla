package main

import (
	"app/eigorilla/server/handler"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Origin", "Content-Type", "X-Requested-with", "Authorization"})

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set")
	}
	handler.DBInit()
	r := mux.NewRouter()
	r.HandleFunc("/post/:{userID}", handler.WritePost)
	r.HandleFunc("/get/timeline", handler.GetTimeLine)
	r.HandleFunc("/get/userpost/:{userID}", handler.GetUserPost)
	r.HandleFunc("/post/login/:{userID}", handler.UserCheck)
	r.HandleFunc("/post/room/:{userID}", handler.CreateRoom)
	r.HandleFunc("/get/allroom", handler.GetRoomList)
	r.HandleFunc("/get/oneroom/:{roomID}", handler.GetOneRoom)
	r.HandleFunc("/get/roompost/:{roomID}", handler.GetRoomPost)
	log.Printf("server start port localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))

}
