package basics

import (
	"fmt"
	"time"
	"math"
)

type Bio struct {
	Name, LastName, Country string
	BirthYear int
}

func IntroduceYourself(bio Bio) {
	currentYear := time.Now().Year()

	fmt.Printf("Hello, my name is %v %v, I am %v y.o. and I was born in %v\n", bio.Name, bio.LastName, currentYear - bio.BirthYear, bio.Country)
	
}

func CalcCircumference(radius float32) (float32, string) {
	circumference := 2 * math.Pi * radius
	
	return circumference, fmt.Sprintf("For a radius of %v, the circle circumference is %.2f.", radius, circumference)
}