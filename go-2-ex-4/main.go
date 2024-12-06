package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Students []Student
}

func main() {
	classes := map[string]Class{
		"Class A": {
			Students: []Student{
				{"Alice", "Smith"},
				{"Bob", "Brown"},
				{"Charlie", "Davis"},
			},
		},
		"Class B": {
			Students: []Student{
				{"David", "Evans"},
				{"Eve", "Foster"},
				{"Frank", "Green"},
			},
		},
	}

	modules := map[int][]string{
		104: {"Class A", "Class B"},
		117: {"Class A"},
		346: {"Class B"},
	}

	fmt.Println("Classes:")
	for className, class := range classes {
		fmt.Println(" ", className, ":")
		for _, student := range class.Students {
			fmt.Println("   -", student.FirstName, student.LastName)
		}
	}

	fmt.Println("Modules:")
	for moduleID, attendingClasses := range modules {
		fmt.Println(" Module", moduleID, "attended by:")
		for _, className := range attendingClasses {
			fmt.Println("   -", className)
		}
	}
}
