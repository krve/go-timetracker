package main

import "testing"

// SetEntries sets the entries on the save
func TestSetEntries(t *testing.T) {
	data := SaveData{}
	data.SetEntries([]TimeEntry{
		TimeEntry{ID: 1337},
	})

	if len(data.Entries) != 1 {
		t.Errorf("The result was incorrect, the length of the array was %d, expected %d", len(data.Entries), 1)
	}

	if data.Entries[0].ID != 1337 {
		t.Errorf("The result was incorrect, did not find entry with ID: %d.", 1337)
	}
}

func TestClearEntries(t *testing.T) {
	data := SaveData{
		Entries: []TimeEntry{
			TimeEntry{ID: 1},
		},
	}

	data.ClearEntries()

	if len(data.Entries) != 0 {
		t.Errorf("The result was incorrect, the length of the array was %d, expected %d", len(data.Entries), 0)
	}
}

func TestGetLatestEntryID(t *testing.T) {
	tables := []struct {
		x SaveData
		y int
	}{
		{
			SaveData{
				Entries: []TimeEntry{
					TimeEntry{ID: 1},
					TimeEntry{ID: 2},
					TimeEntry{ID: 3},
				},
			},
			3,
		},
		{
			SaveData{
				Entries: []TimeEntry{},
			},
			0,
		},
	}

	for _, table := range tables {
		result := table.x.GetLatestEntryID()

		if result != table.y {
			t.Errorf("The result was incorrect, got: %q, want: %q.", result, table.y)
		}
	}
}

func ExampleListEntries() {
	data := SaveData{
		Entries: []TimeEntry{
			TimeEntry{ID: 1, Description: "Foo"},
			TimeEntry{ID: 2, Description: "Bar"},
		},
	}

	data.ListEntries()
	// Output:
	// ID: 		 1
	// Description: 	 Foo
	// Duration: 	 0 seconds
	// ---
	// ID: 		 2
	// Description: 	 Bar
	// Duration: 	 0 seconds
	// ---
}
