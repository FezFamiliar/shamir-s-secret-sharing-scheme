package main

import (
	"fmt"
	"time"
	"strconv"
	"math/rand"
	"github.com/pborman/getopt"
	"lib"
)

var exponent = map[int] string {
	0: "\xe2\x81\xb0",
	1: "x",
	2: "x\xc2\xb2",
	3: "x\xc2\xb3",
	4: "x\xe2\x81\xb4",
	5: "x\xe2\x81\xb5",
	6: "x\xe2\x81\xb6",
	7: "x\xe2\x81\xb7",
	8: "x\xe2\x81\xb8",
	9: "x\xe2\x81\xb9",
}


func main () {

	rand.Seed(time.Now().UTC().UnixNano())
	secret_ptr := getopt.Int32Long("secret", 's', 0, "Your super secret integer which you want to distribute")
	shares_ptr := getopt.Int32Long("shares", 'n', 0, "The number of shares | The secret will be divided this many times")
    threshold_ptr := getopt.Int32Long("threshold", 'k', 0, "Threshold | The minimum number of shares required to reconstruct the secret")
    
    getopt.Parse()

    secret := int(*secret_ptr)
	n := int(*shares_ptr)
	k := int(*threshold_ptr)


    if secret == 0 || n == 0 || k == 0{
    	getopt.Usage()
    	return
    } else if k > n {
    	fmt.Println("Number of shares must be greater than the minimum shares required to reconstruct the secret!")
    	return
    }



	fmt.Println("The secret is:", secret)
	fmt.Println("Dividing into", n, "number of shares...")
	fmt.Println("Setting threshold value to", k, "points...")
	split_secret(secret, n, k)
}


func split_secret(secret, shares, threshold int) {
	
	
	coefficients := make([]int, 0)
	var polynomial string = "f(x) = " + strconv.Itoa(secret)
	fmt.Println("Initializing polynomial...")

	for i := 0; i < threshold - 1; i++ {

		coefficients = append(coefficients, rand.Intn(1000))
		polynomial += " + " + strconv.Itoa(coefficients[i]) + exponent[i + 1]
	}

	fmt.Println("Coefficients are", coefficients)
	fmt.Println("Done!")
	fmt.Println("Polynomial is:", polynomial)
	fmt.Println("Generating", shares, "points...")

	points := make([]int, 0)
	points_map := map[int] int {}
	for i := 1; i <= shares; i++ {
		x := rand.Intn(1000)
		y := lib.Calc_single_point(x, coefficients, threshold - 1, secret)
		points = append(points, y)
		points_map[x] = y
	}
	fmt.Println(points)
	fmt.Println("Done!")
	fmt.Println("The", shares, "random points are:")

	var c int = 1
	for x, y := range points_map {
		fmt.Printf("%d. (%d, %d)\n", c, x, y)
		c++
	}

	var flag string
	fmt.Println("Do you wish to recover the secret? [y/n]")
	fmt.Scanln(&flag)
    if flag == "y" {
		recovered_secret := recover(points_map, threshold, shares)
		fmt.Println("Recovering...")
		fmt.Println("Recovered secret is:", recovered_secret)
    } else {
    	fmt.Println("Exiting...")
    }

}


func recover(points map[int]int, threshold int, shares int) int {

	fmt.Println("Please enter", threshold, "points in order to recover the secret (only the x-coordinate is needed)")


	x_input := make([]int, threshold)

	for k := range x_input {
		fmt.Scanln(&x_input[k])

		if points[x_input[k]] == 0 {
			fmt.Printf("The point (%d, f(%d)) is not valid!\n", x_input[k], x_input[k])
		} else {
			if k == threshold - 2 {
				fmt.Println("1 more..")
			} else if k != threshold - 1{
				fmt.Println("Go on, don't be shy..")
			}
			
		}
	}

	fmt.Println(x_input)
	return lib.Find_secret(x_input, points)
	
}