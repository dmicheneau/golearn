package banque

import (
	"errors"
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

type Account struct {
	surname string
	name    string
}

type Employee struct {
	credits float64
	Account
}

func (e *Employee) GetCredits() float64 {
	return e.credits
}
func (e *Employee) AddCredits(credits float64) float64 {
	e.credits += credits
	return e.credits
}
func (e *Employee) RemoveCredits(credits float64) (float64, error) {
	if e.credits < credits {
		return e.credits, errors.New("You don't have enough credits")
	}
	e.credits -= credits
	return e.credits, nil
}

func (e Employee) String() string {
	return fmt.Sprintf(style.Render("Mr ou Mme %s %s have credits:%v"), e.surname, e.name, e.credits)
}

func (e *Employee) ChangeName() {
	for {
		fmt.Println(style.Render("Enter your name:"))
		_, err := fmt.Scan(&e.name)
		if err != nil {
			fmt.Println(char.Render("Please enter a valid name"))
			continue
		}
		break
	}
	for {
		fmt.Println(style.Render("Enter your surname:"))
		_, err := fmt.Scan(&e.surname)
		if err != nil {
			fmt.Println(char.Render("Please enter a valid name"))
			continue
		}
		break
	}
}

func Start() {
	fmt.Println(style.Render("Utilisation methode et instance de struct"))
	employee := Employee{credits: 1000, Account: Account{name: "John", surname: "Doe"}}
	fmt.Printf(style.Render("%v"), employee)
	employee.ChangeName()
	fmt.Printf(char.Render("%v"), employee)
	for {
		fmt.Println(style.Render("Enter the amount of credits to add:"))
		var credits float64
		_, err := fmt.Scanln(&credits)
		if err != nil {
			fmt.Println(style.Render("Please enter a valid number"))
			continue
		}
		employee.AddCredits(credits)
		break
	}
	fmt.Printf(char.Render("%s"), employee)
	//fmt.Println(style.Render("Add Credits:>"), employee.AddCredits(100))
	for {
		fmt.Println(style.Render("Enter the amount of credits to remove:"))
		var credits float64
		_, err := fmt.Scanln(&credits)
		if err != nil {
			fmt.Println(style.Render("Please enter a valid number"))
			continue
		}
		_, err = employee.RemoveCredits(credits)
		if err != nil {
			fmt.Println(style.Render("Error:>"), err)
			continue
		}
		break

	}
	fmt.Println(employee)
}
