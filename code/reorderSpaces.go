package main

import "strings"

func reorderSpaces(s string) string {
	words := strings.Fields(s)
	space := strings.Count(s," ")
	lw := len(words) - 1
	if lw == 0{
		return words[0] + strings.Repeat(" ",space)
	}
	return strings.Join(words,strings.Repeat(" ",space/lw)) + strings.Repeat(" ",space%lw)
}

func main() {
	reorderSpaces("  this   is  a sentence ")
}
