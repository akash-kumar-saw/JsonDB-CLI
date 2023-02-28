package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func loadUsers() ([]User, error) {
	_, err := os.Stat("users.json")
	if err != nil {
		return []User{}, nil
	}

	data, err := ioutil.ReadFile("users.json")
	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func saveUsers(users []User) error {
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("users.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}
