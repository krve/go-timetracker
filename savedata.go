package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

var (
	saveFile, _ = filepath.Abs("./entries.json")
)

// SaveData : The save file structure
type SaveData struct {
	Entries []TimeEntry `json:"entries"`
}

// SetEntries sets the entries on the save
func (data *SaveData) SetEntries(entries []TimeEntry) {
	data.Entries = entries
}

// ListEntries lists all the entries in the terminal
func (data *SaveData) ListEntries(filter string) {
	var tableData [][]string

	for _, entry := range data.Entries {
		if filter != "" && strings.Contains(strings.ToLower(entry.Description), filter) == false {
			continue
		}

		tableData = append(tableData, []string{
			strconv.Itoa(entry.ID),
			entry.Description,
			FormatDuration(entry.Duration),
			entry.StartTime.Format("2006-01-02 15:04:05"),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Description", "Duration", "Date"})
	table.SetBorder(false)
	table.SetRowSeparator("-")
	table.SetRowLine(true)
	table.AppendBulk(tableData)
	table.Render()
}

// ClearEntries clears all entries from the save
func (data *SaveData) ClearEntries(filter string) {
	if filter == "" {
		data.Entries = nil

		return
	}

	entries := data.Entries

	for i := 0; i < len(entries); i++ {
		if filter != "" && strings.Contains(strings.ToLower(entries[i].Description), filter) == false {
			continue
		}

		entries = append(entries[:i], entries[i+1:]...)

		// Since we just deleted the index, we must redo it
		i--
	}

	data.SetEntries(entries)
}

// DeleteEntry deletes the entry with the specified id
func (data *SaveData) DeleteEntry(ID int) bool {
	entries := data.Entries

	for index, v := range entries {
		if v.ID == ID {
			Data.SetEntries(append(entries[:index], entries[index+1:]...))
			Data.Save()

			return true
		}
	}

	return false
}

// GetLatestEntryID fetches the latest ID used
func (data *SaveData) GetLatestEntryID() int {
	entries := data.Entries

	if len(entries) == 0 {
		return 0
	}

	return entries[len(entries)-1].ID
}

// Load loads the data from the save file
func (data *SaveData) Load() {
	*data = LoadSaveDataFromFile()
}

// Save saves the data to the save file
func (data *SaveData) Save() {
	SaveDataToFile(*data)
}

// LoadSaveDataFromFile loads the save data from the save file
func LoadSaveDataFromFile() SaveData {
	log.Println("Loading data from file")

	var contents, _ = ioutil.ReadFile(saveFile)
	var jsonBlob = []byte(contents)
	var save = SaveData{}

	json.Unmarshal(jsonBlob, &save)

	return save
}

// SaveDataToFile saves the input struct to the specified save file
func SaveDataToFile(save SaveData) {
	log.Println("Saving data to file")

	saveJSON, _ := json.Marshal(save)

	err := ioutil.WriteFile(saveFile, saveJSON, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
