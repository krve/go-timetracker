package main

import (
	"fmt"
	"time"
)

// TimeEntry : An entry of time
type TimeEntry struct {
	ID          int           `json:"id"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration"`
	StartTime   time.Time     `json:"start_time"`
	EndTime     time.Time     `json:"end_time"`
}

// Print prints out the entry in a readable fashion
func (entry TimeEntry) Print() {
	fmt.Println("ID: \t\t", entry.ID)
	fmt.Println("Description: \t", entry.Description)
	fmt.Println("Duration: \t", FormatDuration(entry.Duration))
	// fmt.Println("Start time: \t", entry.StartTime)
	// fmt.Println("End time: \t", entry.EndTime)
}
