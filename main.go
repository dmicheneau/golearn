package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	fibonacci "github.com/dmicheneau/golearn/fibonacci"
	morpion "github.com/dmicheneau/golearn/morpion"
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
		fmt.Println(style.Render("2 - play fibonacci"))
		fmt.Println(style.Render("3 - quit"))
		fmt.Println(style.Render("what do you want to do ?"))
		var reponse string
		fmt.Scanln(&reponse)
		switch reponse {
		case "1":
			morpion.Start()
		case "2":
			fibonacci.Start()
		case "3":
			os.Exit(0)
		default:
			fmt.Println(style.Render("please enter a valid number"))
			break
		}
	}
}
