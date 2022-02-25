package controller

import (
	"auth-app/entity"
	"auth-app/repository"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// Login is an action that authenticates user by `email` and `password`
func Login(sessionRepository repository.SessionRepository) func (w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		defer func () {
			if r := recover(); r != nil {
				switch x := r.(type) {
				case string:
					errorResponse(w, x, http.StatusInternalServerError)
					break
				case error:
					errorResponse(w, x.Error(), http.StatusInternalServerError)
					break
				default:
					errorResponse(w, "unknown error", http.StatusInternalServerError)
					break
				}
			}
		}()

		err := r.ParseForm()
		if err != nil {
			errorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		username := r.Form.Get("username")
		if username == "" {
			errorResponse(w, "Username is empty", http.StatusBadRequest)
			return
		}

		password := r.Form.Get("password")
		if password == "" {
			errorResponse(w, "Password is empty", http.StatusBadRequest)
			return
		}

		user, err := getUserByUsername(username, password)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusNotFound)
			return
		}

		newSession := sessionRepository.CreateSession(*user)
		newSessionCookie := http.Cookie{
			Name:     	"session",
			Value:    	newSession.Id,
			Path: 		"/",
			Expires:  	newSession.ExpiresIn,
			HttpOnly: 	true,
		}

		err = sessionRepository.StoreSession(newSession)
		if err != nil {
			errorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &newSessionCookie)
		successResponse(w, "Authenticated!")
	}
}

func getUserByUsername(username, password string) (user *entity.User, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
		}
	}()

	endpoint := fmt.Sprintf("%s/users?username=%s&password=%s",
		os.Getenv("USER_SERVICE"),
		username,
		password,
	)

	response, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusInternalServerError {
		return nil, errors.New("internal error")
	}

	fmt.Println(endpoint, response)
	if response.StatusCode == http.StatusBadRequest {
		return nil, errors.New("login or password is not correct")
	}

	user = &entity.User{}
	err = json.NewDecoder(response.Body).Decode(&user)

	if user.Id == 0 {
		return nil, errors.New("login or password is not correct")
	}

	return user, nil
}