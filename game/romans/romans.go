package romans

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/lipgloss"
)

var (
	romains2int = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

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
	var reponse string
	s := bufio.NewScanner(os.Stdin)
	fmt.Printf(style.Render("Enter a Romans Number (M D C L X V I):"))
	for s.Scan() {
		reponse = s.Text()
		if s.Err() != nil {
			log.Println(s.Err())
		}
		break
	}
	fmt.Printf(char.Render("%v"), getRomans(reponse))
}

func getRomans(n string) int {
	var result int
	for i := 0; i < len(n); i++ {
		if i > 0 && romains2int[string(n[i])] > romains2int[string(n[i-1])] {
			result += romains2int[string(n[i])] - 2*romains2int[string(n[i-1])]
		} else {
			result += romains2int[string(n[i])]
		}
	}
	return result
}
