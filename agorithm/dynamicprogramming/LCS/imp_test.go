package LCS

import (
	"testing"
	"fmt"
)

func alloc2dimensional(x, y int) [][]int{
	ret := make([][]int, x)
	for index := range ret{
		ret[index] = make([]int, y)
	}
	return ret
}

func TestLCS(t *testing.T) {
	x := []int{5,3,4,5,2,1,9,8}
	y := []int{3,4,5,1,5,3,9,7,5,0}
	c := alloc2dimensional(len(x),len(y))
	fmt.Println("\n",LCS(x,y,len(x)-1, len(y)-1, c), Count, TotalCount)
	for i := range c {
		//fmt.Println(c[i]...)
		for j := range c[i] {
			fmt.Printf("%7d ", c[i][j])
		}
		fmt.Println("")
	}
}