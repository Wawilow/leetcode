class FindSumPairs:

    def __init__(self, nums1: list[int], nums2: list[int]):
        self.nums1 = nums1
        self.nums2 = nums2

    def add(self, index: int, val: int) -> None:
        # Add a positive integer to an element of a given index in the array nums2.
        self.nums2[index] += val

    def count(self, tot: int) -> int:
        res = 0
        for i in self.nums1:
            for j in self.nums2:
                if i + j == tot:
                    res += 1
        return res
        # print(res, self.nums1, self.nums2) 


# Your FindSumPairs object will be instantiated and called as such:
# obj = FindSumPairs(nums1, nums2)
# obj.add(index,val)
# param_2 = obj.count(tot)

s = FindSumPairs(nums1=[1, 1, 2, 2, 2, 3], nums2=[1, 4, 5, 2, 5, 4])
s.count(7)
s.add(3, 2)
s.count(8)
s.count(4)
s.add(0, 1)
s.add(1, 1)
s.count(7)
