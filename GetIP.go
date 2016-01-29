package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "log"
    
    "github.com/sajari/fuzzy"
    )
    
func main() {

    model := fuzzy.NewModel()
    model.SetDepth(5)
    words := []string{"all", "ip", "trace", "epoch"}
    model.Train(words)

    commands := []string{"all", "ip", "trace", "epoch"}
    incorrectWords := make([]string, len(os.Args)-1)
    matches := 0
    argArray := make([]string, len(os.Args)-1)
    choice := "2"

    for i := 1; i < len(os.Args); i++ {
        argArray[i-1] = os.Args[i]
    }
    
    for i, choice := range argArray {
		for _, command := range commands {
            if choice == command {
                matches++
            } else {
                incorrectWords[i] = choice
            }
        }
	}
    
    if matches := len(os.Args) {
        incorrect := len(os.Args) - matches
        fmt.Println("There are", incorrect, "incorrect words")
        for _, word := range incorrectWords {
            fmt.Println("\nThe word", word, "was incorrect. Autocorrect commencing...")
            fmt.Println("Did you mean:", model.Suggestions(saword, false), "?")
            fmt.Println("Press 1 for yes, anything else for no.")
            if _, err := fmt.Scanln(&choice); err != nil {
                log.Fatal(err)
            }
            if choice == "1" {
                fmt.Println("Replacing word...")
            } else {
                fmt.Println("Exiting...")
                os.Exit(1)
            }
        }
    }
    
    for _, choice := range argArray {
        switch choice {
            default:
                fmt.Println("Command not recognized. Please try again. Use lowercase")
                fmt.Println("Usage: go run main.go (<all>, <ip>, <trace>, or <epoch>)")
                fmt.Println("The problem choice is:", choice)
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