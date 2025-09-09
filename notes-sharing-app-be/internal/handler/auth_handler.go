package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"notes-app/internal/models"
	"notes-app/internal/repository"
	"notes-app/internal/service"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()	

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": http.StatusBadRequest,
				"message":    "Invalid input",
			})
			return
		}

		if err := validate.Struct(user); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": http.StatusBadRequest,
				"message":    err.Error(),
			})
			return
		}

		if err := service.RegisterUser(db, user); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": http.StatusInternalServerError,
				"message":    "Failed to register user",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusCreated,
			"message":    "User registered successfully",
		})
	}
}


func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Invalid input",
			})
			return
		}

		if err := validate.Struct(user); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		// Authenticate user and generate access token
		accessToken, err := service.LoginUser(db, user.Username, user.Password)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": "Invalid username or password",
			})
			return
		}
		
		userFetch, err := repository.GetUserByUsername(db, user.Username)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": http.StatusNotFound,
				"message":    "User not found",
			})
			return
		}

		refreshToken, err := service.GenerateRefreshToken(db, userFetch.ID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			// log.Println("Error creating refresh token:", err)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Failed to generate refresh token",
			})
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   false, 
			SameSite: http.SameSiteStrictMode,
			Expires:  time.Now().Add(15 * time.Minute),
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteStrictMode,
			Expires:  time.Now().Add(7 * 24 * time.Hour),
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Login successful",
		})
	}
}


func RefreshTokenHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var req struct {
		// 	Token string `json:"refresh_token"`
		// }

		cookie, err := r.Cookie("refresh_token")
		if err != nil {
			http.Error(w, "No refresh token provided", http.StatusUnauthorized)
			return
		}
		refreshToken := cookie.Value
		// if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Token == "" {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	json.NewEncoder(w).Encode(map[string]interface{}{
		// 		"status":  http.StatusBadRequest,
		// 		"message": "Invalid request",
		// 	})
		// 	return
		// }

		accessToken, err := service.RefreshAccessToken(db, refreshToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": err.Error(),
			})
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().Add(15 * time.Minute),
		})

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Access token refreshed",
		})
	}
}


func GetProfileHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value("user_id").(string)
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": http.StatusUnauthorized,
				"message":    "Unauthorized: no user ID in context",
				"data":       nil,
			})
			return
		}

		user, err := service.GetUserProfile(db, userID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": http.StatusNotFound,
				"message":    "User not found",
				"data":       nil,
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"message":    "Profile fetched successfully",
			"data":       user,
		})
	}
}


func LogoutHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		refreshCookie, err := r.Cookie("refresh_token")
		if err == nil {
			_ = service.RevokeRefreshTokenService(db, refreshCookie.Value)
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
			MaxAge:   -1,
			SameSite: http.SameSiteStrictMode,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
			MaxAge:   -1,
			SameSite: http.SameSiteStrictMode,
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Logout successful",
		})
	}
}



