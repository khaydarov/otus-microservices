package controller

import (
	"encoding/json"
	"net/http"
)

func successResponse(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"data": data,
	})
}

func errorResponse(w http.ResponseWriter, message string) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 500,
		"message": message,
	})
}
