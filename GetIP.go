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

    incorrectWords := make([]string, 0, len(os.Args)-1)
    argArray := make([]string, len(os.Args)-1)
    choice := "2"

    for i := 1; i < len(os.Args); i++ {
        argArray[i-1] = os.Args[i]
    }
    
    for _, arg := range argArray {
        switch arg {
            default:
                fmt.Println(arg, "is spelled incorrectly. Adding to array")
                incorrectWords = append(incorrectWords, arg)
                //index++
            case
                "all",
                "ip",
                "trace",
                "epoch":
                fmt.Println(arg, "is spelled correctly")
        }
    }
    fmt.Println(incorrectWords)
    
    incorrect := len(incorrectWords)
    fmt.Println("There are", incorrect, "incorrect words")
    for _, word := range incorrectWords {
        fmt.Println("\nThe word", word, "was incorrect. Autocorrect commencing...")
        fmt.Println("Did you mean:", model.SpellCheck(word), "?")
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