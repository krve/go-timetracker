package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
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

// ClearEntries clears all entries from the save
func (data *SaveData) ClearEntries() {
	data.Entries = nil
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
