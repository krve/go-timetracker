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
	tables := []struct {
		x SaveData
		y string
		z int
	}{
		{
			SaveData{
				Entries: []TimeEntry{
					TimeEntry{Description: "foo"},
					TimeEntry{Description: "bar"},
					TimeEntry{Description: "foobar"},
				},
			},
			"foo",
			1,
		},
		{
			SaveData{
				Entries: []TimeEntry{
					TimeEntry{Description: "foo was here"},
					TimeEntry{Description: "foockin"},
					TimeEntry{Description: "barfood"},
				},
			},
			"foo",
			0,
		},
		{
			SaveData{
				Entries: []TimeEntry{
					TimeEntry{Description: "foo was here"},
					TimeEntry{Description: "foockin"},
					TimeEntry{Description: "barfood"},
				},
			},
			"",
			0,
		},
		{
			SaveData{
				Entries: []TimeEntry{
					TimeEntry{Description: "foo was here"},
					TimeEntry{Description: "foockin"},
					TimeEntry{Description: "no food here"},
				},
			},
			"bar",
			3,
		},
	}

	for _, table := range tables {
		table.x.ClearEntries(table.y)

		result := len(table.x.Entries)

		if result != table.z {
			t.Errorf("The result was incorrect, the length of the array was %d, expected %d.", result, table.z)
		}
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
