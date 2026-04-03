func exist(board [][]byte, word string) bool {
	var found bool
	visit := make(map[string]struct{})

	var back func( map[string]struct{}, int, int, int) bool
	back = func( vis map[string]struct{}, x, y, cur int) bool {
		if cur == len(word) {
			return true
		} else if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) {return false}
	
		key := strconv.Itoa(x) + "-" + strconv.Itoa(y)
		_,there := vis[key]
		if board[x][y] != word[cur] || there {return false}

		vis[key] = struct{}{}
		found := back(vis, x+1,y,cur+1) || back(vis, x-1,y, cur+1) || back(vis, x, y+1, cur+1) || back(vis, x, y-1,cur+1)
		delete(vis, key)
		return found
	}
	
	for i:=0; i < len(board); i++  {
		for j:=0; j < len(board[0]); j++ {
			if back(visit,i,j,0) {return true}
		}
	}
	return found
}