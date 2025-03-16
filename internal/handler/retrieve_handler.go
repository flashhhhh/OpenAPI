package handler

import (
	"encoding/json"
	"fmt"
	"go-swag/internal/services"
	"net/http"
	"strconv"
	"strings"
)

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "GET") {
		users, err := services.GetAllUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		var userArray []map[string]string
		for _, userData := range users {
			user := map[string]string{
				"id":       fmt.Sprintf("%d", userData.ID),
				"username": userData.Username,
				"name":     userData.Name,
			}
			userArray = append(userArray, user)
		}

		usersJSON, err := json.Marshal(userArray)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(usersJSON)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "GET") {
		pathParts := strings.Split(r.URL.Path, "/")
		id, err := strconv.Atoi(pathParts[len(pathParts)-1])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		user, err := services.GetUserByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		userData := map[string]string{
			"id":       fmt.Sprintf("%d", user.ID),
			"username": user.Username,
			"name":     user.Name,
		}

		userJSON, err := json.Marshal(userData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(userJSON)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}