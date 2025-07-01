# class Solution:
#     def possibleStringCount(self, word: str) -> int:
#         raw_list = []
#         res = 1
#
#         last_letter = ""
#         current_seqence = ""
#         for i in word:
#             if i == last_letter:
#                 current_seqence += i
#             else:
#                 if current_seqence != "":
#                     if len(current_seqence) > 1:
#                         res += len(current_seqence) - 1
#                     raw_list.append(current_seqence)
#                 current_seqence = i
#             last_letter = i
#         raw_list.append(current_seqence)
#
#         if len(current_seqence) > 1:
#             res += len(current_seqence) - 1
#
#         del last_letter, current_seqence, i
#
#         print(raw_list)
#
#         # res = 0
#         # for i in raw_list:
#         #     if len(i) != 1:
#         #         res += len(i)
#         # if res != 0:
#         #     res -= 1
#         return res


class Solution:
    def possibleStringCount(self, word: str) -> int:
        raw_list = []
        res = 1

        last_letter = ""
        current_seqence = ""
        
        for i in word:
            if i == last_letter:
                current_seqence += i
            else:
                if current_seqence != "":
                    if len(current_seqence) > 1:
                        res += len(current_seqence) - 1
                    raw_list.append(current_seqence)
                current_seqence = i
            last_letter = i
        raw_list.append(current_seqence)
        if len(current_seqence) > 1:
            res += len(current_seqence) - 1

        del last_letter, current_seqence, i
        return res



s = Solution()
print(s.possibleStringCount("abbcccc"))
print(s.possibleStringCount("abcd"))
print(s.possibleStringCount("all"))
print(s.possibleStringCount("allbb"))
# print(s.possibleStringCount("abbcccc"))
