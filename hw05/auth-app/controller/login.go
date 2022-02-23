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
					errorResponse(w, x)
					break
				case error:
					errorResponse(w, x.Error())
					break
				default:
					errorResponse(w, "unknown error")
					break
				}
			}
		}()

		err := r.ParseForm()
		if err != nil {
			errorResponse(w, err.Error())
			return
		}

		sessionCookie, err := r.Cookie("session")
		if err == nil && sessionCookie != nil {
			session := sessionRepository.GetSession(sessionCookie.Value)
			if session != nil {
				newSessionCookie := http.Cookie{
					Name:     "session",
					Value:    session.Id,
					Path: "/",
					Expires:  session.ExpiresIn,
					HttpOnly: true,
				}

				http.SetCookie(w, &newSessionCookie)
				successResponse(w, "Already authenticated!")
				return
			}
		}

		email := r.Form.Get("email")
		if email == "" {
			errorResponse(w, "Email is empty")
			return
		}

		password := r.Form.Get("password")
		if password == "" {
			errorResponse(w, "Password is empty")
			return
		}

		user, err := getUserByEmail(email, password)
		if err != nil {
			errorResponse(w, err.Error())
			return
		}

		newSession := sessionRepository.CreateSession(*user)
		newSessionCookie := http.Cookie{
			Name:     "session",
			Value:    newSession.Id,
			Path: "/",
			Expires:  newSession.ExpiresIn,
			HttpOnly: true,
		}

		err = sessionRepository.StoreSession(newSession)
		if err != nil {
			errorResponse(w, err.Error())
			return
		}

		http.SetCookie(w, &newSessionCookie)
		successResponse(w, "Authenticated!")
	}
}

func getUserByEmail(email, password string) (user *entity.User, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
		}
	}()

	endpoint := fmt.Sprintf("%s/users?email=%s&password=%s",
		os.Getenv("USER_SERVICE"),
		email,
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