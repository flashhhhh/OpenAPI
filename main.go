package main

import (
	"encoding/json"
	bloomfilter "flash/bloomFilter"
	"fmt"
	"net/http"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Users struct {
	ID int `gorm:"primaryKey"`
	Username string `gorm:"unique,not null"`
	Password string `gorm:"not null"`
	Name string `gorm:"not null"`
}

func main() {
	db_path := "host=localhost port=5432 user=postgres password=12345678 dbname=swagger_example"
	db, err := gorm.Open(postgres.Open(db_path), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Users{})

	// Load all usernames to bloom filter
	bf := bloomfilter.NewBloomFilter(10000000, 3)
	var names []string
	db.Table("users").Pluck("username", &names)

	for _, name := range names {
		bf.Add(name)
	}

	/*
		POST /users/register
	*/
	http.HandleFunc("/users/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			username := r.FormValue("username")
			password := r.FormValue("password")
			name := r.FormValue("name")

			if !bf.Contains(username) {
				bf.Add(username)
				db.Create(&Users{Username: username, Password: password, Name: name})
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte("User created."))

				fmt.Println("Bloom filer HIT!")
			} else {
				// Check db if exist user
				var user Users
				db.Where("username = ?", username).First(&user)

				if user.Username == username {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte("User already exist."))

					fmt.Println("Bloom filer HIT!")
				} else {
					user := Users{Username: username, Password: password, Name: name}
					db.Create(&user)
					w.WriteHeader(http.StatusCreated)

					fmt.Println("Bloom filer MISS!")
				}
			}
		}
	})

	/*
		POST /users/login
	*/
	http.HandleFunc("/users/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			username := r.FormValue("username")
			password := r.FormValue("password")

			if !bf.Contains(username) {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Username does not exist."))

				fmt.Println("Bloom filer HIT!")
			} else {
				var user Users
				db.Where("username = ?", username).First(&user)

				if user.Username == username {
					fmt.Println("Bloom filer HIT!")

					if (user.Password != password) {
						w.WriteHeader(http.StatusUnauthorized)
						w.Write([]byte("Invalid password."))
					} else {
						w.WriteHeader(http.StatusOK)
						w.Write([]byte("Login successfully!"))
					}
				} else {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Username does not exist."))

					fmt.Println("Bloom filer MISS!")
				}
			}
		}
	})

	/*
		GET /users
	*/

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var users []Users
			db.Select("id", "username", "name").Order("id asc").Find(&users)
			
			jsonData, err := json.Marshal(users)
			if err != nil {
				http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		}
	})

	/*
		GET /users/
	*/
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			pathParts := strings.Split(r.URL.Path, "/")
			if len(pathParts) != 3 || pathParts[2] == "" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Invalid URL."))
				return
			}

			user_id := pathParts[2]

			var user Users
			db.Table("users").Where("id = ?", user_id).Find(&user)

			if user.ID == 0 {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("User not found."))
			} else {
				jsonData, err := json.Marshal(user)
				if err != nil {
					http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(jsonData)
			}
		}
	})
	
	http.ListenAndServe(":1234", nil)
}