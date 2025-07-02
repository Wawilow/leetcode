word = "aaabbb"
res = 1
print(enumerate(word[1:]))
for i,c in enumerate(word[1:]):
    print(i, c, word[i])
    if c==word[i]: res+=1
print(res)
