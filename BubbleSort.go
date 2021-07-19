package main

import (
	"fmt"
)

func main() {
	//create array of unsorted elements and make a slice of it
	A := [6]int{5, 2, 7, 1, 4, 0}
	var S []int = A[:]
	fmt.Println("Unsorted: ", S)
	//setup variables for use later
	n := len(S)
	flag := true
	
	for flag {
		flag = false
		
		//loop through each element of slice
		for i := 1; i < n; i++{
			//if previous is smaller then swap
			if S[i-1] > S[i]{
				temp := S[i-1]
				S[i-1] = S[i]
				S[i] = temp
				flag = true
			}
		}
		//remove unnecisary steps
		n = n-1
	}
	//print sorted slice
	fmt.Println("Sorted: ", S)
}
