# 6 lines, 3 couners, 1 loop

1. The best solutions is always went left -> right
2. Idc how many 0 are there in the sequence 10 == 10000000000000
3. Idc about 1 on the far right end 10111111111 == 10

So moving from left to right I count how many ones do I have on the far left side and how many new ones i've met.

## Code
```golang []
func maxOperations(s string) (res int) {
    counter := 0
    leftOnes := 0
    for _, j := range s {if j == 49 {counter += 1} else if counter != 0{res += leftOnes + counter
            leftOnes += counter
            counter = 0}}
    return res
}
```