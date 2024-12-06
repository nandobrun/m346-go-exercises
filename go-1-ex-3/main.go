package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed für Zufallszahlen

	var eyes = rand.Intn(6) + 1
	var when = time.Now()

	// Ausgabe mit fmt.Fprintln direkt in os.Stdout und os.Stderr
	fmt.Fprintln(os.Stdout, "The dice shows", eyes, "eyes")
	fmt.Fprintln(os.Stderr, "The dice was rolled at", when)
}
