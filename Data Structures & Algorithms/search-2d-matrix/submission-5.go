func searchMatrix(matrix [][]int, target int) bool {
    left := 0
    n := len(matrix)
    right := n -1

    for left <= right {
        middle := left + (right - left) / 2
        middle_len := len(matrix[middle]) -1
        if target >= matrix[middle][0] && target <= matrix[middle][middle_len] {
            ileft := 0
            iright := middle_len

            for ileft <= iright {
                imiddle := ileft + (iright - ileft) /2

                if target == matrix[middle][imiddle] {
                    return true
                }

                if target > matrix[middle][imiddle] {
                    ileft = imiddle + 1
                }else {
                    iright = imiddle -1
                }
               
            }
            return false
        }
        if target > matrix[middle][0] {
            left = middle +1
        }
        if target < matrix[middle][0] {
            right = middle-1
        }

    }

return false
}
