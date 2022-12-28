import "math"

func maxWidthOfVerticalArea(points [][]int) int {
    xAxis := make([]int, len(points))

    for i := 0; i < len(points); i++ {
        xAxis[i] = points[i][0]
    }

    sort.Ints(xAxis)
    var max float64 = 0

    for i := 0; i < len(xAxis) - 1; i++ {
        res := math.Abs(float64(xAxis[i] - xAxis[i + 1]))
        if res > max {
            max = res
        }
    }

    return int(max)
}
