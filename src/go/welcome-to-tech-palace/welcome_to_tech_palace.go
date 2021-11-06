package techpalace

import (
	"fmt"
	"strings"
)

var welcomeMsg = "Welcome to the Tech Palace, %s"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf(welcomeMsg, strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := strings.Repeat("*", numStarsPerLine)
	return strings.Join([]string{stars, welcomeMsg, stars}, "\n")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.Trim(strings.Split(oldMsg, "\n")[1], " *")
}
