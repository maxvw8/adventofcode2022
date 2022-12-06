package advent

import (
	"advent/pkg/utils"
	"fmt"
	"io"
)

type Day4 struct {
	pairs []SectionPair
}
type elfAssignment struct {
	from, to int
}

func (a elfAssignment) contains(b elfAssignment) bool {
	return a.from <= b.from && a.to >= b.to
}
func (a elfAssignment) overlap(b elfAssignment) bool {
	return a.to >= b.from && a.from <= b.from || a.from <= b.to && b.to <= a.to
}

type SectionPair struct {
	e1, e2 elfAssignment
}

func (s SectionPair) isContaining() bool {
	return s.e1.contains(s.e2) || s.e2.contains(s.e1)
}
func (s SectionPair) isOverlap() bool {
	return s.e1.overlap(s.e2) || s.isContaining()
}

func NewDay4() *Day4 {
	return &Day4{}
}

func (d Day4) FileName() string {
	return "input4.txt"
}

func (d *Day4) Load(reader io.Reader) error {
	pair, err := utils.ReadFile(reader, readSectionPair)
	d.pairs = pair
	return err
}
func readSectionPair(l string) (SectionPair, error) {
	var a, b elfAssignment
	_, err := fmt.Sscanf(l, "%d-%d,%d-%d", &a.from, &a.to, &b.from, &b.to)
	if err != nil {
		return SectionPair{}, err
	}
	return SectionPair{a, b}, nil
}

func (d *Day4) Solve() (string, error) {
	sum := 0
	for _, pair := range d.pairs {
		if pair.isOverlap() {
			sum++
		}
	}
	return fmt.Sprintf("%d", sum), nil
}
