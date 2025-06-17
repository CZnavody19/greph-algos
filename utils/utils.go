package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintMatrix(matrix [][]uint16, label string) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		fmt.Println(label)
		fmt.Println("(empty matrix)")
		return
	}

	maxWidth := len("∞")
	for _, row := range matrix {
		for _, value := range row {
			if value != ^uint16(0) {
				width := len(strconv.Itoa(int(value)))
				if width > maxWidth {
					maxWidth = width
				}
			}
		}
	}

	colCount := len(matrix[0])
	totalWidth := colCount*(maxWidth+1) - 1

	padding := (totalWidth - len(label)) / 2
	if padding > 0 {
		fmt.Println(strings.Repeat(" ", padding) + label)
	} else {
		fmt.Println(label)
	}

	for _, row := range matrix {
		for _, value := range row {
			if value == ^uint16(0) {
				fmt.Printf("%*s ", maxWidth, "∞")
			} else {
				fmt.Printf("%*d ", maxWidth, value)
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func PrintPath(path []uint16) {
	if len(path) == 0 {
		fmt.Println("No path found")
		return
	}

	for i, v := range path {
		if i == len(path)-1 {
			fmt.Print(v)
		} else {
			fmt.Print(v, " -> ")
		}
	}

	fmt.Println()
}
