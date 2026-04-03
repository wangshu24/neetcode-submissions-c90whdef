class Solution {
    /**
     * @param {character[][]} board
     * @param {string} word
     * @return {boolean}
     */
    exist(board, word) {
        const ROW = board.length
        const COL = board[0].length
        let path = new Set()

        for (let x = 0; x <  ROW; x++) {
            for (let y = 0; y < COL; y++) {
                if (this.backtrack(board, x,y,0,ROW,COL,path,word)) {return true}
            }
        }

        return false
    }

    backtrack(board,x,y,i,ROW, COL,path,word) {
        // console.log(x,y, word, path, i === word.length)
        if (i === word.length) {return true}
        if (x < 0 || y < 0 || x >= ROW || y >= COL || 
        board[x][y] !== word[i] || path.has(`${x},${y}`)) {return false}

        path.add(`${x},${y}`)
        const res = this.backtrack(board, x+1,y,i+1,ROW,COL,path,word) ||
            this.backtrack(board, x-1,y,i+1,ROW,COL,path,word) ||
            this.backtrack(board, x,y+1,i+1,ROW,COL,path,word) ||
            this.backtrack(board,x,y-1,i+1,ROW,COL,path,word)
        path.delete(`${x},${y}`)
        return res
    }
}
