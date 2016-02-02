package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    )

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run http.go <website>")
		os.Exit(1)
	}
    
    if response, err := http.Get(os.Args[1]); err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        if contents, err := ioutil.ReadAll(response.Body); err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        } else {
            fmt.Printf("%s\n", string(contents))
        }
    }
}
