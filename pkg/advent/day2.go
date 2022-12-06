package advent

import (
	"advent/pkg/utils"
	"fmt"
	"io"
	"reflect"
)

type Shape interface {
	score() int
	beats(shape Shape) bool
	opponentShapeForResult(result RoundResult) Shape
}
type Paper rune
type Rock rune
type Scissor rune
type RoundResult rune

func (p Paper) score() int {
	return 2
}

func (p Paper) opponentShapeForResult(expected RoundResult) Shape {
	switch expected {
	case Lose:
		return Rock(RockRune)
	case Win:
		return Scissor(ScissorRune)
	default:
		return Paper(PaperRune)
	}
}

func (p Paper) beats(shape Shape) bool {
	switch shape.(type) {
	case Rock:
		return true
	default:
		return false
	}
}

func (r Rock) score() int {
	return 1
}

func (r Rock) beats(shape Shape) bool {
	switch shape.(type) {
	case Scissor:
		return true
	default:
		return false
	}
}

func (r Rock) opponentShapeForResult(result RoundResult) Shape {
	switch result {
	case Lose:
		return Scissor(ScissorRune)
	case Win:
		return Paper(PaperRune)
	default:
		return Rock(RockRune)
	}
}

func (s Scissor) score() int {
	return 3
}

func (s Scissor) beats(shape Shape) bool {
	switch shape.(type) {
	case Paper:
		return true
	default:
		return false
	}
}

func (s Scissor) opponentShapeForResult(result RoundResult) Shape {
	switch result {
	case Lose:
		return Paper(PaperRune)
	case Win:
		return Rock(RockRune)
	default:
		return Scissor(ScissorRune)
	}
}

const (
	RockRune    rune        = 'A'
	PaperRune   rune        = 'B'
	ScissorRune rune        = 'C'
	Lose        RoundResult = 'X'
	Draw        RoundResult = 'Y'
	Win         RoundResult = 'Z'
)

type Round struct {
	OpponentPlay Shape
	RoundResult
}
type Day2 struct {
	rounds []Round
}

func NewDay2() *Day2 {
	return &Day2{}
}
func (d Day2) FileName() string {
	return "input2.txt"
}

func (d *Day2) Load(reader io.Reader) error {
	rounds, err := utils.ReadFile(reader, ParseRound)
	d.rounds = rounds
	return err
}

func ParseRound(line string) (Round, error) {
	var opponent rune
	var result RoundResult
	_, err := fmt.Sscanf(line, "%c %c", &opponent, &result)
	if err != nil {
		return Round{}, err
	}
	return NewRound(opponent, result)
}

func NewRound(opponent rune, result RoundResult) (Round, error) {
	oppShape, err := parseShape(opponent)
	if err != nil {
		return Round{}, err
	}
	return Round{OpponentPlay: oppShape, RoundResult: result}, nil
}

func parseShape(s rune) (Shape, error) {
	switch s {
	case PaperRune:
		return Paper(s), nil
	case ScissorRune:
		return Scissor(s), nil
	case RockRune:
		return Rock(s), nil
	}
	return nil, fmt.Errorf("Invalid Shape for %q", s)
}
func (r Round) Score() int {
	expectedPlay := r.OpponentPlay.opponentShapeForResult(r.RoundResult)
	score := expectedPlay.score()
	if expectedPlay.beats(r.OpponentPlay) {
		score += 6
	} else if reflect.TypeOf(expectedPlay) == reflect.TypeOf(r.OpponentPlay) {
		score += 3
	}
	return score
}

func (d *Day2) Solve() (string, error) {
	score := 0

	for _, round := range d.rounds {
		score += round.Score()
	}
	return fmt.Sprintf("%d", score), nil
}
