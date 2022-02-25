package controller

import (
	"encoding/json"
	"net/http"
)

func successResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": http.StatusOK,
		"data": data,
	})
}

func errorResponse(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": status,
		"message": message,
	})
}
