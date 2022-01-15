package main

func imageSmoother(M [][]int) [][]int {
	// 构建结果数组
	res := make([][]int, len(M))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(M[0]))
	}

	// 开始计算
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[0]); j++ {
			res[i][j] = smooth(M, i, j)
		}
	}

	return res
}

func smooth(M [][]int, x int, y int) int {
	cols := len(M)
	rows := len(M[0])
	sum, num := 0, 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if x+i < 0 || x+i >= cols || y+j < 0 || y+j >= rows {
				continue
			}
			num += 1
			sum += M[x+i][y+j]
		}
	}
	return sum / num
}
