package utils

import (
	"bufio"
	"os"
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
