package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// AskForConfirmation asks the user for confirmation. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user.
func AskForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

// AwaitInput waits for the user to input something and then returns the input
func AwaitInput(description string) string {
	fmt.Println(description)

	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	sentence, err := buf.ReadBytes('\n')

	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	return strings.TrimSuffix(string(sentence), "\n")
}

// ItoaTwoDigits time.Clock returns one digit on values, so we make sure to convert to two digits
func ItoaTwoDigits(i int) string {
	b := "0" + strconv.Itoa(i)
	return b[len(b)-2:]
}

// FormatDuration formats the duration into hours and minutes
func FormatDuration(d time.Duration) string {
	hoursRaw := math.Floor(d.Minutes() / 60)
	minutesRaw := math.Mod(d.Minutes(), 60)

	hours := int(math.Round(hoursRaw))
	minutes := int(math.Round(minutesRaw))

	if hours > 0 && minutes > 0 {
		return fmt.Sprintf("%d hours %d minutes", hours, minutes)
	} else if hours > 0 && minutes == 0 {
		return fmt.Sprintf("%d hours", hours)
	}

	return fmt.Sprintf("%d minutes", minutes)
}
