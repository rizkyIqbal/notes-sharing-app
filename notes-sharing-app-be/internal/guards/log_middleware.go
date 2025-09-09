package guards

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"notes-app/internal/models"
	"notes-app/internal/repository"
)

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func newResponseRecorder(w http.ResponseWriter) *responseRecorder {
	return &responseRecorder{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		body:           new(bytes.Buffer),
	}
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}

func (rr *responseRecorder) Write(b []byte) (int, error) {
	rr.body.Write(b) // capture response body
	return rr.ResponseWriter.Write(b)
}

func LoggingMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		if r.Method == http.MethodOptions {	
			next.ServeHTTP(w, r)
			return
		}	

		var requestBody []byte
		if r.Body != nil {
			requestBody, _ = io.ReadAll(r.Body)
			r.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		headers := make(map[string]string)
		for k, v := range r.Header {
			val := strings.Join(v, ", ")
			if strings.ToLower(k) == "authorization" {
				if len(val) > 10 {
					val = val[:10] + "****" 
				} else {
					val = "****"
				}
			}
			headers[k] = val
		}
		headersJSON, _ := json.Marshal(headers)

		// --- Wrap response writer ---
		rec := newResponseRecorder(w)

		// --- Call next handler ---
		next.ServeHTTP(rec, r)

		// --- Build log ---
		log := models.Log{
			Datetime:     start,
			Method:       r.Method,
			Endpoint:     r.URL.Path,
			Headers:      string(headersJSON),
			Payload:      string(requestBody),
			ResponseBody: rec.body.String(),
			StatusCode:   rec.statusCode,
		}

		// --- Store log in DB ---
		_ = repository.InsertLog(db, log)
	})
}
