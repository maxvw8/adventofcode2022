package utils_test

import (
	"advent/pkg/utils"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		Input    string
		Expected []int
		Fn       utils.StringParseFn[int]
	}{
		{"1\n2\n3\n4\n5\n", []int{1, 2, 3, 4, 5}, strconv.Atoi},
		{"1\n2\n3\n4\n\n5\n", []int{1, 2, 3, 4, 5}, strconv.Atoi},
		{"\n5\n4\n3\n2\n\n1\n", []int{5, 4, 3, 2, 1}, strconv.Atoi},
	}

	for _, tc := range tests {
		t.Run(tc.Input, func(t *testing.T) {
			obtained, err := utils.ReadFile(strings.NewReader(tc.Input), tc.Fn)
			assert.NoError(t, err)
			assert.EqualValues(t, tc.Expected, obtained)
		})
	}
}
func TestReadBlocks(t *testing.T) {
	tests := []struct {
		Input    string
		Expected [][]int
	}{
		{"1\n\n2\n\n3", [][]int{{1}, {2}, {3}}},
		{"1\n2\n\n3", [][]int{{1, 2}, {3}}},
		//{"1\n2\n\n\n3", [][]int{{1, 2}, {3}}},
	}
	for _, tc := range tests {
		t.Run(tc.Input, func(t *testing.T) {
			obtained, err := utils.ReadBlocks(strings.NewReader(tc.Input), strconv.Atoi)
			assert.NoError(t, err)
			assert.EqualValues(t, tc.Expected, obtained)
		})
	}
}
