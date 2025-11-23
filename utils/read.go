package utils

import (
	"fmt"
	"log"
	"os"
)

func ReadString() string {
	var input string
	fmt.Scanf("%s", &input)
	return input
}

func ReadInt() int {
	var input int
	fmt.Scanf("%d", &input)
	return input
}

func ReadFile(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}
