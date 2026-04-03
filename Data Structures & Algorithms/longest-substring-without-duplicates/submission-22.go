func lengthOfLongestSubstring(s string) int {
    visited := make(map[byte]struct{})
    tail := 0
    maxLen := 0

    for i := 0; i < len(s); i++ {

        // 🔥 shrink window while duplicate exists
        for {
            if _, exists := visited[s[i]]; !exists {
                break
            }
            delete(visited, s[tail])
            tail++
        }

        visited[s[i]] = struct{}{}
        maxLen = max(maxLen, i-tail+1)
    }

    return maxLen
}
