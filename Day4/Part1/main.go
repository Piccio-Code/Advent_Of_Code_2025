package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	matrix, err := GetMatrix("../input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	result := AccessByForklift(matrix)

	fmt.Println(result)
}

func AccessByForklift(matrix [][]int) (sum int) {
	sum = 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			value := 0

			if matrix[i][j] != 0 {
				value = CanBeAccessed(i, j, matrix)
			}

			sum += value
		}
	}

	return sum
}

func CanBeAccessed(i, j int, matrix [][]int) int {
	sum := 0

	for k := -1; k <= 1; k++ {
		if i+k >= 0 && i+k < len(matrix) {
			sum += GetSumSide(matrix[i+k][max(j-1, 0):min(len(matrix[0]), j+2)])
		}
	}

	if sum <= 4 {
		return 1
	}

	return 0
}

func GetSumSide(side []int) int {
	sum := 0

	for _, v := range side {
		sum += v
	}

	return sum
}

func GetMatrix(path string) ([][]int, error) {
	file, err := os.Open(path)
	matrix := make([][]int, 0)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	reader := bufio.NewScanner(file)

	for reader.Scan() {
		text := reader.Text()
		matrix = append(matrix, NewRow(text))
	}

	return matrix, nil
}

func NewRow(line string) []int {
	row := make([]int, 0, len(line))

	for _, c := range line {
		if c == '.' {
			row = append(row, 0)
		} else {
			row = append(row, 1)
		}
	}

	return row
}
