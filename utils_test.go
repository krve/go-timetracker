package main

import (
	"testing"
	"time"
)

func TestItoaTwoDigits(t *testing.T) {
	tables := []struct {
		x int
		y string
	}{
		{1, "01"},
		{9, "09"},
		{10, "10"},
		{59, "59"},
		{0, "00"},
	}

	for _, table := range tables {
		result := ItoaTwoDigits(table.x)
		if result != table.y {
			t.Errorf("The result was incorrect, got: %q, want: %q.", result, table.y)
		}
	}
}

func TestFormatDuration(t *testing.T) {
	tables := []struct {
		x time.Duration
		y string
	}{
		{time.Minute * 55, "55 minutes"},
		{time.Hour + time.Minute*35, "1 hours 35 minutes"},
		{time.Hour * 4, "4 hours"},
		{time.Second, "0 minutes"},
	}

	for _, table := range tables {
		result := FormatDuration(table.x)
		if result != table.y {
			t.Errorf("The result was incorrect, got: %q, want: %q.", result, table.y)
		}
	}
}
