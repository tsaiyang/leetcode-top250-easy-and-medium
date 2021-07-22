package main

func getDigitMap() map[byte]string {
    m := make(map[byte]string)
    m['2'] = "abc"
    m['3'] = "def"
    m['4'] = "ghi"
    m['5'] = "jkl"
    m['6'] = "mno"
    m['7'] = "pqrs"
    m['8'] = "tuv"
    m['9'] = "wxyz"
    return m
}

func letterCombinations(digits string) []string {
    digitMap := getDigitMap()
    result := make([]string, 0)
    str := make([]byte, 0)
    var fn func(level int)
    fn = func(level int) {
        if level == len(digits) {
            tmp := make([]byte, len(str))
            copy(tmp, str)
            result = append(result, string(tmp))
            return
        }
        for i := range digitMap[digits[level]] {
            str = append(str, digitMap[digits[level]][i])
            fn(level + 1)
            str = str[:len(str)-1]
        }
    }
    fn(0)
    return result
}
