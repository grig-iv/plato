package main

import (
	"fmt"
	"os"
	"strings"
)

func init() {
	// for debugg
	if len(os.Args) == 2 {
		workingDir = replaceTilde(os.Args[1])
	}

	if strings.HasSuffix(workingDir, "/") == false {
		workingDir = workingDir + "/"
	}
}

func main() {
	fmt.Print("\n")

	print(lable())

	fmt.Print("  ")

	print(promptSymbol())
}

func lable() (string, style) {
	if os.Getenv("IN_NIX_SHELL") != "" {
		return "󱄅", style{fg: "#7eb7e2"}
	}

	if match("~/.config/*") {
		return "", style{}
	}

	if match("~/*") {
		return "", style{}
	}

	if match("/") {
		return "/", style{}
	}

	return "", style{}
}

func promptSymbol() (string, style) {
	return " ", style{fg: "#fab387"}
}
