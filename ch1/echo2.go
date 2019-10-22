package main

import (
	"strconv"
	"fmt"
	"os"
)

func main() {
	s, sep := "", "\n"
	for i, arg := range os.Args[0:] {
		s += strconv.Itoa(i) + " " + arg + sep
	}
	fmt.Println(s)
}
