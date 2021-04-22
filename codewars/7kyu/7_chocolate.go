// https://www.codewars.com/kata/534ea96ebb17181947000ada

func BreakChocolate(n, m int) int {
	if n < 1 || m < 1 {
		return 0
	}

	return (n * m) - 1
}