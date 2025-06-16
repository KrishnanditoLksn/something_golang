package main

import (
	"fmt"
	"os"
)

func osargs() {
	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[0]
		fmt.Println(i)
		sep = " "
	}
	fmt.Println(s)
}

func main() {
	osargs()
}
