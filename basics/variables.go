package basics

import (
	"fmt"
	"time"
)

type Bio struct {
	Name, LastName, Country string
	BirthYear int
}

func IntroduceYourself(bio Bio) {
	currentYear := time.Now().Year()

	fmt.Printf("Hello, my name is %v %v, I am %v y.o. and I was born in %v\n", bio.Name, bio.LastName, currentYear - bio.BirthYear, bio.Country)
	
}