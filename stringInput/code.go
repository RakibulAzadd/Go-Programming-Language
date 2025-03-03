package main

import (
	"bufio"
	"fmt"
	"os"
)

// Scan, Scanf, bufio
func main(){
	fmt.Print("Enter your full name and age: ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println(input)


}