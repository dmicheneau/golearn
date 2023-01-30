package fibonacci

import (
	"fmt"
	"io"

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
	for {
		fmt.Printf(style.Render("Enter a number:"))
		sizerep, err := fmt.Scanln(&reponse)
		if sizerep < 1 || (err != nil && err != io.EOF) || reponse < 2 {
			fmt.Println("please enter a valid number superior to 1, try again")
			continue
		} else {
			fmt.Printf(char.Render("%d"), getFibonacci(reponse))
			break
		}
	}
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
