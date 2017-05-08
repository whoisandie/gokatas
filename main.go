package main

import (
	"fmt"

	"github.com/whoisandie/gokatas/problems"
)

func main() {
	fmt.Println(problems.Thirt(154))

	fmt.Println(problems.RemovNb(10))

	fmt.Println(problems.TortoiseRace(720, 850, 70))

	fmt.Println(problems.ToNato("If you can read"))

	var reduceSteps = []int{357, 112, 28, -52, 644, 119}
	fmt.Println(problems.ReduceBySteps(problems.Som, reduceSteps, 0))

	var reduceSteps1 = []int{10, -32, 190, 300, -42, -38, 50, 405, -46, 225, -31}
	fmt.Println(problems.ReduceBySteps(problems.Maxi, reduceSteps1, reduceSteps1[0]))

	var reduceSteps2 = []int{18, 69, -90, -78, 65, 40}
	fmt.Println(problems.ReduceBySteps(problems.Gcdi, reduceSteps2, reduceSteps2[0]))

	var reduceSteps3 = []int{6, -72, -62, -22, -23, 80}
	fmt.Println(problems.ReduceBySteps(problems.Lcmu, reduceSteps3, reduceSteps3[0]))

	fmt.Println(problems.ListSquared(1, 250))

	fmt.Println(problems.HasUniqueChar("  nAa"))

	var peaksInput = [][]int{
		{1, 2, 3, 3, 0},
		{1, 2, 3, 4, 5, 6, 7},
		{2, 1, 3, 1, 2, 2, 2, 2},
		{2, 1, 3, 1, 2, 2, 2, 2, 1},
		{-3, 12, 12, 9, -3, 14, 5, 1, -4},
		{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3},
		{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1},
		{1, 2, 5, 4, 3, 2, 3, 6, 4, 1, 2, 3, 3, 4, 5, 3, 2, 1, 2, 3, 5, 5, 4}}

	for _, input := range peaksInput {
		fmt.Println(problems.PickPeaks(input))
	}
}
