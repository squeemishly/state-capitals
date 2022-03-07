package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("gimme an state!")
		os.Exit(1)
	}

	state := os.Args[1]

	abbrev := stateToAbbrev(state)
	fmt.Printf("The abbreviation for Oregon is %s", abbrev)
}

func stateToAbbrev(state string) string {
	stateAbbrev := map[string]string{
		"Oregon":   "OR",
		"Colorado": "CO",
	}

	return stateAbbrev[state]
}
