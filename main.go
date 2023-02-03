package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	banque "github.com/dmicheneau/golearn/game/banque"
	fibonacci "github.com/dmicheneau/golearn/game/fibonacci"
	geometry "github.com/dmicheneau/golearn/game/geometry"
	morpion "github.com/dmicheneau/golearn/game/morpion"
	romans "github.com/dmicheneau/golearn/game/romans"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	// style is the style of the title
	style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(1).
		PaddingLeft(4).
		Width(40)
	// redalert is the style of the error message
	redalert = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("9")). // red
			Width(40).
			Align(lipgloss.Center)
)

// main is the main function
func main() {
	// log to file
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to open log file")
	} else {
		log.Logger = log.Output(logFile)
	}
	log.Info().Msg("Enter menu game")
	// Main loop
	for {
		fmt.Println(style.Render("welcome to game"))
		fmt.Println(style.Render("1 - play morpion"))
		fmt.Println(style.Render("2 - play fibonacci"))
		fmt.Println(style.Render("3 - play romans"))
		fmt.Println(style.Render("4 - play geometry"))
		fmt.Println(style.Render("5 - play banque"))
		fmt.Println(style.Render("6 - quit"))
		fmt.Println(style.Render("what do you want to do ?"))

		var reponse string
		fmt.Scanln(&reponse)
		switch reponse {
		case "1":
			log.Info().Msg("starting morpion")
			morpion.Start()
		case "2":
			log.Info().Msg("starting fibonacci")
			fibonacci.Start()
		case "3":
			log.Info().Msg("starting romans")
			romans.Start()
		case "4":
			log.Info().Msg("starting geometry")
			geometry.Start()
		case "5":
			log.Info().Msg("starting banque")
			banque.Start()
		case "6":
			log.Info().Msg("quitting")
			os.Exit(0)
		default:
			log.Warn().Msg("invalid choice")
			fmt.Println(redalert.Render("please enter a valid number"))
			break
		}
	}
}
