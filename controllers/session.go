package controllers

import (
	"encoding/json"
	"goaml-api/config"
	"goaml-api/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterHandlers registers handlers with the given router.
func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/session/login", Login).Methods("POST")
	r.HandleFunc("/session/logout", Logout).Methods("POST")
}

// Login handler for user login
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	log.Println("Mencoba mengambil data pengguna dari database...")

	// Validate user credentials from the database
	var fetchedUser models.User
	if err := config.DB.Where("username = ?", user.Username).First(&fetchedUser).Error; err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"status": false})
		return
	}
	log.Println("Berhasil mengambil data pengguna dari database")

	if user.Password != fetchedUser.Password || user.CustomerID != fetchedUser.CustomerID {
		json.NewEncoder(w).Encode(map[string]interface{}{"status": false})
		return
	}

	// If credentials are valid, create session
	session := models.Session{
		Username:   user.Username,
		CustomerID: user.CustomerID,
		SessionID:  "123456", // Mock session ID
	}
	log.Println("Membuat sesi baru untuk pengguna:", user.Username)

	config.DB.Create(&session)
	json.NewEncoder(w).Encode(map[string]interface{}{"status": true, "session_id": session.ID})
}

// Logout handler for user logout
func Logout(w http.ResponseWriter, r *http.Request) {
	var session models.Session
	json.NewDecoder(r.Body).Decode(&session)

	// Delete session from database
	config.DB.Where("username = ? AND customer_id = ? AND session_id = ?", session.Username, session.CustomerID, session.SessionID).Delete(&session)
	json.NewEncoder(w).Encode(map[string]interface{}{"status": true})
}
