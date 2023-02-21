import "fmt"

func oddCells(m int, n int, indices [][]int) int {
    initialMatrix := make([][]int, m)

    for i := 0; i < m; i++ {
        initialMatrix[i] = make([]int, n)
    }

    for i := 0; i < len(indices); i++ {
        r := indices[i][0] 
        c := indices[i][1]
        
        for j := 0; j < n; j++ {
            initialMatrix[r][j] += 1
        }

        for j := 0; j < m; j++ {
            initialMatrix[j][c] += 1
        }
    }

    counter := 0

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if initialMatrix[i][j] % 2 == 1 {
                counter++
            }
        }
    }

    return counter
}
