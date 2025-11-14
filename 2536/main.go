func rangeAddQueries(n int, queries [][]int) (res [][]int) {
    for y := 0; y < n; y++ {        // y - col
        for x := 0; x < n; x++ {    // x - row
            if len(res) != n {
                res = append(res, []int{})
            }
            num := 0
            for _, q := range queries{
                if x >= q[0] && x <= q[2] && y >= q[1] && y <= q[3] {
                    num += 1
                }
            }
            res[x] = append(res[x], num)
        }
    }
    return res
}