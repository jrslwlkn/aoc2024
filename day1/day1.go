package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/jrslwlkn/aoc2024/common"
)

func main() {
	data, err := os.ReadFile(common.AbsPath("input.txt"))
	common.Oops(err)

	var list1, list2 []int
	counts := make(map[int]int)

	for _, line := range common.Lines(data) {
		if line == "" {
			break
		}
		nums := strings.Split(line, "   ")
		num1, err := strconv.Atoi(nums[0])
		common.Oops(err)
		list1 = append(list1, num1)
		num2, err := strconv.Atoi(nums[1])
		common.Oops(err)
		list2 = append(list2, num2)

		counts[num2] += 1
	}

	slices.Sort(list1)
	slices.Sort(list2)

	ret1 := 0
	ret2 := 0
	for i, x := range list1 {
		ret1 += common.Abs(list1[i] - list2[i])
		ret2 += x * counts[x]
	}

	fmt.Println("part 1:", ret1)
	fmt.Println("part 2:", ret2)
}
