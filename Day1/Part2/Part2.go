package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Lock struct {
	currentValue int
	password     int
}

func (l *Lock) UpdateValue(valueString string) error {
	rotation, err := GetRotationFromString(valueString)

	if err != nil {
		return err
	}

	prev := l.currentValue
	l.currentValue += rotation

	if l.currentValue < 0 {
		rotationToNormal := int(math.Ceil(-float64(l.currentValue) / 100))

		l.currentValue += rotationToNormal * 100

		if prev != 0 {
			l.password += rotationToNormal
		} else {
			l.password += rotationToNormal - 1
		}

	} else if l.currentValue > 0 {
		rotationToNormal := int(math.Floor(float64(l.currentValue) / 100))

		l.currentValue = l.currentValue % 100

		if l.currentValue == 0 {
			l.password += rotationToNormal - 1
		} else {
			l.password += rotationToNormal
		}
	}

	if l.currentValue == 0 {
		l.password++
	}

	return nil
}

func GetRotationFromString(valueString string) (int, error) {

	value, err := strconv.Atoi(valueString[1:])

	if err != nil {
		return 0, err
	}

	switch strings.ToLower(string(valueString[0])) {
	case "l":
		return value * -1, nil
	case "r":
		return value, nil
	default:
		return 0, fmt.Errorf("the value %s is not in the right format", valueString)
	}
}

func main() {
	input, err := GetInputValues("input.txt")
	lock := Lock{currentValue: 50}

	if err != nil {
		return
	}

	for _, s := range input {
		err := lock.UpdateValue(s)

		if err != nil {
			fmt.Println(err)
			return
		}

	}

	fmt.Println("============================")
	fmt.Printf("The password is %v\n", lock.password)
	fmt.Println("============================")
}

func GetInputValues(path string) ([]string, error) {
	input, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer func(fi *os.File) {
		err := fi.Close()
		if err != nil {

		}
	}(input)

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	values := make([]string, 0)

	for fileScanner.Scan() {
		values = append(values, fileScanner.Text())
	}

	return values, nil
}
