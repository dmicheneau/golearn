package testurl

import (
	"fmt"
	"net/http"
	"time"

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
		Width(40).
		PaddingLeft(4).
		Align(lipgloss.Center)
	red = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("124")).
		Width(40).
		PaddingLeft(4).
		Align(lipgloss.Center)
)

func Start() {
	ch := make(chan string, 10)

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for {
		var url string
		urlstart := "https://"
		fmt.Println(style.Render("Enter your URL to test:"))
		_, err := fmt.Scan(&url)
		if err != nil {
			fmt.Println(char.Render("Please enter a valid name"))
			continue
		}
		apis = append(apis, urlstart+url)
		break
	}

	start := time.Now()

	// go routine
	for _, api := range apis {
		go checkAPI(api, ch)
	}

	// wait for all go routines to finish
	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}
	elapsed := time.Since(start)
	fmt.Printf(char.Render("Done! It took %v seconds!\n"), elapsed.Seconds())
}

func checkAPI(api string, ch chan string) {

	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf(red.Render("ERROR: %s is down!\n"), api)
		return
	}

	ch <- fmt.Sprintf(char.Render("SUCCESS: %s is up and running!\n"), api)
}
