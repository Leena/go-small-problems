// https://www.codewars.com/kata/56541980fa08ab47a0000040/train/go

package main

import "fmt"

var validChars = [14]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"}

func PrinterError(s string) string {
	demonimator := len(s)
	numerator := 0
	for _, char := range s {
		if !contains(validChars, string(char)) {
			numerator++
		}
	}
	return fmt.Sprint(numerator, "/", demonimator)

}

func contains(arr [14]string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func main() {

	fmt.Println(PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz") == "3/56")
	fmt.Println(PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz") == "6/60")
	fmt.Println(PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu") == "11/65")

}
