// compare efficience between my own version and the one that uses strings.Join

// for time, refer to 1.6
// for systematic performance evaluation, refer to 11.4 on writing benchmark tests

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	for i:= 0; i<= 10000000; i++ {
		MyStringsJoin(os.Args, ", ")
	}
	fmt.Println("run my own strings join takes: %.2fs seconds.", time.Since(start).Seconds())
	start = time.Now()
	for i:= 0; i<= 10000000; i++ {
		MyStringsJoin2(os.Args, ", ")
	}
	fmt.Println("run my own strings join takes: %.2fs seconds.", time.Since(start).Seconds())
}

func MyStringsJoin(Args []string, Sep string) string {
	s := ""
	for _, Arg := range Args {
		s += Arg + Sep
	}
	return s
}

func MyStringsJoin2(Args []string, Sep string) string {
	return strings.Join(Args, Sep)
}
