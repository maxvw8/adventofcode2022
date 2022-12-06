package advent

import (
	"advent/pkg/utils"
	"fmt"
	"io"
	"unicode"
)

const GroupSize = 3

type Item rune
type Day3 struct {
	rucksacks []Rucksack
}
type Rucksack struct {
	r1 map[rune]int
	r2 map[rune]int
	r3 map[rune]int
}

func NewDay3() *Day3 {
	return &Day3{}
}
func (d Day3) FileName() string {
	return "input3.txt"
}

func (d *Day3) Load(reader io.Reader) error {
	//every 3 lines
	lines, err := utils.Readlines(reader)
	var rucksacks []Rucksack
	if err != nil {
		return err
	}
	if len(lines) < GroupSize || len(lines)%GroupSize != 0 {
		return fmt.Errorf("cannot process rucksacks, %q rucksacks read, not multiple of 3", len(lines))
	}
	for i := GroupSize - 1; i < len(lines); i += GroupSize {
		if (i+1)%GroupSize == 0 {
			rucksacks = append(rucksacks, createRucksack(lines[i], lines[i-1], lines[i-2]))
		}
	}

	d.rucksacks = rucksacks
	return nil
}

func createRucksack(r1, r2, r3 string) Rucksack {
	return Rucksack{letterMap(r1), letterMap(r2), letterMap(r3)}
}

func letterMap(s string) map[rune]int {
	var keys = make(map[rune]int)
	for _, l := range s {
		keys[l] += 1
	}
	return keys
}

func (r Rucksack) getShared() Item {
	//if key is found in all rucksacks
	for k, _ := range r.r1 {
		if r.r2[k] != 0 && r.r3[k] != 0 {
			return Item(k)
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
		sum += item.ToPriority()
	}
	return fmt.Sprintf("%d", sum), nil
}

func (i Item) ToPriority() int {
	if unicode.IsUpper(rune(i)) {
		return int(i - 'A' + 27)
	}
	return int(i - 'a' + 1)
}
