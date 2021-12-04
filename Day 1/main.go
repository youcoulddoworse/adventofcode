package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	data, err := getData("report.txt")
	if err != nil {
		log.Fatal(err)
	}

	increases := 0
	previous := data[0]

	for _, thisData := range data {
		if thisData > previous {
			increases++
		}

		previous = thisData
	}

	fmt.Println(increases)

}

func getData(filename string) ([]int, error) {

	var result []int

	// open file
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file word by word using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil

}
