// Why the program hang at line 15?
package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        return
    })

    err := http.ListenAndServe(":8888", nil)
    if err != nil {
        log.Println("start server error")
    }

    log.Println("in main")

    defer log.Println("in main again")
    for i := 0; i < 10000; i++ {
        go func() {
            log.Printf("%v: in goroutine", i)
        }()
        if i == 1000 {
            os.Exit(1)
        }
    }
}
