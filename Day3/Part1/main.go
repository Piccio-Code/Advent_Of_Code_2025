package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	banks, err := GetInput("../input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0

	for _, bank := range banks {
		fmt.Println(GetMaxBank(bank))

		sum += GetMaxBank(bank)
	}

	fmt.Println(sum)
}

func GetMaxBank(bank string) int {
	maxBattery1 := -1

	firstMaxIndex := math.MaxInt

	for i, charBattery := range bank[:len(bank)-1] {
		battery := int(charBattery - '0')

		if battery > maxBattery1 {
			maxBattery1 = battery
			firstMaxIndex = i
		} else if battery == maxBattery1 {
			firstMaxIndex = min(i, firstMaxIndex)
		}
	}

	maxBattery2 := -1

	for _, charBattery := range bank[firstMaxIndex+1:] {
		maxBattery2 = max(maxBattery2, int(charBattery-'0'))
	}

	return maxBattery1*10 + maxBattery2
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
