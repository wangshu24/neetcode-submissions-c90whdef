func solveNQueens(n int) [][]string {
	res := [][]string{}

	board := make([][]byte, n)
	
	for i:=0;i < n; i++ {
		board[i] = make([]byte, n)
		for j:=0; j < n; j++ {
			board[i][j] = '.'
		}	
	}

	var back func(int, [][]byte)
	back = func(row int, cboard [][]byte) {
		if row == n  {
			add := []string{}
			for i:=0; i < n; i++ {
				add = append(add, string(cboard[i]))
			}
			res = append(res, add)
			return
		}
		
		for i:=0; i < n; i++ {
			if valid(n, row, i, cboard) {
				cboard[row][i] = 'Q'
				back(row+1, cboard)
				cboard[row][i] = '.'
			}
		} 
	}


	
	back(0, board)
	

	return res
}

func valid(n, row, col int, board [][]byte) bool {
	for i:=0; i < n; i++ {
		if board[i][col] == 'Q' {return false}
	}
	x, y := row, col
	for x >=0 && y >=0 {
		if board[x][y] == 'Q' {return false}
		x--
		y--
	}
	x, y = row, col
	for x >=0 && y <n {
		if board[x][y] == 'Q' {return false}
		x--
		y++
	}
	x, y = row, col
	for x <n && y <n {
		if board[x][y] == 'Q' {return false}
		x++
		y++
	}
	x, y = row, col
	for x <n && y >=0 {
		if board[x][y] == 'Q' {return false}
		x++
		y--
	}
	return true
}