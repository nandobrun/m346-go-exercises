package main

import "fmt"

func main() {
	var firstName string = "Nando"
    var lastName string = "Brun"
    var dayOfBirth int = 11
    var monthOfBirth int = 01
    var yearOfBirth int = 2006
    var numberOfSiblings int = 1
    var heightInMeters float64 = 1
	var zodiacSign string = "Steinbock"

	
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
