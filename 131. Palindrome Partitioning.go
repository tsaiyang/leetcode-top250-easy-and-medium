package main

func partition1(s string) [][]string {
    palindrome := make([][]bool, len(s))
    for i := range palindrome {
        palindrome[i] = make([]bool, len(s))
    }
    for i := range s {
        for j := i; j >= 0; i-- {
            if s[i] == s[j] && (j-i < 2 || palindrome[i+1][j-1]) {
                palindrome[i][j] = true
            }
        }
    }
    var result [][]string
    var strs []string
    var fn func(level int)
    fn = func(level int) {
        if level == len(s) {
            tmp := make([]string, len(strs))
            copy(tmp, strs)
            result = append(result, tmp)
            return
        }
        for i := level; i < len(s); i++ {
            if palindrome[level][i] {
                strs = append(strs, s[level:i+1])
                fn(i + 1)
                strs = strs[:len(strs)-1]
            }
        }
    }
    fn(0)
    return result
}

func isPalindrome2(s string, i, j int) bool {
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}
