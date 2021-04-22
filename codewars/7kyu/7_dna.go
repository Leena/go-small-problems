// https://www.codewars.com/kata/554e4a2f232cdd87d9000038

package main

import "fmt"

var DNAmap = map[string]string{
	"A": "T",
	"T": "A",
	"C": "G",
	"G": "C",
}

func DNAStrand(dna string) (compliment string) {
	for _, char := range dna {
		compliment = compliment + DNAmap[string(char)]
	}
	return compliment
}

func main() {
	fmt.Println(DNAStrand("AAAA") == "TTTT")
	fmt.Println(DNAStrand("ATTGC") == "TAACG")
	fmt.Println(DNAStrand("GTAT") == "CATA")
}

/* Alternative Solutions
 import "strings"

func DNAStrand(dna string) string {
  replacer := strings.NewReplacer("A", "T", "T", "A", "G", "C", "C", "G")
  return(replacer.Replace(dna))
}

func DNAStrand(dna string) string {
  var complements = map[string]string{"A":"T", "C":"G", "G":"C", "T":"A"}
  var and string
  for _, char := range dna {
    and += complements[string(char)]
  }
  return and
}
*/
