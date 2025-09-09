package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"notes-app/internal/service"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ImageRequest struct {
	URLs []string `json:"image_path" validate:"required,min=1,dive,required"`
}


func AddImagesHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        vars := mux.Vars(r)
        noteID := vars["id"]

        var req ImageRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "status":  http.StatusBadRequest,
                "message": "Invalid request body",
            })
            return
        }

        if err := validate.Struct(req); err != nil {
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "status":  http.StatusBadRequest,
                "message": "Missing fields and cannot be empty",
            })
            return
        }

        if err := service.AddImagesToNote(db, noteID, req.URLs); err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "status":  http.StatusInternalServerError,
                "message": err.Error(),
            })
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status":  http.StatusCreated,
            "message": "Images added successfully",
        })
    }
}

func GetImagesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		noteID := vars["id"]

		images, err := service.GetImages(db, noteID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  http.StatusOK,
			"message": "Images fetched successfully",
			"data":    images,
		})
	}
}

func DeleteImageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        vars := mux.Vars(r)
        imageID := vars["id"]

        if _, err := uuid.Parse(imageID); err != nil {
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "status":  http.StatusBadRequest,
                "message": "Invalid image ID",
            })
            return
        }

        if err := service.DeleteImageService(db, imageID); err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]interface{}{
                "status":  http.StatusInternalServerError,
                "message": "Failed to delete image",
            })
            return
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status":  http.StatusOK,
            "message": "Image deleted successfully",
        })
    }
}