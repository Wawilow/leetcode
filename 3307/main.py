def next_alpha(s):
    # a -> z = 97 -> 122
    return chr(((ord(s)+1-97) % 26) + 97)

class Solution:
    def kthCharacter(self, k: int, operations: list[int]) -> str:
        word = "a"
        for op in operations:
            if op == 0:
                word += word
            else:
                word += "".join([next_alpha(i) for i in word])
           
            if len(word) >= k:
                return word[k-1]
        return word[k-1]

s = Solution()
print(s.kthCharacter(10, [0,1,0,1]))
