package main

import (
	"fmt"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	fmt.Println("Welcome to the JsonDB CLI")

	for {
		fmt.Println()
		fmt.Println("Please select an operation:")
		fmt.Println("1. Create")
		fmt.Println("2. Read")
		fmt.Println("3. Update")
		fmt.Println("4. Delete")
		fmt.Println("5. Exit")
		fmt.Println()
		fmt.Print("Enter your choice (1-5): ")

		var choice string
		fmt.Scanln(&choice)
		fmt.Println()

		switch choice {
		case "1":
			createUser()
		case "2":
			readUser()
		case "3":
			updateUser()
		case "4":
			deleteUser()
		case "5":
			fmt.Println("Exiting JsonDB CLI... Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid command.")
		}
		fmt.Printf("\nDo you want to continue (Y/N): ")
		var option string
		fmt.Scanln(&option)

		if option == "Y" || option == "y" {
			continue
		}
		if option == "N" || option == "n" {
			fmt.Println()
			fmt.Println("Exiting JsonDB CLI... Goodbye!")
			break
		}
	}
}
