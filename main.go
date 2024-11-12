package main

import (
	"fmt"
	"os"
)

var (
	promptSymbolStyle = style{fg: "#FAB387"}
	nixIconStyle      = style{fg: "#7eb7e2"}
)

func main() {
	fmt.Print("\n")
	printLable()
	print("  ", promptSymbolStyle)
}

func printLable() {
	if os.Getenv("IN_NIX_SHELL") != "" {
		print("󱄅", nixIconStyle)
		return
	}

	fmt.Printf("")
}
