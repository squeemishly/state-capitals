package main

import (
	"flag"
	"fmt"
	"os"
)

type State struct {
	name         string
	abbreviation string
	capital      string
}

var states = []State{
	{
		name:         "Oregon",
		abbreviation: "OR",
		capital:      "Portland",
	},
	{
		name:         "Colorado",
		abbreviation: "CO",
		capital:      "Denver",
	},
}

func main() {
	statePtr := flag.String("state", "", "a state name")
	capitalPtr := flag.String("capital", "", "a state capital")
	flag.Parse()

	if *statePtr != "" {
		fmt.Printf("state: %v", *statePtr)
		fmt.Println()

		state := findState(*statePtr)

		if state.abbreviation == "" {
			fmt.Println("We don't know that state. Try again.")
			os.Exit(1)
		}

		fmt.Printf("The abbreviation for %s is %s", *statePtr, state.abbreviation)
		fmt.Println()
		fmt.Printf("The capital for %s is %s", *statePtr, state.capital)
	}

	if *capitalPtr != "" {
		fmt.Printf("capital: %v", *capitalPtr)
		fmt.Println()

		state := findStateFromCapital(*capitalPtr)

		if state.abbreviation == "" {
			fmt.Println("We don't know that state. Try again.")
			os.Exit(1)
		}

		fmt.Printf("The state where %s is the capital is %s", *capitalPtr, state.name)
	}
}

func findState(state string) State {
	var myState State
	for _, s := range states {
		if s.name == state {
			myState = s
			break
		}
	}

	return myState
}

func findStateFromCapital(capital string) State {
	var myState State
	for _, s := range states {
		if s.capital == capital {
			myState = s
			break
		}
	}

	return myState
}
