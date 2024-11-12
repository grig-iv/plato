package main

import (
	"fmt"
	"os"
	"strings"
)

type lableMatch struct {
	pattern string
	lable   string
	style   style
}

var (
	homeDir       = os.Getenv("HOME") + "/"
	workingDir, _ = os.Getwd()
)

var dirLabels = []lableMatch{
	{pattern: "~/.config/*", lable: ""},
	{pattern: "~/extended mind/*", lable: "󰧑"},
	{pattern: "~/desktop/*", lable: "󰇄"},
	{pattern: "~/cloud/*", lable: ""},
	{pattern: "~/sources/*", lable: ""},
	{pattern: "~/", lable: ""},

	{pattern: "*/downloads/*", lable: "󱃩"},
	{pattern: "*/books/*", lable: "󱉟"},

	{pattern: "/etc/nixos/*", lable: "󱄅"},
	{pattern: "/run/media/*", lable: "󰕓"},
	{pattern: "/mnt/c/*", lable: ""},
	{pattern: "/mnt/*", lable: "󰋊"},
	{pattern: "/", lable: "/"},
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
	fmt.Print("\n")

	print(lable())

	fmt.Print("  ")

	print(promptSymbol())
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
	return " ", style{fg: "#fab387"}
}
