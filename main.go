package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
)

var (
	// Data contains the cached/loaded data from the save file
	Data = SaveData{}
)

func main() {
	app := cli.NewApp()
	app.Name = "Time Tracker CLI"
	app.Version = "0.0.1"
	app.Usage = "Track time spent on different things"
	// app.Flags = []cli.Flag{
	// 	cli.BoolFlag{
	// 		Name:  "cli",
	// 		Usage: "use CLI to stop the timer",
	// 	},
	// }
	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "show the list of entries",
			Action: func(c *cli.Context) error {
				Data.ListEntries()

				return nil
			},
			// Flags: []cli.Flag{
			// 	cli.StringFlag{
			// 		Name:  "lang",
			// 		Value: "english",
			// 		Usage: "language for the greeting",
			// 	},
			// },
		},
		{
			Name:  "clear",
			Usage: "clear all entries",
			Action: func(c *cli.Context) error {
				confirm := AskForConfirmation("Are you sure you want to erase all current entries?")

				if confirm {
					Data.ClearEntries()
					Data.Save()
					fmt.Println("Cleared all entries")
				}

				return nil
			},
		},
		{
			Name:  "delete",
			Usage: "delete the entry with the specified id",
			Action: func(c *cli.Context) error {
				input := c.Args().Get(0)

				if input == "" {
					fmt.Println("You need to specify a ID")
				}

				ID, _ := strconv.Atoi(input)

				confirm := AskForConfirmation("Are you sure you want to erase the entry?")

				if confirm {
					success := Data.DeleteEntry(ID)

					if success == false {
						fmt.Println("No entry with the specified ID found")
					} else {
						fmt.Println("Removed the entry")
					}
				}

				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		createEntry()

		return nil
	}

	Data.Load()

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}

func createEntry() {
	start := time.Now()

	description := AwaitInput("Enter a description and press enter to start time tracking:")
	var entryDuration time.Duration

	fmt.Println("Press \"enter\" to finish tracking time.")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	entryDuration = time.Since(start)

	description = strings.TrimSuffix(description, "\n")
	end := time.Now()

	entry := &TimeEntry{
		ID:          (Data.GetLatestEntryID() + 1),
		Description: description,
		Duration:    entryDuration,
		StartTime:   start,
		EndTime:     end,
	}
	entry.Print()

	Data.SetEntries(append(Data.Entries, *entry))
	Data.Save()
}
