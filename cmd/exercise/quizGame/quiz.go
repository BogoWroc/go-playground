package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format question,answer")
	flag.Parse()

	questions, e := readCsvFile(*csvFilename)
	if e == nil {
		for q, a := range questions {
			fmt.Println(q)
			n, _ := fmt.Scanln()
			if n == a {
				fmt.Printf("Correct!\n")
			} else {
				fmt.Printf("Incorrect!\n")
			}

		}
	} else {
		fmt.Printf(e.Error())
		os.Exit(-1)
	}

}

func readCsvFile(pathToFile string) (map[string]int, error) {
	csvFile, err := os.Open(pathToFile)
	if err != nil {
		return nil, fmt.Errorf("unable to open %v file", pathToFile)
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))
	lines := make(map[string]int, 100)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		expression := line[0]
		expectedValue, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, fmt.Errorf("line '%v' contains invalid expression", line)
		}
		lines[expression] = expectedValue
	}

	return lines, err
}
