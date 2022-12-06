package advent

import "io"

type Problem interface {
	FileName() string
	Load(reader io.Reader) error
	Solve() (string, error)
}
