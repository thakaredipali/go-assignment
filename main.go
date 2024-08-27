package main

import (
	"fmt"
	"go-lang-assignments/utils"
	"strconv"
	"strings"
)

func main(){

   /// Calculate simple interest
	var input string
	fmt.Println("Enter Principal, Rate, and Time:")
	 _, err:= fmt.Scan(&input)

	if err != nil{
     fmt.Println("Invalid Input")
	 return 
	}
	values := strings.Split(input, ",")
	principal, _ := strconv.ParseFloat(values[0], 64)
	rate, _ := strconv.ParseFloat(values[1], 64)
	time, _ := strconv.ParseFloat(values[2], 64)
	simpleInterest := utils.CalculateSimpleInterest(principal,rate, time)
	fmt.Printf("Simple Interest: %.2f\n", simpleInterest)


	/// Calculate area of circle
	var radius float64
	fmt.Println("Enter radius")
	fmt.Scanln(&radius)
	area := utils.CalculateAreaOfCircle(radius)
	roundedArea := fmt.Sprintf("%.2f", area)
	fmt.Printf("Area of the circle: %s\n", roundedArea)
}