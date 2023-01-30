package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	fibonacci "github.com/dmicheneau/golearn/game/fibonacci"
	morpion "github.com/dmicheneau/golearn/game/morpion"
	romans "github.com/dmicheneau/golearn/game/romans"
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
	redalert = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("9")).
			Width(40).
			Align(lipgloss.Center)
)

// main is the main function
func main() {
	for {
		fmt.Println(style.Render("welcome to game"))
		fmt.Println(style.Render("1 - play morpion"))
		fmt.Println(style.Render("2 - play fibonacci"))
		fmt.Println(style.Render("3 - play romans"))
		fmt.Println(style.Render("4 - quit"))
		fmt.Println(style.Render("what do you want to do ?"))
		var reponse string
		fmt.Scanln(&reponse)
		switch reponse {
		case "1":
			morpion.Start()
		case "2":
			fibonacci.Start()
		case "3":
			romans.Start()
		case "4":
			os.Exit(0)
		default:
			fmt.Println(redalert.Render("please enter a valid number"))
			break
		}
	}
}
