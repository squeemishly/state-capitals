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
	statePtr := flag.String("state", "unk", "a state name")
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
