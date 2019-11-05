package main

import(
	"os"
	"log"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"app/eigorilla/server/handler"
	"net/http"
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
	log.Printf("server start port localhost:"+port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))

}