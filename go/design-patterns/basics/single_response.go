package basics

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)

	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) int {
	entryCount--
	j.entries = append(j.entries[:index], j.entries[index+1:]...)
	return entryCount
}

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	slice := strings.Split(string(data), "\n")
	for _, s := range slice {
		entryCount++
		j.entries = append(j.entries, s)
	}
	return entryCount
}
