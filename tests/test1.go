package main	


import (
	"fmt"
	"time"
	"math/rand"
	"lib"
)

var share_limit int = 12
var k_limit int = 6

func main () {

	n := 500
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Gnerating random secrets...")
	secrets := make([]int, n)
	shares := make([]int, n)
	threshold := make([]int, n)

	for i := 0; i < n; i++ {
		secrets[i] = rand.Intn(10000)
		shares[i] = rand.Intn(10000) % share_limit + 10   
		threshold[i] = rand.Intn(10000) % k_limit + 2    
		fmt.Printf("S=%d, n=%d, k=%d\n", secrets[i], shares[i], threshold[i])
		co := make([]int, 0)
		x_input := make([]int, 0)
		points_test := map[int] int {}
		for c := 0; c < threshold[i] - 1; c++ {
			co = append(co, rand.Intn(1000))
		}
		fmt.Println("Coefficients:", co)

		for j := 1; j <= shares[i]; j++ {
			x := rand.Intn(1000)
			x_input = append(x_input, x)
			y := lib.Calc_single_point(x, co, threshold[i] - 1, secrets[i])
			points_test[x] = y
		}
		fmt.Println(shares[i], "random points:", points_test)
		fmt.Println("x-cooridnates:",x_input)
		fmt.Println("selecting a", threshold[i], "length subset:", x_input[0:threshold[i]])

		result := lib.Find_secret(x_input[0:threshold[i]], points_test)

		if result == secrets[i] {
			fmt.Println("Successfully recovered secret")
		} else {
			fmt.Println("Couldn't recover secret")
		}
		//return
	}
}