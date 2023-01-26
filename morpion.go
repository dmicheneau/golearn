package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	sigle1 string = "O"
	sigle2 string = "X"
	grid1  string = "1"
	grid2  string = "2"
	grid3  string = "3"
	grid4  string = "4"
	grid5  string = "5"
	grid6  string = "6"
	grid7  string = "7"
	grid8  string = "8"
	grid9  string = "9"
)

var (
	number  int = 1
	player1 Player
	player2 Player
	grid    = [9]string{grid1, grid2, grid3, grid4, grid5, grid6, grid7, grid8, grid9}
	sigle   = [2]string{sigle1, sigle2}
	reponse string
	style   = lipgloss.NewStyle().
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

type Player struct {
	Name string
	Sigl string
	Num  int
}

// setPlayer is the function that will set the name, sigle and number of the player
func setPlayer(number int) Player {
	buf := bufio.NewReader(os.Stdin)
	fmt.Printf(style.Render("player %d what's your name ?"), number)
	player_name, _ := buf.ReadString('\n')
	player_name = strings.TrimSuffix(player_name, "\n")
	p := Player{player_name, sigle[(number - 1)], number}
	return p
}

// printGridV2 is the function that will print the grid with color
func printGridV2(grid [9]string) {
	for i := 0; i < 9; i++ {
		if grid[i] == sigle[0] || grid[i] == sigle[1] {
			fmt.Printf(char.Render("%s"), grid[i])
		} else {
			fmt.Printf(" %s ", grid[i])
		}
		if i == 2 || i == 5 || i == 8 {
			fmt.Println()
		}
	}
}

// readEntry is the function that will read the entry of the player
func readEntry(s bufio.Scanner) int {
	for s.Scan() {
		reponse = s.Text()
		if s.Err() != nil {
			log.Println(s.Err())
		}
		break
	}
	i, err := strconv.Atoi(reponse)
	if err != nil {
		log.Println(err, reponse)
	}
	return i
}

// playerTour is the function that will be called for each player
func playerTour(p Player) {
	var i int
	buf := bufio.NewScanner(os.Stdin)
	fmt.Printf(style.Render("Player %s (sigle %s) please enter a number"), p.Name, p.Sigl)
	for {
		i = readEntry(*buf)
		if i >= 1 && i <= 9 {
			break
		}
		fmt.Println(style.Render("please enter a number between 1 and 9"))
	}
	i--
	for {
		if grid[i] != sigle[0] && grid[i] != sigle[1] {
			break
		}
		fmt.Println(style.Render("please enter a number between 1 and 9 and not already used"))
		i = readEntry(*buf) - 1
	}
	grid[i] = p.Sigl
	printGridV2(grid)
}

func checkWin(g [9]string) bool {
	if g[0] == g[1] && g[1] == g[2] {
		return true
	}
	if g[3] == g[4] && g[4] == g[5] {
		return true
	}
	if g[6] == g[7] && g[7] == g[8] {
		return true
	}
	if g[0] == g[3] && g[3] == g[6] {
		return true
	}
	if g[1] == g[4] && g[4] == g[7] {
		return true
	}
	if g[2] == g[5] && g[5] == g[8] {
		return true
	}
	if g[0] == g[4] && g[4] == g[8] {
		return true
	}
	if g[2] == g[4] && g[4] == g[6] {
		return true
	}
	return false
}

// main is the main function
func main() {
	player1 := setPlayer(1)
	player2 := setPlayer(2)
	printGridV2(grid)
	for i := 0; i < 5; i++ {
		playerTour(player1)
		if checkWin(grid) {
			fmt.Printf(style.Render("player %s win !!!"), player1.Name)
			break
		}
		if i == 4 {
			fmt.Println(style.Render("draw match !!!"))
			break
		}
		playerTour(player2)
		if checkWin(grid) {
			fmt.Printf(style.Render("player %s win !!!"), player2.Name)
			break
		}
	}
}
