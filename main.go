package main

import (
	"fmt"
	li "lemin/funcs"
	"os"
)

// Where it is all coming together.

func main() {
	var graph li.Graph
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Print("\nERROR: Invalid number of arguments\n\nUSAGE: go run main.go [INPUT FILENAME]\n")
		return
	}
	ants, rooms, errs := graph.ParseInput(args[0])
	if errs != nil {
		for _, err := range errs {
			fmt.Println(err)
		}
		return
	}
	paths := graph.FindPaths(rooms)
	
	li.MoveAnts(ants, paths)
}
