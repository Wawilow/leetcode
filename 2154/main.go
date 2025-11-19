func findFinalValue(nums []int, original int) int {
    m := make(map[int]bool)
    for _, v := range nums {
        if v % original == 0 {
            m[v] = true
        }
    }
    loop := true
    for loop {
        if m[original] {
            original *= 2
        } else {
            loop = false
        }
    }
    return original
}

