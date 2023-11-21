package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var colorPtr string

var validColors = map[string]struct{}{
	"rood":   {},
	"blauw":  {},
	"groen":  {},
	"geel":   {},
	"oranje": {},
	"roze":   {},
	"paars":  {},
}

//Her wordt een map aangemaakt validColors die geldige kleuren bevat als sleutels.

// Wanneer een kleur wordt opgegeven via de vlag -kleur in de init functie, wordt de opgegeven kleur gecontrolleerd tegen deze map van geldige kleuren.
func init() {
	// flag voor de kleur
	flag.StringVar(&colorPtr, "kleur", "", "Kleur (rood, blauw, groen, geel, oranje, roze, paars).")
	flag.Parse()
	colorPtr = strings.ToLower(colorPtr)
	//Hier wordt colorPtr omgezet naar kleine letters (lettergevoeligheid) zodat het niet uitmaakt of je kleine of grote letters gebruikt.

	if _, isValid := validColors[colorPtr]; !isValid {
		fmt.Println("Geen geldige kleur. Kies uit rood, blauw, groen, geel, oranje, roze, paars.")
		os.Exit(1) // Stop het programma bij een ongeldige kleur
	}
}

//Hier wordt voordat de main functie aangeroepen wordt gecontrolleerd of de kleur voorkomt in de lijst die al aangemaakt is. Komt de kleur niet vol dan wordt het programma afgesloten.

func main() {

	// Geef kleuren aan voor verschillende gedichten
	// map[string]string{ dit is de aanduiding van een map. Het geeft aan dat gedichten een map is waarvan de sleutels en waarden beide string zijn.
	gedichten := map[string]string{
		"rood":   "Rood met passie.",
		"blauw":  "Blauw zoals de lucht.",
		"groen":  "Groen van de natuur.",
		"geel":   "Geel als de stralen van de zon.",
		"oranje": "Oranje de kleur van de zon",
		"roze":   "Roze is de kleur van een verse vrucht",
		"paars":  "Paars is de kleur van luxe.",
	}

	// Selecteer het gedicht op basis van de opgegeven kleur
	gekozenKleur, exists := gedichten[colorPtr]
	if !exists {
		fmt.Println("Geen geldige kleur. Kies uit rood, blauw, groen, geel, oranje, roze, paars.")
		return
	}

	// Open een text bestand om het gedicht naartoe te zenden
	fileName := "gedicht.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Fout bij het maken van het tekstbestand:", err)
		return
	}
	defer file.Close()

	// Schrijf het gedicht naar het text file
	_, err = file.WriteString(gekozenKleur)
	if err != nil {
		fmt.Println("Fout bij het schrijven naar het tekstbestand:", err)
		return
	}

	fmt.Printf("Gedicht is opgeslagen in %s op basis van jouw opgegeven kleur %s. \n", fileName, colorPtr)
}
