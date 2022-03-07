package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type State struct {
	abbreviation string
	capital      string
}

var states = map[string]State{
	"oregon": {
		abbreviation: "OR",
		capital:      "Salem",
	},
	"colorado": {
		abbreviation: "CO",
		capital:      "Denver",
	},
	"new mexico": {
		abbreviation: "NM",
		capital:      "Santa Fe",
	},
}

func main() {
	statePtr := flag.String("state", "", "a state name")
	capitalPtr := flag.String("capital", "", "a state capital")
	flag.Parse()

	if *statePtr != "" {
		fmt.Printf("state: %v", *statePtr)
		fmt.Println()

		state, ok := states[strings.ToLower(*statePtr)]

		if ok {
			correctedState := strings.Title(*statePtr)

			fmt.Printf("The abbreviation for %s is %s", correctedState, state.abbreviation)
			fmt.Println()
			fmt.Printf("The capital for %s is %s", correctedState, state.capital)
		} else {
			fmt.Println("We don't know that state. Try again.")
			os.Exit(1)
		}
	}

	if *capitalPtr != "" {
		fmt.Printf("capital: %v", *capitalPtr)
		fmt.Println()

		capitals := makeCapitals()
		state, ok := capitals[strings.ToLower(*capitalPtr)]

		if ok {
			correctedState := strings.Title(state)
			correctedCapital := strings.Title(*capitalPtr)
			fmt.Printf("The state where %s is the capital is %s", correctedCapital, correctedState)
		} else {
			fmt.Println("We don't know that capital. Try again.")
			os.Exit(1)
		}
	}
}

func makeCapitals() map[string]string {
	myCapitals := make(map[string]string)
	for i, s := range states {
		myCapitals[strings.ToLower(s.capital)] = i
	}
	return myCapitals
}
