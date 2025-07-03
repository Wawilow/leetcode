class Solution:
    def generate_string(self) -> str:
        def next_alpha(s):
            return chr((ord(s.upper())+1 - 65) % 26 + 65).lower()
        
        val = "a"
        for i in range(9): # 2 ^ 9 -> 512 characters
            val += "".join([next_alpha(i) for i in val])
        return val

    def kthCharacter(self, k: int) -> str:
        return "abbcbccdbccdcddebccdcddecddedeefbccdcddecddedeefcddedeefdeefeffgbccdcddecddedeefcddedeefdeefeffgcddedeefdeefeffgdeefeffgeffgfgghbccdcddecddedeefcddedeefdeefeffgcddedeefdeefeffgdeefeffgeffgfgghcddedeefdeefeffgdeefeffgeffgfgghdeefeffgeffgfggheffgfgghfgghghhibccdcddecddedeefcddedeefdeefeffgcddedeefdeefeffgdeefeffgeffgfgghcddedeefdeefeffgdeefeffgeffgfgghdeefeffgeffgfggheffgfgghfgghghhicddedeefdeefeffgdeefeffgeffgfgghdeefeffgeffgfggheffgfgghfgghghhideefeffgeffgfggheffgfgghfgghghhieffgfgghfgghghhifgghghhighhihiij"[k-1] 


        # iteration_num = 0
        # itr_i = 0
        # while iteration_num == 0:
        #     if (2 ** itr_i) > k:
        #         iteration_num = itr_i
        #     itr_i += 1
        # del itr_i
        #
        # print(k, iteration_num)
    
        # a -> ab -> ab bc -> abbc bccd -> a b bc bccd bccdcdde
        # 1 -> 2 -> 4 -> 8

s = Solution()
# s.kthCharacter(1)
# s.kthCharacter(2)
# s.kthCharacter(3)
# s.kthCharacter(4)
print(s.kthCharacter(5))
