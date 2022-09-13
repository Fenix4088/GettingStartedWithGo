package structures

import "time"

type User struct {
	firstName, lastName, birthDate string
	created time.Time
}

func createNewUser(fName, lName, bDate string) *User {
	return &User{
		fName,
		lName,
		bDate,
		time.Now(),
	}
}