package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func main() {
	const version = 1

	idx := 0
	input := make([][]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimRight(line, "\n")
		if len(line) == 0 {
			break
		}

		input = append(input, make([]int, 0))

		for jdx, ch := range line {
			num, err := strconv.Atoi(string(ch))
			if err != nil {
				slog.Error("bad input", "err", err, "idx", idx, "jdx", jdx, "ch", string(ch))
				return
			}

			input[idx] = append(input[idx], num)
		}

		slog.Info("processed line", "line", line, "idx", idx, "input", input[idx])
		idx++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	slog.Info("input", "x", len(input), "y", len(input[0]))
	vis := InitVisibleMap(input)
	slog.Info("init visible map", "vis", CountVisible(vis))
	PopulateVisibleMap(vis, input)
	fmt.Println(String2DArray(vis))
	slog.Info("visible map", "vis", String2DArray(vis))
	fmt.Println(CountVisible(vis))

	slog.Info("done")
}

type Visibility struct {
	Top, Bottom, Left, Right             bool
	TopVal, BottomVal, LeftVal, RightVal int
}

func (v Visibility) Visible() bool {
	return v.Top || v.Bottom || v.Left || v.Right
}

func (v Visibility) String() string {
	if v.Visible() {
		return "X"
	}
	return "o"
}

// InitVisibleMap creates a 2D array of booleans with the given dimensions
// and sets the border to true.
func InitVisibleMap(input [][]int) [][]Visibility {
	x := len(input)
	y := len(input[0])

	visibleMap := make([][]Visibility, x)
	for idx := 0; idx < x; idx++ {
		visibleMap[idx] = make([]Visibility, y)
		visibleMap[idx][0].Left = true
		visibleMap[idx][0].LeftVal = input[idx][0]
		visibleMap[idx][y-1].Right = true
		visibleMap[idx][y-1].RightVal = input[idx][y-1]

		if idx == 0 {
			for jdx := 0; jdx < y; jdx++ {
				visibleMap[idx][jdx].Top = true
				visibleMap[idx][jdx].TopVal = input[idx][jdx]
			}
		}

		if idx == x-1 {
			for jdx := 0; jdx < y; jdx++ {
				visibleMap[idx][jdx].Bottom = true
				visibleMap[idx][jdx].BottomVal = input[idx][jdx]
			}
		}
	}
	return visibleMap
}

// PopulateVisibleMap takes a map of height of trees and an initial map of visibility
// and populates the visibility map with the visible trees.
func PopulateVisibleMap(visibleMap [][]Visibility, input [][]int) {
	for idx := 1; idx < len(input)-1; idx++ {
		for jdx := 1; jdx < len(input[idx])-1; jdx++ {
			visibleMap[idx][jdx].Top = input[idx][jdx] > visibleMap[idx-1][jdx].TopVal
			visibleMap[idx][jdx].Left = input[idx][jdx] > visibleMap[idx][jdx-1].LeftVal

			if visibleMap[idx][jdx].Top {
				visibleMap[idx][jdx].TopVal = input[idx][jdx]
			} else {
				visibleMap[idx][jdx].TopVal = visibleMap[idx-1][jdx].TopVal
			}

			if visibleMap[idx][jdx].Left {
				visibleMap[idx][jdx].LeftVal = input[idx][jdx]
			} else {
				visibleMap[idx][jdx].LeftVal = visibleMap[idx][jdx-1].LeftVal
			}
		}
	}

	for idx := len(input) - 2; idx >= 1; idx-- {
		for jdx := len(input[idx]) - 2; jdx >= 1; jdx-- {
			visibleMap[idx][jdx].Bottom = input[idx][jdx] > visibleMap[idx+1][jdx].BottomVal
			visibleMap[idx][jdx].Right = input[idx][jdx] > visibleMap[idx][jdx+1].RightVal

			if visibleMap[idx][jdx].Bottom {
				visibleMap[idx][jdx].BottomVal = input[idx][jdx]
			} else {
				visibleMap[idx][jdx].BottomVal = visibleMap[idx+1][jdx].BottomVal
			}

			if visibleMap[idx][jdx].Right {
				visibleMap[idx][jdx].RightVal = input[idx][jdx]
			} else {
				visibleMap[idx][jdx].RightVal = visibleMap[idx][jdx+1].RightVal
			}
		}
	}
}

// CountVisible counts the number of visible trees in the map.
func CountVisible(visibleMap [][]Visibility) int {
	total := 0
	for idx := 0; idx < len(visibleMap); idx++ {
		for jdx := 0; jdx < len(visibleMap[idx]); jdx++ {
			if visibleMap[idx][jdx].Visible() {
				total++
			}
		}
	}
	return total
}

func String2DArray[T any](arr [][]T) string {
	sb := strings.Builder{}
	for idx := 0; idx < len(arr); idx++ {
		sb.WriteString(fmt.Sprintln(arr[idx]))
	}
	return sb.String()
}
