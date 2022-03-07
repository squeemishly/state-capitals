package main

import (
	"fmt"
	"os"
)

type State struct {
	name         string
	abbreviation string
	capital      string
}

var states = []State{
	State{
		name:         "Oregon",
		abbreviation: "OR",
		capital:      "Portland",
	},
	State{
		name:         "Colorado",
		abbreviation: "CO",
		capital:      "Denver",
	},
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("gimme an state!")
		os.Exit(1)
	}

	state := os.Args[1]

	abbrev := stateToAbbrev(state)
	if abbrev == "" {
		fmt.Println("We don't know that state. Try again.")
		os.Exit(1)
	} else {
		fmt.Printf("The abbreviation for %s is %s", state, abbrev)
	}

	fmt.Println()

	capital := stateCapital(state)
	if capital == "" {
		fmt.Println("We don't know that state. Try again.")
		os.Exit(1)
	} else {
		fmt.Printf("The capital for %s is %s", state, capital)
	}
}

func stateToAbbrev(state string) string {
	var abbrev string
	for _, s := range states {
		if s.name == state {
			abbrev = s.abbreviation
			break
		}
	}

	return abbrev
}

func stateCapital(state string) string {
	var capital string
	for _, s := range states {
		if s.name == state {
			capital = s.capital
			break
		}
	}

	return capital
}
