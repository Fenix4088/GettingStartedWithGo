package bmicalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)

func Bmi() {
	fmt.Printf("-----%v-----", appName)

	 weight, _ := strconv.ParseFloat(prompt(weightMsg), 32)

	 height, _ := strconv.ParseFloat(prompt(heightMsg), 32)

	 fmt.Printf("Your BMI is: %.2f\n", calcBmi(weight, height))
}
