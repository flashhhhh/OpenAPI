package handler

import (
	"encoding/json"
	"fmt"
	"go-swag/internal/services"
	"io/ioutil"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	data := make(map[string]string)
	json.Unmarshal(body, &data)

	username := data["username"]
	password := data["password"]
	name := data["name"]

	fmt.Println("username: ", username)
	fmt.Println("password: ", password)
	fmt.Println("name: ", name)

	err = services.Signup(username, password, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User created successfully"))
	w.WriteHeader(http.StatusOK)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get token in cookie
	cookie, err := r.Cookie("auth_token")
	if err != http.ErrNoCookie {
		fmt.Println("cookie: ", cookie.Name, cookie.Value,  cookie.MaxAge)

		err = services.ValidateToken(cookie.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	data := make(map[string]string)
	json.Unmarshal(body, &data)

	username := data["username"]
	password := data["password"]

	err = services.Login(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	userData := []map[string]string{
		{
			"username": username,
			"password": password,
		},
	}
	
	token, err := services.GenerateToken(userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("token: ", token)
	
	// Write token to cookie
	http.SetCookie(w, &http.Cookie{
		Name: "auth_token",
		Value: token,
		MaxAge: 3600,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}