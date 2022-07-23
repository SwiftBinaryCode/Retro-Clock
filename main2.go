package main

import (
	"os"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// Render colorless output to a file
	f, _ := os.Create("out.txt")
	defer f.Close()

	nums := []int{1, 3, 5, 2, 4, 8}

	s.Writer = f
	s.Colors(false)
	s.Show("nums", nums)
}
