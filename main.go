package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("\n")

	print(lable())
	print(promptSymbol())
}

func lable() (string, style) {
	if os.Getenv("IN_NIX_SHELL") != "" {
		return "󱄅", style{fg: "#7eb7e2"}
	}

	_, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return "ඞ", style{fg: "#b01000"}
}

func promptSymbol() (string, style) {
	return "  ", style{fg: "#fab387"}
}
