package geometry

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	// draw the style of the input
	style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(1).
		PaddingLeft(4).
		Width(40)
	// draw entered number of stars
	char = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("202")).
		Width(2).
		Align(lipgloss.Center)
)

type triangle struct {
	size int
}

type square struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (t triangle) draw() {
	for i := 0; i < t.size; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(char.Render("*"))
		}
		fmt.Println()
	}
}

func (s square) perimeter() int {
	return s.size * 4
}

func (s square) draw() {
	for i := 0; i < s.size; i++ {
		for j := 0; j < s.size; j++ {
			fmt.Print(char.Render("*"))
		}
		fmt.Println()
	}
}

func Start() {
	for {
		var size int
		fmt.Print(style.Render("Enter a size: "))
		_, err := fmt.Scanln(&size)
		if err != nil || size < 1 {
			fmt.Println(style.Render("Please enter a valid number (>= 1)"))
			continue
		}
		t := triangle{size}
		t.draw()
		fmt.Printf(style.Render("Perimeter (triangle): %d"), t.perimeter())
		fmt.Println(style.Render("===================================="))
		s := square{size}
		s.draw()
		fmt.Printf(style.Render("Perimeter (square): %d"), s.perimeter())
		fmt.Println(style.Render("===================================="))
		break
	}
}
