package lib

var p int = 108347
func Exp(base, exp int) int {

	result := 1

	for i := 1; i <= exp; i++ {
		result *= base
	}

	return result
}


func Mod_inverse(a, p int) int {

	if a < 0 {
		a = ((a % p) + p) % p
	}

	for i := 1; i < p; i++ {
		if ((a % p) * (i % p)) % p == 1 {
			return i
		}
	}

	return 0
}

func Find_secret(x_input[]int, points map[int]int) int{

	reconstructed := 0

	for _, k := range x_input { 
		y := points[k]
		pr_x := 1
		for _, l := range x_input {
			if l != k {
				inv := Mod_inverse(l - k, p)

				pr_x *= inv % p * l % p
			}
		}
		y *= pr_x % p
		reconstructed += y
	}

	return reconstructed % p
}


func Calc_single_point(x int, co []int, l int, secret int) int{

	y := 0
	for k := 0; k < l; k++ {
		y += co[k] * Exp(x, k + 1) 
	}
	
	return (y + secret) % p 
}

func In_slice(needle int, haystack[] int) bool {

	n := len(haystack)

	for i := 0; i < n; i++ {

		if needle == haystack[i] {
			return true
		}
	}
	return false
}