package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	if t.Hour() < 12 {
		fmt.Printf("Good morning\n");
	} else if t.Hour() < 17 {
		fmt.Printf("Good afternoon\n");
	} else {
		fmt.Printf("Good evening\n");
	}

	fmt.Printf("The time is %s\n", t.Local())
}
