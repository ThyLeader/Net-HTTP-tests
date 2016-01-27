package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    )
    
func main() {
    if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go (<all>, <ip>, <trace>, or <epoch>)")
		os.Exit(1)
	}
    switch os.Args[1] {
        default:
            fmt.Println("Command not recognized. Please try again. Use lowercase")
            fmt.Println("Usage: go run main.go (<all>, <ip>, <trace>, or <epoch>)")
        case "ip":
            fmt.Println("Your IP is:")
            GetIP()
        case "trace":
            fmt.Println("Your Traceroute is:")
            GetTrace()
        case "epoch":
            fmt.Println("The current Unix Time/Epoch is:")
            GetEpoch()  
        case "all":
            fmt.Println("Your IP is:")
            GetIP()
            fmt.Println("Your Traceroute is:")
            GetTrace()
            fmt.Println("The current Unix Time/Epoch is:")
            GetEpoch()
    }
}

func GetIP() {
    if response, err := http.Get("http://icanhazip.com"); err != nil {
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

func GetTrace() {
    if response, err := http.Get("http://icanhaztrace.com"); err != nil {
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

func GetEpoch() {
    if response, err := http.Get("http://icanhazepoch.com"); err != nil {
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