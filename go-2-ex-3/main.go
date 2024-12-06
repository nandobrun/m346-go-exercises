package main

import "fmt"

func main() {
	modules := map[int]string{
		104: "Einführung in die Informatik",
		117: "Datenbanken erstellen und verwalten",
		346: "Cloud-Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 117)
	modules[200] = "Algorithmen und Datenstrukturen"
	modules[346] = "Erweiterte Cloud-Lösungen"

	fmt.Println(modules)
}
