package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Starting http server")

	route := http.NewServeMux()

	route.HandleFunc("GET /users", getUsers)
	route.HandleFunc("POST /users", addUser)

	err := http.ListenAndServe(":808", route)
	if err != nil {
		fmt.Println("error occurred while starting the server")
	}

	fmt.Println("Server started on port 8080")
}

type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []Users{}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET requets recevived")
	result, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(result)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST requets recevived")
	user := Users{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error occurred in reading body")
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("error occurred in unmarshaling body")
	}

	users = append(users, user)

	res, _ := json.Marshal(users)
	w.Write(res)
}
