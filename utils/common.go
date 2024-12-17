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
