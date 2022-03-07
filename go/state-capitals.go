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
	if len(os.Args) == 1 {
		fmt.Println("gimme an state!")
		os.Exit(1)
	}

	state := os.Args[1]
	myState := findState(state)

	if myState.abbreviation == "" {
		fmt.Println("We don't know that state. Try again.")
		os.Exit(1)
	}

	fmt.Printf("The abbreviation for %s is %s", state, myState.abbreviation)
	fmt.Println()
	fmt.Printf("The capital for %s is %s", state, myState.capital)
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
