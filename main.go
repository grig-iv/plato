package main

import (
	"fmt"
	"strconv"
)

type style struct {
	fg string
	bg string
}

var (
	promptSymbolStyle = style{fg: "#FAB387"}
)

func main() {
	fmt.Print("\n")
	print("  ", promptSymbolStyle)
}


func print(text string, style style) {
	if style.fg != "" {
		r, g, b := colorValues(style.fg)
		fmt.Printf("\033[38;2;%d;%d;%dm", r, g, b)
	}

	if style.bg != "" {
		r, g, b := colorValues(style.bg)
		fmt.Printf("\033[48;2;%d;%d;%dm", r, g, b)
	}

	// reset colors
	fmt.Printf("\033[0m")
}

func colorValues(color string) (r, g, b int64) {
	r, err := strconv.ParseInt(color[1:3], 16, 16)
	if err != nil {
		panic(err)
	}

	g, err = strconv.ParseInt(color[3:5], 16, 16)
	if err != nil {
		panic(err)
	}

	b, err = strconv.ParseInt(color[5:], 16, 16)
	if err != nil {
		panic(err)
	}

	return r, g, b
}
