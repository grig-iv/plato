package main

import "fmt"

func main() {
	fmt.Print("\n")
	fg(0xfa, 0xb3, 0x87)
	fmt.Print("   ")
	resetColors()
}

func fg(r, g, b int) {
	fmt.Printf("\033[38;2;%d;%d;%dm", r, g, b)
}

func bg(r, g, b int) {
	fmt.Printf("\033[48;2;%d;%d;%dm", r, g, b)
}

func resetColors() {
	fmt.Printf("\033[0m")
}
