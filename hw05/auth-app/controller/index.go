package controller

import "net/http"

// Index is Main page action
func Index(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			errorResponse(w, r.(string))
		}
	}()

	str := "Hello!"
	w.Write([]byte(str))
}