package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

const baseURL = "http://localhost:8000"

func TestGetUsers(t *testing.T) {
	response, err := http.Get(baseURL + "/users")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("Response code should be %d but got %d", http.StatusOK, response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		t.Fatal(err)
	}

	if len(users) == 0 {
		t.Fatal("Users should not be empty")
	}
}

func TestGetUser(t *testing.T) {
	// Add a user before getting it
	addUser(t)

	response, err := http.Get(baseURL + "/users/1")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("Response code should be %d but got %d", http.StatusOK, response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		t.Fatal(err)
	}

	if user.ID != 1 || user.Username != "testUser" || user.Email != "test@example.com" {
		t.Fatalf("User should be {ID: 1, Username: 'testUser', Email: 'test@example.com'} but got %v", user)
	}
}

func addUser(t *testing.T) {
	user := User{
		Username: "testUser",
		Email:    "test@example.com",
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	response, err := http.Post(baseURL+"/users", "application/json", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		t.Fatalf("Response code should be %d but got %d", http.StatusCreated, response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	var addedUser User
	err = json.Unmarshal(body, &addedUser)
	if err != nil {
		t.Fatal(err)
	}

	if addedUser.ID != 1 || addedUser.Username != "testUser" || addedUser.Email != "test@example.com" {
		t.Fatalf("User should be {ID: 1, Username: 'testUser', Email: 'test@example.com'} but got %v", addedUser)
	}
}