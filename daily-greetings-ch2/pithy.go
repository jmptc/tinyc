package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"math/rand"
	"time"
)

func main() {
	quotes := make([]string, 0, 10)

	file, err := os.Open("quotes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		quotes = append(quotes, scanner.Text())
	}
	
	seed := rand.NewSource(time.Now().UnixNano())
	ran := rand.New(seed)

	index := ran.Intn(len(quotes))
	fmt.Println(index, quotes[index])
}
