package backtracking

func solveSudoku(board [][]byte) {
	n := len(board)
	m := make([]bool, n*n*3)
	a, b, count := -1, -1, 0
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if board[x][y] != '.' {
				m[n*x+int(board[x][y]-'1')] = true                 // 行
				m[81+n*y+int(board[x][y]-'1')] = true              // 列
				m[162+n*((x/3)*3+y/3)+int(board[x][y]-'1')] = true // 格
			} else {
				a = x
				b = y
				count++
			}
		}
	}
	stack, top := make([]int, count*3), -1
	try(board, m, 0, 0, 0, n, stack, top, a, b)
}

func try(board [][]byte, m []bool, x, y, z, n int, stack []int, top, a, b int) {
	for ; x < n; x++ {
		for ; y < n; y++ {
			if board[x][y] == '.' {
				for ; z < n; z++ {
					if m[n*x+z] || m[81+n*y+z] || m[162+n*((x/3)*3+y/3)+z] {
						continue
					} else {
						board[x][y] = byte(z + '1')
						m[n*x+z] = true
						m[81+n*y+z] = true
						m[162+n*((x/3)*3+y/3)+z] = true
						if z < n {
							top++
							stack[top*3+0] = x
							stack[top*3+1] = y
							stack[top*3+2] = z
						}
						if x == a && y == b {
							correct := true
							for _, item := range m {
								if !item {
									correct = false
									break
								}
							}
							if !correct {
								if top >= 0 && stack[top*3+2] < n {
									try(board, m, stack[top*3+0], stack[top*3+1], stack[top*3+2]+1, n, stack, top-1, a, b)
								}
							}
						}
						break
					}
				}
				z = 0
			}
		}
		y = 0
	}
}
