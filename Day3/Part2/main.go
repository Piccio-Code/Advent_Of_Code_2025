package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	banks, err := GetInput("../input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0

	for _, bank := range banks {
		sum += GetMaxBank(bank)
		fmt.Println(GetMaxBank(bank))
	}

	fmt.Println(sum)
}

func GetMaxBank(bank string) int {
	largestNumber := ""

	start := 0
	end := len(bank) - 11

	for i := 0; i < 12; i++ {
		maxVal, index := GetMaxString(bank[start:end])

		largestNumber += string(maxVal)

		start += index + 1
		end = len(bank) - 10 + i

		if end > len(bank) {
			end = len(bank)
		}
	}

	result, err := strconv.Atoi(largestNumber)

	if err != nil {
		fmt.Println(err)
	}

	return result
}

func GetMaxString(str string) (maxVal uint8, index int) {
	maxVal = '0'

	for i := len(str) - 1; i >= 0; i-- {

		if str[i] >= maxVal {
			maxVal = str[i]
			index = i
		}

	}

	return maxVal, index
}

func GetInput(path string) (input []string, err error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		input = append(input, reader.Text())
	}

	if reader.Err() != nil {
		return nil, err
	}

	return input, nil
}
