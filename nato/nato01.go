package main

import (
    "bufio"
    "fmt"
    "os"
    "unicode"
)

func main() {
    nato := []string { "Alpha", "Bravo", "Charile", "Delta", "Echo", "Foxtrot",
        "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November",
        "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor",
        "Whiskey", "Xray", "Yankee", "Zulu" }

    fmt.Print("Enter a word or a phrase: ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()

    for _, c := range text {
        if unicode.IsLetter(c) {
            c = unicode.ToLower(c)
            idx := int(c - 'a')
            if idx < len(nato) && idx >= 0 {
                fmt.Print(nato[idx])
            }
        } else {
            fmt.Printf("%c", c)
        }
    }
}
