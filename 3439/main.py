class Solution:
    def maxFreeTime(self, eventTime: int, k: int, startTime: list[int], endTime: list[int]) -> int:
        b_val = self.maxGap(eventTime, startTime, endTime)
        if k == 0:
            return b_val
        # print(b_val, startTime, endTime, len(startTime))
        k -= 1
        for v in range(len(startTime)):
            V2Flag = True
            V3Flag = True

            startTimeV2 = startTime[:]
            endTimeV2 = endTime[:]
            startTimeV3 = startTime[:]
            endTimeV3 = endTime[:]

            delta = endTime[v] - startTime[v]

            if v == 0:
                # already near the wall, just skip it
                V2Flag = False
            else:
                startTimeV2[v] = startTimeV2[v-1]+1
                endTimeV2[v] = startTimeV2[v]+delta

            if v == len(startTime)-1:
                # already near the wall, just skip it
                V3Flag = False
            else:
                endTimeV3[v] = startTime[v+1]
                startTimeV3[v] = endTimeV3[v]-delta

            if V2Flag:
                b_valV2 = self.maxFreeTime(eventTime, k, startTimeV2, endTimeV2)
                if b_valV2 > b_val:
                    b_val = b_valV2

            if V3Flag:
                b_valV3 = self.maxFreeTime(eventTime, k, startTimeV3, endTimeV3)
                if b_valV3 > b_val:
                    b_val = b_valV3
            # print("v: ", b_val, v, V3Flag, V2Flag)
        return b_val
    
    def maxGap(self, eventTime, startTime, endTime) -> int:
        res = 0
        for i in range(len(startTime)):
            if i == 0:
                if res < startTime[i]:
                    res = startTime[i]
            elif i == len(startTime) - 1:
                if res < eventTime - endTime[i]:
                    res = eventTime - endTime[i]
            else:
                if res < startTime[i] - endTime[i-1]:
                    res = startTime[i] - endTime[i-1]
                elif res > startTime[i] - endTime[i+1]:
                    res = startTime[i] - endTime[i+1]
        return res

    def maxGapIDK(self, eventTime, startTime, endTime):
        n = len(startTime)
        # res = []
        res_val = 0

        last_event_end = 0
        for i in range(n):
            eeT = 0
            if i!=0:
                eeT = endTime[i-1]
            ssT = eventTime
            if i!=n-1:
                ssT = startTime[i+1]

            eT = endTime[i]
            sT = startTime[i]

            if (sT-eeT) + (ssT - eT) > res_val:
                res_val = (sT-eeT)+(ssT-eT)
                print(res_val, sT, eeT, ssT, eT)
                # res = [i]
            # elif (sT-eeT)+(ssT-eT) == res_val:
                # res.append(i)
        # return (res, res_val)
        return None, res_val

s = Solution()
print(s.maxFreeTime(5, 1, [1,3],[2,5]), "\n")
print(s.maxFreeTime(11, 1, [0, 2, 9], [1, 4, 10]), "\n")
# print(s.maxFreeTime(11, 2, [0, 2, 9], [1, 4, 10]))
