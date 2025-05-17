package main
import (
    "fmt"
)

func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return []int{}
    }

    rows := len(matrix)
    cols := len(matrix[0])
    result := make([]int, 0, rows*cols)

    top, bottom := 0, rows-1
    left, right := 0, cols-1

    for left <= right && top <= bottom {
        // dari kiri ke kanan
        for col := left; col <= right; col++ {
            result = append(result, matrix[top][col])
        }
        top++

        // dari atas ke bawah
        for row := top; row <= bottom; row++ {
            result = append(result, matrix[row][right])
        }
        right--

        if top <= bottom {
            // dari kanan ke kiri
            for col := right; col >= left; col-- {
                result = append(result, matrix[bottom][col])
            }
            bottom--
        }

        if left <= right {
            // dari bawah ke atas
            for row := bottom; row >= top; row-- {
                result = append(result, matrix[row][left])
            }
            left++
        }
    }

    return result
}
func main() {
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    result := spiralOrder(matrix)
    fmt.Println("Hasil spiral order:")
    fmt.Println(result)
}
