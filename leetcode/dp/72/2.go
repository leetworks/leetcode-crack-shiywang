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

var dp[1000][1000]int


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
    
    if dp[w1][w2] < max {
        return dp[w1][w2]
    }
    
    if word1[w1] == word2[w2] {
        a = slove(w1+1, w2+1, word1, word2)
    }
    
    b = slove(w1+1, w2, word1, word2) + 1
    c = slove(w1, w2+1, word1, word2) + 1
    d = slove(w1+1, w2+1, word1, word2) + 1
    
    dp[w1][w2] = Min(a,b,c,d)
    return dp[w1][w2]
}



func minDistance(word1 string, word2 string) int {
    for i := 0; i < 1000; i++ {
        for j:=0; j < 1000;j++ {
            dp[i][j]= max
        }
    }
    return slove(0,0,word1,word2)
}
