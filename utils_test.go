package main

import "testing"

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
