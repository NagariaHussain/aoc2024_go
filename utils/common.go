package utils

import (
	"bufio"
	"os"
	"strconv"
)

func GetFileAsScanner(filePath string) (*bufio.Scanner, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	return scanner, nil
}

func GetLines(filePath string) (lines []string) {
	scanner, _ := GetFileAsScanner(filePath)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return
}

func Map[T any, M any](items []T, f func(T) M) []M {
	mappedItems := make([]M, len(items))
	for index, item := range items {
		mappedItems[index] = f(item)
	}
	return mappedItems
}

func ToInt(number string) int64 {
	parsedNum, _ := strconv.ParseInt(number, 10, 0)
	return parsedNum
}

type Set struct {
	setMap map[int64]bool
}

func NewSet() Set {
	set := new(Set)
	set.setMap = make(map[int64]bool)
	return *set
}

func (s *Set) Add(ele int64) {
	s.setMap[ele] = true
}

func (s *Set) Has(ele int64) bool {
	_, ok := s.setMap[ele]
	return ok
}
