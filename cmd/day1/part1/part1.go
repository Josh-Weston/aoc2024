package part1

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	left := []int{}
	right := []int{}
	for scanner.Scan() {
		if scanner.Err() != nil {
			fmt.Println("Error reading input:", scanner.Err())
			os.Exit(1)
		}
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 2 {
			fmt.Println("Invalid input format. Expected two integers per line.")
			os.Exit(1)
		}
		leftInt, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Println("Error converting left field to integer:", err)
			os.Exit(1)
		}
		rightInt, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println("Error converting right field to integer:", err)
			os.Exit(1)
		}
		left = append(left, leftInt)
		right = append(right, rightInt)

	}
	if len(left) != len(right) {
		fmt.Println("Left and right slices must be of the same length.")
		os.Exit(1)
	}

	// Sort th slices in ascending order
	slices.Sort(left)
	slices.Sort(right)

	total := 0
	for i := range len(left) {
		total += int(math.Abs(float64(left[i]) - float64(right[i])))
	}
	return total, nil
}
