var max = 9999999999
func Min2(a,b int) int {
    if a < b {
        return a
    }
    return b
}

func Min(a,b,c,d int) int {
    return Min2(d,Min2(a,Min2(b,c)))
}

func slove(w1, w2 int, word1, word2 string) int {
    a := max
    b := max
    c := max
    d := max
    if w1 == len(word1) && w2 == len(word2) {
        return 0
    }
    if w1 >= len(word1){
        return len(word2) - w2
    }
    if w2 >= len(word2) {
        return len(word1) - w1
    }

    if word1[w1] == word2[w2] {
        a = slove(w1+1, w2+1, word1, word2)
    }
    b = slove(w1+1, w2, word1, word2) + 1
    c = slove(w1, w2+1, word1, word2) + 1
    d = slove(w1+1, w2+1, word1, word2) + 1
    return Min(a,b,c,d)
}



func minDistance(word1 string, word2 string) int {
    return slove(0,0,word1,word2)
}
