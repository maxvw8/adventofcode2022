package advent

import (
	"advent/pkg/utils"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Item rune
type Day3 struct {
	rucksacks []Rucksack
}
type Rucksack struct {
	c1 string
	c2 string
}

func NewDay3() *Day3 {
	return &Day3{}
}
func (d Day3) FileName() string {
	return "input3.txt"
}

func (d *Day3) Load(reader io.Reader) error {
	lines, err := utils.ReadFile(reader, func(s string) (Rucksack, error) {
		mid := len(s) / 2
		return Rucksack{c1: s[:mid], c2: s[mid:]}, nil
	})
	if err != nil {
		return err
	}
	d.rucksacks = lines
	return nil
}
func (r Rucksack) getShared() Item {
	//var shared []Item
	for _, letter := range r.c1 {
		if strings.ContainsRune(r.c2, letter) {
			//shared = append(shared, Item(letter))
			return Item(letter)
		}
	}
	return 0
}
func (d Day3) Solve() (string, error) {
	var matches []Item
	for _, rucksack := range d.rucksacks {
		matches = append(matches, rucksack.getShared())
	}
	sum := 0
	for _, item := range matches {
		sum += item.toPriority()
	}
	return fmt.Sprintf("%d", sum), nil
}

func (i Item) toPriority() int {
	if unicode.IsUpper(rune(i)) {
		return int(i - 'A' + 27)
	}
	return int(i - 'a' + 1)
}
