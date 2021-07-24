package main

func generateParenthesis(n int) []string {
    var result []string
    str := make([]byte, 0, 2*n)
    var fn func(level int)
    fn = func(level int) {
        if level == 2*n {
            tmp := make([]byte, len(str))
            copy(tmp, str)
            result = append(result, string(tmp))
            return
        }
        if byteNum(str, '(') < n {
            str = append(str, '(')
            fn(level + 1)
            str = str[:len(str)-1]
        }
        if byteNum(str, ')') < byteNum(str, '(') {
            str = append(str, ')')
            fn(level + 1)
            str = str[:len(str)-1]
        }
    }
    fn(0)
    return result
}

func byteNum(str []byte, char byte) int {
    count := 0
    for i := range str {
        if char == str[i] {
            count++
        }
    }
    return count
}
