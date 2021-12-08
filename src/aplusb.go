package main

/**
 * @param a: An integer
 * @param b: An integer
 * @return: The sum of a and b
 */
func aplusb(a int, b int) int {
	// write your code here
	return a + b
}

func calculate(r int) []float64 {
	const PI = 3.14
	r2 := float64(r)
	var result = []float64{2 * PI * float64(r2), PI * r2 * r2}
	return result
}
