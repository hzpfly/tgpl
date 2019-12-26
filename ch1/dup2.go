// Find duplicate line
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		input := bufio.NewScanner(os.Stdin)
		CountLines(input, "stdin")
	} else {
		for _, file := range files {
			scan, _ := os.Open(file)
			input := bufio.NewScanner(scan)
			CountLines(input, file)
			scan.Close() //I forgot this line.
		}
	}
}

func CountLines(input *bufio.Scanner, f string) {
	counts := make(map[string]int)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, f, line)
		}
	}
}
