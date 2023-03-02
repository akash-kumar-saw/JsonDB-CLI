# JsonDB CLI
The JsonDB CLI is a simple command line application that allows users to create, read, update, and delete user records in a JSON file. The application is built using Go and utilizes a simple user struct with name, email, and phone number fields.

## Installation
- Clone the repository from GitHub:
```bash
git clone https://github.com/akash-kumar-saw/JsonDB-CLI
```
- Install the required dependencies:
```bash
cd JsonDB-CLI
go mod download
```
- Run the following command to build the application:
```bash
go build .
```
- Run the following command to start the application:
```bash
JsonDB-CLI
```

## Features
The JsonDB CLI provides the following features:
- Create user records with name, email, and phone number fields
- Read user records by email address
- Update user records with new name, email, and phone number fields
- Delete user records by email address
- Store user records in a JSON file for persistence
- Simple command line interface for ease of use

## Usage
When you run the JsonDB CLI, you will be prompted to select an operation from a list of options:
```bash
Welcome to the JsonDB CLI

Please select an operation:
1. Create
2. Read
3. Update
4. Delete
5. Exit

Enter your choice (1-5):
```
To perform an operation, simply enter the corresponding number and follow the prompts.
### Create User
To create a new user, select option 1 from the main menu. You will be prompted to enter the user's name, email, and phone number. Once you have entered this information, the user will be added to the database and you will see a success message.
### Read User
To read a user's details, select option 2 from the main menu. You will be prompted to enter the user's email address. If the user exists in the database, their name, email, and phone number will be displayed.
### Update User
To update a user's details, select option 3 from the main menu. You will be prompted to enter the user's email address, followed by the new name, email, and phone number (if applicable). Once you have entered this information, the user's details will be updated in the database and you will see a success message.
### Delete User
To delete a user, select option 4 from the main menu. You will be prompted to enter the user's email address. If the user exists in the database, they will be deleted and you will see a success message.
### Exit
To exit the application, select option 5 from the main menu.

## Contribution
Contributions to the JsonDB-CLI are welcome. If you have a bug fix, feature request, or other improvement, please submit a pull request or open an issue.

