func isValidSudoku(board [][]byte) bool {
    row, col, bloc := make([]map[byte]bool,9), make([]map[byte]bool,9), make([]map[byte]bool,9)

    for i:=0; i < 9; i++ {
        row[i] = make(map[byte]bool)
        col[i] = make(map[byte]bool)
        bloc[i] = make(map[byte]bool)
    }


    for r:=0; r < 9; r++ {
        for c:=0; c < 9; c++ {
            if board[r][c] == '.' {continue}
            val := board[r][c]

            sqrid := (r/3)*3 + c/3

            if row[r][val] || col[c][val] || bloc[sqrid][val] {
                return false
            }

            row[r][val] = true
            col[c][val] = true
            bloc[sqrid][val] = true
        }
    } 
    

    return true
}
