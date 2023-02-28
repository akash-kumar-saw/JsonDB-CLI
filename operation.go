package main

import "fmt"

func createUser() {
	fmt.Println("Create User")
	fmt.Println("-----------")
	var user User
	fmt.Print("Name: ")
	fmt.Scanln(&user.Name)
	fmt.Print("Email: ")
	fmt.Scanln(&user.Email)
	fmt.Print("Phone: ")
	fmt.Scanln(&user.Phone)

	// Load existing users from JSON file
	users, err := loadUsers()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add new user to slice of users
	users = append(users, user)

	// Save updated users to JSON file
	saveUsers(users)

	fmt.Println("User created successfully!")
}

func readUser() {
	fmt.Println("Read User")
	fmt.Println("---------")
	fmt.Print("Enter email: ")
	var email string
	fmt.Scanln(&email)

	// Load existing users from JSON file
	users, err := loadUsers()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find user with matching email
	var user User
	for _, u := range users {
		if u.Email == email {
			user = u
			break
		}
	}

	fmt.Println()
	fmt.Println("User Details")
	fmt.Println("------------")
	if user.Email == "" {
		fmt.Println("User not found.")
	} else {
		fmt.Printf("Name: %s\nEmail: %s\nPhone: %s\n", user.Name, user.Email, user.Phone)
	}
}

func updateUser() {
	fmt.Println("Update User")
	fmt.Println("-----------")
	fmt.Print("Enter email: ")
	var email string
	fmt.Scanln(&email)

	// Load existing users from JSON file
	users, err := loadUsers()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find index of user with matching email
	var index int
	for i, u := range users {
		if u.Email == email {
			index = i
			break
		}
	}

	fmt.Println()
	if users[index].Email == "" {
		fmt.Println("User not found.")
	} else {
		var user User
		fmt.Print("Enter new name (press enter to skip): ")
		fmt.Scanln(&user.Name)
		fmt.Print("Enter new phone number (press enter to skip): ")
		fmt.Scanln(&user.Phone)

		if user.Name != "" {
			// Update user details
			users[index].Name = user.Name
		}

		if user.Phone != "" {
			// Update user details
			users[index].Phone = user.Phone
		}

		// Save updated users to JSON file
		saveUsers(users)

		fmt.Println()
		fmt.Println("User updated successfully!")
	}
}

func deleteUser() {
	fmt.Println("Delete User")
	fmt.Println("-----------")
	fmt.Print("Enter email: ")
	var email string
	fmt.Scanln(&email)

	// Load existing users from JSON file
	users, err := loadUsers()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find index of user with matching email
	var index int
	for i, u := range users {
		if u.Email == email {
			index = i
			break
		}
	}

	fmt.Println()
	if index == len(users) {
		fmt.Println("User not found.")
		return
	}

	// Remove user from slice of users
	users = append(users[:index], users[index+1:]...)

	// Save updated users to JSON file
	err = saveUsers(users)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User deleted successfully!")
}
