class Solution:
    def possibleStringCount(self, word: str, k: int) -> int:
        if len(word) == k:
            return 1
        elif len(word) < k:
            return 0

        res = 1
        raw_list = []
        last_letter = ""
        raw_list_last_val = ""
        for i in word:
            if last_letter == "":
                last_letter = i
                raw_list_last_val = i
            elif i == last_letter:
                raw_list_last_val += i
            else:
                if last_letter != "":
                    raw_list.append(len(raw_list_last_val))
                    raw_list_last_val = ""
                    last_letter = ""
                last_letter = i
                raw_list_last_val = i
        if last_letter != "":
            raw_list.append(len(raw_list_last_val))
        del raw_list_last_val, last_letter, i
   

        word_sum = sum(raw_list)
        sublist = list()
        sublist.append(({"v": raw_list, "k": [False for _ in raw_list]}))

        for _ in range(len(raw_list)):
            tmp_sublist = sublist.copy()
            for sublist_elem in tmp_sublist:
                any_changes = False
                for j in range(len(sublist_elem["k"])):
                    if sublist_elem["k"][j] == False:
                        any_changes = True
                        new_k = sublist_elem["k"].copy()
                        new_k[j] = True
                        for ii in range(sublist_elem["v"][j]):
                            sublist_val = sublist_elem["v"].copy()
                            sublist_val[j] = ii + 1
                            if sum(sublist_val) >=k:
                                sublist.append({"v": sublist_val, "k": new_k})
                            del sublist_val
                        del ii, new_k,

                if any_changes:
                    sublist.remove(sublist_elem)
                del j, any_changes
            del tmp_sublist

        unique_sublist = set([str(i) for i in sublist if all(i["k"])])

        return len(unique_sublist)


s = Solution()
# print(s.possibleStringCount("aabbccdd", 7))
# print(s.possibleStringCount("aabbccdd", 8))
# print(s.possibleStringCount("aaabbb", 3))
print(s.possibleStringCount("aan", 2))
# print(s.possibleStringCount("aakf", 3))
# print(s.possibleStringCount("", ))
# print(s.possibleStringCount("", ))
# print(s.possibleStringCount("", ))
# print(s.possibleStringCount("", ))
# print(s.possibleStringCount("", ))
# print(s.possibleStringCount("", ))
# print(s.possibleStringCount("", ))
# print(s.possibleStringCount("", ))

