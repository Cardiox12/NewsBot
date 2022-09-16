package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Database struct {
	Path	string
	Schema	map[string]interface{}
}

func GetDatabaseFullpath(filename string) string {
	path, err := os.Getwd()

	if err != nil {
		log.Fatalf("Error reading working directory")
	}
	return filepath.Join(path, filename)
}

func NewDatabase(path string) *Database {
	return &Database{ 
		Path: path, 
		Schema: make(map[string]interface{}, 0),
	}
}

func (d *Database) Init() {
	f, err := os.Open(d.Path)

	if os.IsNotExist(err) {
		return
	} else if err != nil {
		log.Fatalf("Error opening database file")
	}
	defer f.Close()

	data, _ := ioutil.ReadAll(f)
	json.Unmarshal(data, &d.Schema)
}

func (d *Database) Set(key, value string) {
	d.Schema[key] = value
}

func (d *Database) Get(key string) (interface{}, bool) {
	v, ok := d.Schema[key]

	return v, ok
}

func (d *Database) Write() {
	file, _ := json.MarshalIndent(d.Schema, "", " ")

	err := ioutil.WriteFile(d.Path, file, 0644)
	if err != nil {
		log.Fatal("Failed to write database file")
	}	
}