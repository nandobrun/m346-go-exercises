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
			FirstName: "Nando",
			LastName:  "Brun",
		},
		BirthDate: BirthDate{
			Day:   11,
			Month: 1,
			Year:  2006,
		},
		NumberOfSiblings: 2,
		ZodiacSign:       'A',
	}

	fmt.Println(me)
	fmt.Println("Siblings Before:", me.NumberOfSiblings)

	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
