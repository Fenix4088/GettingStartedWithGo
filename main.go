package main

import (
	"fmt"
	"gettingStartedWithGo/basics"
	bmicalculator "gettingStartedWithGo/bmiCalculator"
)

func main() {
	basics.IntroduceYourself(basics.Bio{"John", "Dou", "USA", 2000})

	circ, explenation := basics.CalcCircumference(22)

	fmt.Println(circ)
	fmt.Println(explenation)
	bmicalculator.Bmi()

}