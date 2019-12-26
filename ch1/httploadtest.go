package main

import (
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var myClient *http.Client

func startWebserver() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//time.Sleep(time.Millisecond * 50)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	go http.ListenAndServe(":8080", nil)
}

func startLoadTest() {
	count := 0
	for {
		resp, err := myClient.Get("http://localhost:8080/")
		if err != nil {
			panic(fmt.Sprintf("Got error: %v", err))
		}
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		log.Printf("Finished GET request #%v", count)
		count += 1
	}
}

func main() {
	defaultRoundTripper := http.DefaultTransport
	defaultTransportPointer, ok := defaultRoundTripper.(*http.Transport)
	if !ok {
		panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport"))
	}
	defaultTransport := *defaultTransportPointer
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100

	myClient = &http.Client{Transport: &defaultTransport}
	startWebserver()
	for i := 0; i < 100; i++ {
		go startLoadTest()
	}
	time.Sleep(time.Second * 3600)
}
