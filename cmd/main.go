package main

import (
	"advent/pkg/advent"
	"fmt"
	"os"
	"path/filepath"
)

const inputPath = "resources/"

var problems = []advent.Problem{
	//advent.NewDay1(),
	//advent.NewDay2(),
	//advent.NewDay3(),
	advent.NewDay4(),
}

func main() {
	for _, p := range problems {
		path, _ := filepath.Abs(fmt.Sprintf("%s/%s", inputPath, p.FileName()))
		file, err := os.Open(path)
		if err != nil {
			fmt.Printf("Error while open file %s", err)
			os.Exit(1)
		}
		err = p.Load(file)
		if err != nil {
			fmt.Printf("Error parsing file %s", err)
			os.Exit(1)
		}
		sol, err := p.Solve()
		if err != nil {
			fmt.Printf("Error solving problem %s", err)
			os.Exit(1)
		}

		fmt.Printf("[%T] %q\n", p, sol)
	}

}
