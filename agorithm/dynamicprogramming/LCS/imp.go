package LCS

import "fmt"

func max(x, y int) int{
	if x >= y{
		return x
	}
	return y
}

var (
	Count = 0
	TotalCount = 0
)
//LCS 最长子序列
func LCS(x,y []int, i, j int, c [][]int) int{
	TotalCount += 1
	if c[i][j] == 0{
		Count += 1
		if x[i] == y[j]{
			fmt.Printf("%2d %2d %2d\n", x[i], i, j)
			c[i][j] = 1
			if i > 0 && j > 0{
				c[i][j] += LCS(x,y,i-1,j-1,c)
			}
		}else{
			if i > 0 && j > 0{
				c[i][j] += max(LCS(x,y,i-1,j,c), LCS(x,y,i,j-1,c))
			}
		}
	}
	return c[i][j]
}
