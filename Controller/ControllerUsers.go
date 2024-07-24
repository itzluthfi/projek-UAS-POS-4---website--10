package Controller

import (
	"THR/Node"
	"log"
)

var Users = []Node.NodeUser{
	{"Admin", "admin123", "admin"},
	{"Adam", "adam123", "admin"},
	{"Habib", "kasir123", "kasir"},
	{"Luthfi", "kasir321", "kasir"},
}

func GetAllUsers() []Node.NodeUser {
	return Users
}

func InsertUser(username, password, role string) {
	newUser := Node.NodeUser{
		Username: username,
		Password: password,
		Role:     role,
	}
	Users = append(Users, newUser)
}

func GetUserByUsername(username string) Node.NodeUser {
	for _, user := range Users {
		if user.Username == username {
			return user
		}
	}
	return Node.NodeUser{}
}

func UpdateUser(oldUsername, newUsername, password, role string) {
	for idx, user := range Users {
		if user.Username == oldUsername {
			Users[idx].Username = newUsername
			Users[idx].Password = password
			Users[idx].Role = role

			// Debugging: Log update confirmation
			log.Printf("User updated: %+v\n", Users[idx])
			break
		}
	}
}


func DeleteUser(username string) {
	for idx, user := range Users {
		if user.Username == username {
			Users = append(Users[:idx], Users[idx+1:]...)
			break
		}
	}
}

func VerifikasiUser(username, password string) (bool, string) {
	for _, user := range Users {
		if user.Username == username && user.Password == password {
			return true, user.Role
		}
	}
	return false, ""
}
