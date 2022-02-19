package controller

import (
	"auth-app/model"
	"auth-app/repository"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Login is a page action
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Println("rendering tmpl")
		tmpl, err := template.ParseFiles("template/login.html")
		if err == nil {
			tmpl.Execute(w, nil)
		}

		return
	}

	log.Println("login...")
	err := r.ParseForm()
	if err != nil {
		errorResponse(w, err.Error())
		return
	}

	email := r.Form.Get("email")
	user := getUserByEmail(email)
	log.Println("user", user)
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	session := repository.CreateSession(*user)
	sessionCookie := http.Cookie{
		Name:     "session",
		Value:    session.Id,
		Expires:  session.ExpiresIn,
		HttpOnly: true,
	}

	err = repository.StoreSession(session)
	if err != nil {
		errorResponse(w, err.Error())
		return
	}

	http.SetCookie(w, &sessionCookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func getUserByEmail(email string) (user *model.User) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recovered: %s\n", r)
		}
	}()

	user = &model.User{}
	log.Printf("Request: %s/users?email=%s", os.Getenv("USER_SERVICE"), email)
	response, err := http.Get(fmt.Sprintf("%s/users?email=%s", os.Getenv("USER_SERVICE"), email))
	if err != nil {
		log.Printf("Response: %s\n", err)
		return nil
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(user)

	log.Println(user)
	return user
}