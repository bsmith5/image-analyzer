package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Application starting at", time.Now().Format(time.RFC3339))
}
