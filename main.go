package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

// single responsibility principle
var entrycount int = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entrycount++
	entry := fmt.Sprintf("%d:  %s", entrycount, text)
	j.entries = append(j.entries, entry)
	return entrycount

}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) RemoveEntry(index int) {

}
func (j *Journal) Load(filename string) {}

func (j *Journal) LoadFromWeb(url *url.URL) {}

// separation of concern { splitiing up of the variables}
// saving the entries of the journal in the file
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

// presisting the given stuff
var Lineseparator = "\n"

func SaveToFile(j *Journal, filename string) {

	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, Lineseparator)), 0644)
}

type Persistance struct {
	lineseparator string
}

func (p *Persistance) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, Lineseparator)), 0644)

}

func main() {
	j := Journal{}
	j.AddEntry("stsrted to worry about future today")
	j.AddEntry("everything changed")
	fmt.Println(j.String())

	SaveToFile(&j, "journal.txt")

	p := Persistance{"\r  \n"}
	p.SaveToFile(&j, " journal entry")
}
