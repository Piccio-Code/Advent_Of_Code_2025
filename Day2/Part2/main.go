package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	startId int
	endId   int
}

func main() {
	ranges, err := getInputValues("../input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0

	for _, r := range ranges {
		sum += checkValidity(r)
	}

	fmt.Println(sum)
}

func checkValidity(r Range) (sum int) {

	for id := r.startId; id <= r.endId; id++ {
		idStr := strconv.Itoa(id)
		pattern := ""

		for _, c := range idStr[:len(idStr)/2] {
			pattern += string(c)

			count := strings.Count(idStr, pattern)

			if count*len(pattern) == len(idStr) {
				sum += id
				break
			}
		}
	}

	return sum
}

func getInputValues(path string) ([]Range, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := bufio.NewReader(file)

	ranges := make([]Range, 0)

	for {
		data, err := reader.ReadString(',')

		if err != nil && err != io.EOF {
			return nil, err
		} else if err == io.EOF {
			break
		}

		startId, endId, err := getIdRange(data)

		if err != nil {
			return nil, err
		}

		ranges = append(ranges, Range{startId: startId, endId: endId})
	}

	return ranges, nil
}

func getIdRange(data string) (startId, endId int, err error) {
	data = data[:len(data)-1]

	strRange := strings.Split(data, "-")

	startId, err = strconv.Atoi(strRange[0])

	if err != nil {
		return 0, 0, err
	}

	endId, err = strconv.Atoi(strRange[1])

	if err != nil {
		return 0, 0, err
	}

	return startId, endId, nil
}
