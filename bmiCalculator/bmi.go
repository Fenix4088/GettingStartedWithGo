package bmicalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func Bmi() {
	fmt.Println("-----BMI Calculator-----")

	 weight, _ := strconv.ParseFloat(prompt("Please enter your weight (kg): "), 32)

	 height, _ := strconv.ParseFloat(prompt("Please enter your height (m): "), 32)

	 fmt.Printf("Your BMI is: %.2f\n", calcBmi(weight, height))
}

func prompt(question string) string {
	fmt.Print(question)
	result, _ := reader.ReadString('\n')
	/*
		-1 meens that we whant to replace all matches in a string
		in our case all \n to empty string
	*/
	return strings.Replace(result, "\n", "", -1) 
}

func calcBmi(weight, height float64) float64 {
	return weight / height * height
}