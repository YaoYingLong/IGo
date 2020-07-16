package test

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortString(t *testing.T) {
	names := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Sort(names)
	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
}
func TestSortInt(t *testing.T) {
	{
		names := sort.StringSlice{
			"3. Triple Kill",
			"5. Penta Kill",
			"2. Double Kill",
			"4. Quadra Kill",
			"1. First Blood",
		}
		sort.Sort(names)
		for _, v := range names {
			fmt.Printf("%s\n", v)
		}
	}
	{
		names := []string{
			"3. Triple Kill",
			"5. Penta Kill",
			"2. Double Kill",
			"4. Quadra Kill",
			"1. First Blood",
		}
		sort.Strings(names)
		// 遍历打印结果
		for _, v := range names {
			fmt.Printf("%s\n", v)
		}
	}
}
