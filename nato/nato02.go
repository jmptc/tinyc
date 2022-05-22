package main

import (
    "os"
    "log"
    "fmt"
    "io"
    "bufio"
    "unicode"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Provide a filename argument")
    }

    f, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    
    nato := []string { "Alpha", "Bravo", "Charile", "Delta", "Echo", "Foxtrot",
        "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November",
        "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor",
        "Whiskey", "Xray", "Yankee", "Zulu" }

    reader := bufio.NewReader(f)
    for {
        r, _, err := reader.ReadRune()
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }

        if unicode.IsLetter(r) {
            r = unicode.ToLower(r)
            idx := int(r - 'a')
            if idx < len(nato) && idx >= 0 {
                fmt.Print(nato[idx])
            }
        } else {
            fmt.Printf("%c", r)
        }
 
    }
}
