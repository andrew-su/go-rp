package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	http.HandleFunc("/", printRequest)

	listenAddr := fmt.Sprintf("%s:%d", "", *port)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		panic(err)
	}
}

func printRequest(w http.ResponseWriter, r *http.Request) {
	var output []string // Add the request string
	output = append(output, fmt.Sprintf("%s %s %s", r.Method, r.URL, r.Proto))
	output = append(output, fmt.Sprintf("Host: %v", r.Host))
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			output = append(output, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			output = append(output, fmt.Sprintf("\nError: %s", err.Error()))
		} else {
			output = append(output, "\n")
			output = append(output, r.Form.Encode())
		}
	}

	fmt.Fprintf(w, strings.Join(output, "\n"))
}
