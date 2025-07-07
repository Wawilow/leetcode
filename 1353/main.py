class Solution:
    def maxEvents(self, events: list[list[int]]) -> int:
        heatmap = {}
        eventmap = {}
        event_num = 0
        for event in events:
            for i in range(event[0], event[1]+1):
                if heatmap.get(i) == None:
                    heatmap[i] = [event_num]
                else:
                    heatmap[i].append(event_num)
            if eventmap.get(event_num) == None:
                eventmap[event_num] = event
            event_num += 1
        del event_num, i, event

        heatmap = dict(sorted(heatmap.items()))

        res = 0
        blacklist = []
        for i in heatmap.values():
            closest_end = -1
            closest_event = None
            print(f"events for the day {i}")
            for e in i:
                # print(e, eventmap[e])
                if e not in blacklist:
                    if closest_end == -1 or closest_end > eventmap[e][1]:
                        print(f"e: {e} {eventmap[e]}")
                        closest_end = eventmap[e][1]
                        closest_event = e
            if closest_event != None:
                print(f"attending to event: {closest_event}")
                blacklist.append(closest_event)
                res += 1
        return res


s = Solution()
# print(s.maxEvents([[1,3],[2,3],[3,4]]))
# print(s.maxEvents([[1,2],[2,3],[3,4],[1,2]]))
print(s.maxEvents([[26,27],[24,26],[1,4],[1,2],[3,6],[2,6],[9,13],[6,6],[25,27],[22,25],[20,24],[8,8],[27,27],[30,32]]))
