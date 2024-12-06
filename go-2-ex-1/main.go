package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   int
	Month int
	Year  int
}

type Profile struct {
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	me := Profile{
		FullName: FullName{
			FirstName: "Max",
			LastName:  "Mustermann",
		},
		BirthDate: BirthDate{
			Day:   1,
			Month: 1,
			Year:  1990,
		},
		NumberOfSiblings: 2,
		ZodiacSign:       'A',
	}

	fmt.Println(me)
	fmt.Println("Siblings Before:", me.NumberOfSiblings)

	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}