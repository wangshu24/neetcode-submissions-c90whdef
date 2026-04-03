func searchMatrix(matrix [][]int, target int) bool {
    l, r := 0, len(matrix[0])
    for _, arr := range matrix {
        if target <= arr[len(arr)-1] && target >= arr[0] {
            for l <= r {
                m := l + (r - l)/2
                if arr[m] == target {
                    return true
                } else {
                    if arr[m] < target {
                        l = m + 1
                    } else {
                        r = m - 1
                    }
                }
            }
        }
    }

    return false
}
