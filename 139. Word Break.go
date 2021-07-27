package main

func wordBreak(s string, wordDict []string) bool {
    wordMap := make(map[string]bool)
    for _, word := range wordDict {
        wordMap[word] = true
    }
    flag := make([]bool, len(s)+1)
    flag[0] = true
    for i := 1; i < len(s)+1; i++ {
        for j := i; j >= 0; j-- {
            if flag[j] && wordMap[s[j:i]] {
                flag[i] = true
                break
            }
        }
    }
    return flag[len(flag)-1]
}
