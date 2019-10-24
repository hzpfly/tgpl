// compare efficience between my own version and the one that uses strings.Join

// for time, refer to 1.6
// for systematic performance evaluation, refer to 11.4 on writing benchmark tests

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(MyStringsJoin(os.Args, ", "))
}

func MyStringsJoin(Args []string, Sep string) string {
	s := ""
	for _, Arg := range Args {
		s += Arg + Sep
	}
	return s
}
