package main

import (
	"fmt"
	"log"
	"net/http"

	"notes-app/internal/config"
	"notes-app/internal/guards"
	"notes-app/internal/handler"

	"github.com/rs/cors"
	// "notes-app/internal/gua"
)

func main() {
	config.InitConfig()
	db := config.InitDB()
	defer db.Close()

	r := handler.NotesRouter(db)

	r.HandleFunc("/register", handler.RegisterHandler(db)).Methods("POST")
	r.HandleFunc("/login", handler.LoginHandler(db)).Methods("POST")
	r.HandleFunc("/profile", handler.AuthMiddleware(db, handler.GetProfileHandler(db))).Methods("GET")
	r.HandleFunc("/logout", handler.AuthMiddleware(db, handler.LogoutHandler(db))).Methods("POST")
	r.HandleFunc("/auth/refresh", handler.RefreshTokenHandler(db)).Methods("POST")



	// http.Handle("/", JSONMiddleware(r))

	loggedRouter := guards.LoggingMiddleware(db, r) // log requests & responses
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(loggedRouter)

	// handler := c.Handler(r)
	fmt.Println("Server running at :4000")
	// log.Fatal(http.ListenAndServe(":8080", r))
	log.Fatal(http.ListenAndServe("0.0.0.0:4000", corsHandler))
}
