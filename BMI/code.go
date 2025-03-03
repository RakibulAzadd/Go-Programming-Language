package main

import "fmt"

func main(){
	var name string
	var weight float64
	var height float64

	fmt.Printf("Enter your name :")
	fmt.Scan(&name)
	fmt.Printf("Enter your weight :")
	fmt.Scan(&weight)
	fmt.Printf("Enter your height :")
    fmt.Scan(&height)
	
	BMI := weight/(height*height)

	fmt.Printf("Your name is %s and your BMI is %f",name,BMI)

}