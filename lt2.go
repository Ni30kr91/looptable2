package main 

import "fmt"

func main() { 

	fmt.Println("Enter any number")
	var input int
	fmt.Scanln(&input)
	printtable(input)
}

func printtable(input int)  {
	var i = 0
	count:= input*10
	j := 1

	for i = input; i <= count; i=i+input {
	fmt.Println(input, "*", j, "=", i)
	j += 1
	
	} 
} 