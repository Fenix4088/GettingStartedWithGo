package bmicalculator

import (
	"fmt"
	"strings"
)

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