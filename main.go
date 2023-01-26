package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	morpion "github.com/dmicheneau/golearn/game"
)

var (
	style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(1).
		PaddingLeft(4).
		Width(40)
	char = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("202")).
		Width(4).
		Align(lipgloss.Center)
)

// main is the main function
func main() {
	for {
		fmt.Println(style.Render("welcome to game"))
		fmt.Println(style.Render("1 - play morpion"))
		fmt.Println(style.Render("2 - quit"))
		fmt.Println(style.Render("what do you want to do ?"))
		var reponse string
		fmt.Scanln(&reponse)
		if reponse == "1" {
			morpion.Start()
		} else if reponse == "2" {
			break
		}
	}
}
