package main

import (
	"fmt"
	"hello-go/gpa"

	"github.com/fatih/color"
)

func main() {
	defer color.Blue("Program finished")

	for i := 1; i <= 3; i++ {
		var percent int

		fmt.Printf("Enter percentage #%d: ", i)
		fmt.Scan(&percent)

		grade, gpaValue, description := gpa.Calculate(percent)

		color.Cyan("Grade: %s", grade)
		color.Green("GPA: %.2f", gpaValue)
		color.Yellow("Description: %s", description)
		fmt.Println()
	}
}
