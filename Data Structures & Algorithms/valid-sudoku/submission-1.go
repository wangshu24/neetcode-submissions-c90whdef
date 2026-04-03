func isValidSudoku(board [][]byte) bool {
    for col:=0; col < 9; col++ {
        found := make(map[byte]bool)
        for i:=0; i < 9; i++ {
            if board[i][col] == '.' {
                continue
            }
            if found[board[i][col]] {
                return false
            } else {
                found[board[i][col]] = true
            }
            
        }
    }

    for row:=0; row < 9; row++ {
        found := make(map[byte]bool)
        for i:=0; i < 9; i++ {
            if board[row][i] == '.' {
                continue
            }
            if found[board[row][i]]  {
                return false
            }  else {
                found[board[row][i]] = true
            }
        }
    }

    for sq:=0; sq < 9; sq++ {
        seen := make(map[byte]bool)

        for i:=0; i < 3; i++ {
            for j:=0; j < 3; j++ {
                row := (sq / 3) * 3 + i
                col := (sq % 3) * 3 + j 
                if board[row][col] == '.' {continue}
                if seen[board[row][col]] { 
                    return false
                } else {
                    seen[board[row][col]] = true
                }
            }
        }
    }

    return true    
}
