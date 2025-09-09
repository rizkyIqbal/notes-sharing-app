package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"notes-app/internal/models"
	"notes-app/internal/service"
	"notes-app/pkg/jwt"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// var validate = validator.New()


func NotesRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/notes", AuthMiddleware(db, CreateNoteHandler(db))).Methods("POST")
	r.HandleFunc("/notes", GetAllNotesHandler(db)).Methods("GET")
	r.HandleFunc("/notes/user", AuthMiddleware(db, GetNotesByUserIDHandler(db))).Methods("GET")
	r.HandleFunc("/notes/images/{id}", AuthMiddleware(db, DeleteImageHandler(db))).Methods("DELETE")
	r.HandleFunc("/notes/{id}", GetNoteHandler(db)).Methods("GET")
	r.HandleFunc("/notes/{id}", AuthMiddleware(db, OwnershipMiddleware(db, UpdateNoteHandler(db)))).Methods("PATCH")
	r.HandleFunc(
		"/notes/{id}",
		AuthMiddleware(db, OwnershipMiddleware(db, DeleteNoteHandler(db))),
	).Methods("DELETE")
	r.HandleFunc("/notes/{id}/images", AddImagesHandler(db)).Methods("POST")
	r.HandleFunc("/notes/{id}/images", GetImagesHandler(db)).Methods("GET")

	return r
}

func GetAllNotesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")  
		pageStr := r.URL.Query().Get("page")
		limitStr := r.URL.Query().Get("limit")

		page := 1
		limit := 10
		var err error

		if pageStr != "" {
			page, err = strconv.Atoi(pageStr)
			if err != nil || page < 1 {
				page = 1
			}
		}

		if limitStr != "" {
			limit, err = strconv.Atoi(limitStr)
			if err != nil || limit < 1 {
				limit = 10
			}
		}

		offset := (page - 1) * limit

		notes, err := service.GetAllNotesService(db, title, limit, offset)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":     http.StatusInternalServerError,
				"message":    err.Error(),
				"data":       nil,
			})
			return
		}

		total, err := service.GetTotalNotesCountService(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":     http.StatusInternalServerError,
				"message":    "Failed to fetch total count",
				"data":       nil,
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":     http.StatusOK,
			"message":    "Notes fetched successfully",
			"data": map[string]interface{}{
				"page":  page,
				"limit": limit,
				"total": total,
				"notes": notes,
			},
		})
	}
}

func GetNotesByUserIDHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value("user_id").(string)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
				"data":    nil,
			})
			return
		}

		title := r.URL.Query().Get("title")  
		pageStr := r.URL.Query().Get("page")
		limitStr := r.URL.Query().Get("limit")

		page := 1
		limit := 10
		var err error

		if pageStr != "" {
			page, err = strconv.Atoi(pageStr)
			if err != nil || page < 1 {
				page = 1
			}
		}

		if limitStr != "" {
			limit, err = strconv.Atoi(limitStr)
			if err != nil || limit < 1 {
				limit = 10
			}
		}

		offset := (page - 1) * limit

		notes, err := service.GetNotesByUserIDService(db, userID, title, limit, offset)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		total, err := service.GetTotalNotesByUserIDCountService(db, userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Failed to fetch total count",
				"data":    nil,
			})
			return
		}

		// Wrap data
		data := map[string]interface{}{
			"notes": notes,
			"page":  page,
			"limit": limit,
			"total": total,
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Notes fetched successfully",
			"data":    data,
		})
	}
}

func GetNoteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// id, err := strconv.Atoi(mux.Vars(r)["id"])
		id := mux.Vars(r)["id"]
		if _, err := uuid.Parse(id); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Invalid note ID",
				"data":    nil,
			})
			return
		}

		note, err := service.GetNoteService(db, id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusNotFound,
				"message": "Note not found",
				"data":    nil,
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Note fetched successfully",
			"data":    note,
		})
	}
}


func CreateNoteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var note models.Note
		userID, ok := r.Context().Value("user_id").(string)

		if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Invalid input",
				"data":    nil,
			})
			return
		}

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized",
				"data":    nil,
			})
			return
		}
		note.UserID = userID

		if err := validate.Struct(note); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		id, err := service.CreateNoteService(db, note)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusInternalServerError,
				"message": "Failed to create note",
				"data":    nil,
			})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusCreated,
			"message": "Note created successfully",
			"data": map[string]string{
				"id": id,
			},
		})
	}
}


func UpdateNoteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value("user_id").(string)
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": http.StatusUnauthorized,
				"message":    "Unauthorized",
			})
			return
		}

		noteID := mux.Vars(r)["id"]
		// noteID, err := strconv.Atoi(idStr)
		if _, err := uuid.Parse(noteID); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": http.StatusBadRequest,
				"message":    "Invalid note ID",
			})
			return
		}

		var note models.Note
		if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": http.StatusBadRequest,
				"message":    "Invalid input",
			})
			return
		}

		note.ID = noteID
		note.UserID = userID

		if err := validate.Struct(note); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		if err := service.UpdateNoteService(db, note); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":     http.StatusInternalServerError,
				"message":    "Failed to update note",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":     http.StatusOK,
			"message":    "Note updated successfully",
		})
	}
}


func DeleteNoteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		noteID := mux.Vars(r)["id"]
		// noteID, err := strconv.Atoi(idStr)
		if _, err := uuid.Parse(noteID); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": http.StatusBadRequest,
				"message":    "Invalid note ID",
			})
			return
		}

		if err := service.DeleteNoteService(db, noteID); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": http.StatusInternalServerError,
				"message":    "Failed to delete note",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": http.StatusOK,
			"message":    "Note deleted successfully",
		})
	}
}


func OwnershipMiddleware(db *sql.DB, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value("user_id").(string)
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": http.StatusUnauthorized,
				"message":    "Unauthorized: missing user ID",
			})
			return
		}

		vars := mux.Vars(r)
		noteID := vars["id"]

		var owner string
		err := db.QueryRow("SELECT user_id FROM notes WHERE id = $1", noteID).Scan(&owner)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": http.StatusNotFound,
				"message":    "Note not found",
			})
			return
		}

		if owner != userID {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": http.StatusForbidden,
				"message":    "Forbidden: you do not own this note",
			})
			return
		}

		next(w, r)
	}
}


func AuthMiddleware(db *sql.DB, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		cookie, err := r.Cookie("access_token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized: no token",
				"data":    nil,
			})
			return
		}

		userID, err := jwt.ValidateToken(cookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": "Unauthorized: invalid token",
				"data":    nil,
			})
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", userID)

		next(w, r.WithContext(ctx))
	}
}

