package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type lableMatch struct {
	pattern string
	lable   string
	style   style
}

var (
	homeDir       = os.Getenv("HOME") + "/"
	workingDir, _ = os.Getwd()
	now           = time.Now()
)

var dirLabels = []lableMatch{
	{lable: "󰒓", pattern: "~/.config/*"},
	{lable: "󰧑", pattern: "~/mind/*"},
	{lable: "󰇄", pattern: "~/desktop/*"},
	{lable: "", pattern: "~/cloud/*"},
	{lable: "", pattern: "~/sources/*"},
	{lable: "", pattern: "~/"},

	{lable: "󱃩", pattern: "*/downloads/*"},
	{lable: "󱉟", pattern: "*/books/*"},

	{lable: "󱄅", pattern: "/etc/nixos/*"},
	{lable: "󰕓", pattern: "/run/media/*"},
	{lable: "", pattern: "/mnt/c/*"},
	{lable: "󰋊", pattern: "/mnt/*"},
	{lable: "/", pattern: "/"},
}

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
	fmt.Println()

	if isBedTime() {
		print(now.Format("󰤄 15:04"), style{fg: "#f38ba8"})
		fmt.Println()
	}

	print(lable())

	fmt.Print("  ")

	print(promptSymbol())

	fmt.Print(" ")
}

func isBedTime() bool {
	return now.Hour() > 23 || now.Hour() < 7
}

func lable() (string, style) {
	if os.Getenv("IN_NIX_SHELL") != "" {
		return "󱄅", style{fg: "#7eb7e2"}
	}

	for _, m := range dirLabels {
		if match(workingDir, m.pattern) {
			return m.lable, m.style
		}
	}

	return "", style{}
}

func promptSymbol() (string, style) {
	defaultStyle := style{fg: "#fab387"}

	_, month, day := now.Date()
	if month == time.November && day == 14 {
		return "󰃩", defaultStyle
	}

	if month == time.April && day == 1 {
		return "ඞ", defaultStyle
	}

	if month == time.October && day == 31 {
		return "󰮿", defaultStyle
	}

	if (month == time.December || month == time.June) && day == 21 {
		return "", defaultStyle
	}

	if (month == time.December && day >= 31-14) ||
		(month == time.January && day <= 14) {
		return "", defaultStyle
	}

	return "", defaultStyle
}
