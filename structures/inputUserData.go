package structures

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func InputUserData() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	newUser := createNewUser(firstName, lastName, birthDate)

	newUser.outputUserData()
	
}

func (user *User) outputUserData() {
	fmt.Printf("My name is %v %v (born on %v)", user.firstName, user.lastName, user.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)

	userInput, _ := reader.ReadString('\n')

	var cleanedInput string

	if runtime.GOOS == "windows" {
		cleanedInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		cleanedInput = strings.Replace(userInput, "\n", "", -1)
	}

	return cleanedInput
}

