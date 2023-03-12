package main

import (
	"fmt"
	"os"

	"github.com/jmptc/glox/scanner"
)

func main() {
	fmt.Println(os.Args)

	scanner := scanner.NewScanner("(){}*/+-")
	toks := scanner.ScanTokens()
	for i, t := range toks {
		fmt.Println(i, t)
	}
}
