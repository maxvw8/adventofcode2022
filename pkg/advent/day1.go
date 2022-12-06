package advent

import (
	"advent/pkg/utils"
	"fmt"
	"io"
	"strconv"
)

type Elf struct {
	calories []int
}
type Day1 struct {
	elves []Elf
}

func NewDay1() *Day1 {
	return &Day1{}
}

func (d Day1) FileName() string {
	return "input1.txt"
}
func (d *Day1) Load(reader io.Reader) error {
	calories, err := utils.ReadBlocks(reader, strconv.Atoi)
	if err != nil {
		return err
	}
	var elves []Elf
	for _, cals := range calories {
		elves = append(elves, Elf{calories: cals})
	}
	d.elves = elves
	return nil
}

func (d *Day1) Solve() (string, error) {
	var highesCal []int
	for _, elf := range d.elves[1:] {
		highesCal = append(highesCal, sum(elf.calories))
	}
	pos := getMax(highesCal)
	first := highesCal[pos]
	highesCal = remove(highesCal, pos)
	pos = getMax(highesCal)
	second := highesCal[pos]
	highesCal = remove(highesCal, pos)
	pos = getMax(highesCal)
	third := highesCal[pos]
	highesCal = remove(highesCal, pos)
	return fmt.Sprintf("%d", first+second+third), nil
}
func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func sum(number []int) (total int) {
	for _, n := range number {
		total += n
	}
	return total
}
func getMax(numbers []int) int {
	max := numbers[0]
	i := 0
	for j, number := range numbers[1:] {
		if number > max {
			max = number
			i = j
		}
	}
	return i + 1
}
