//Euclidean Algorithm Interactive Example using Go
//btylrob

package main

import "fmt"

// Our function euclidGCD, "Euclidean Algorithm" is going to to take 2 integers and give us one integer back.
func euclidGCD(a, b int) int {
	if b == 0 { //According to Euclidean Algorithm if a = 0 then be is the GCD and if b = 0, a is the GCD.
		return a
	}
	//Once the rucursion stops the program will return the last non-zero
	return euclidGCD(b, a%b)
}

func main() {
	a, b := 60, 100 //Example integers 60 and 100
	gcd := euclidGCD(a, b)
	fmt.Printf("GCD of %d and %d is %d\n", a, b, gcd)
}
