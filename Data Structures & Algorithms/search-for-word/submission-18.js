class Solution {
    /**
     * @param {character[][]} board
     * @param {string} word
     * @return {boolean}
     */
    exist(board, word) {
        const ROW = board.length
        const COL = board[0].length
        const visit = new Set()
        
        for (let x = 0; x < ROW; x++) {
            for (let y = 0; y < COL; y++) {
                if (this.search(board, x, y, 0, ROW, COL, visit, word)) {return true}
            }
        }

        return false
    }

    search(board,x,y,i, ROW, COL, visit, word) {
            if (i === word.length) {
                return true
            } else if (x < 0 || y < 0 ||
            x >= ROW || y >= COL || visit.has(`${x},${y}`) || board[x][y] != word[i]) {
                return false
            }
            console.log(visit)
            
            visit.add(`${x},${y}`)

            const res = this.search(board, x+1,y,i+1, ROW, COL, visit, word) || 
            this.search(board, x-1,y,i+1,ROW, COL, visit, word) || 
            this.search(board, x, y+1,i+1, ROW, COL, visit, word) || 
            this.search(board, x, y-1,i+1, ROW, COL, visit, word)
            
            visit.delete(`${x},${y}`)

            return res
        }
}
 