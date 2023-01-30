package fibonacci

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
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

func Start() {
	var reponse int
	fmt.Printf(style.Render("Enter a number:"))
	fmt.Scanln(&reponse)
	if reponse < 2 {
		panic("number is too small")
	}
	fmt.Printf(char.Render("%v"), getFibonacci(reponse))
}

func getFibonacci(n int) []int {
	var fibo []int
	var a, b int = 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
		fibo = append(fibo, a)
	}
	return fibo
}
