package csvier_test

import (
	"fmt"
	"testing"

	"github.com/fabritsius/csvier"
)

func ExampleRead() {
	data, err := csvier.Read("./test_data/data.csv")
	if err != nil {
		panic(err)
	}

	for _, r := range data {
		fmt.Printf("%s fights with %s.\n", r["NAME"], r["WEAPON"])
	}
	// Output:
	// Peter fights with Quad Blasters.
	// Gamora fights with Godslayer.
	// Rocket fights with Ion Cannon.
	// Groot fights with courage.
	// Drax fights with Dual Knives.
	// Nebula fights with Electroshock Batons.
}

func ExampleIndex() {
	data, err := csvier.Read("./test_data/data.csv",
		csvier.Index([]string{"id", "name", "race", "weapon"}),
		csvier.Skip(1),
		csvier.Limit(5),
	)
	if err != nil {
		panic(err)
	}

	for _, r := range data {
		fmt.Printf("%s fights with %s.\n", r["name"], r["weapon"])
	}
	// Output:
	// Peter fights with Quad Blasters.
	// Gamora fights with Godslayer.
	// Rocket fights with Ion Cannon.
	// Groot fights with courage.
	// Drax fights with Dual Knives.
}

func ExampleDelimiter() {
	data, err := csvier.Read("./test_data/data.tsv",
		csvier.Limit(5),
		csvier.Delimiter('\t'),
	)
	if err != nil {
		panic(err)
	}

	for _, r := range data {
		fmt.Printf("%s fights with %s\n", r["NAME"], r["WEAPON"])
	}
	// Output:
	// Peter fights with Quad Blasters
	// Gamora fights with Godslayer
	// Rocket fights with Ion Cannon
	// Groot fights with courage
	// Drax fights with Dual Knives
}

func TestDelimiterError(t *testing.T) {
	for _, d := range []rune{'\r', '\n', 0xFFFD} {
		_, err := csvier.Read("./test_data/data.csv", csvier.Delimiter(d))
		if err == nil {
			t.Error("csvier: wrong delimiter values aren`t handled")
		}
	}
}
