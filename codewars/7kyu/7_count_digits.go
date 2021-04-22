// https://www.codewars.com/kata/566fc12495810954b1000030

package main

import (
	"strconv"
	"strings"
)

func NbDig(n int, d int) (out int) {
	for i := 0; i <= n; i++ {
		out += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}

	return
}
