package main

func getRoman2IntMap() map[string]int {
    m := make(map[string]int, 13)
    m["I"] = 1
    m["IV"] = 4
    m["V"] = 5
    m["IX"] = 9
    m["X"] = 10
    m["XL"] = 40
    m["L"] = 50
    m["XC"] = 90
    m["C"] = 100
    m["CD"] = 400
    m["D"] = 500
    m["CM"] = 900
    m["M"] = 1000
    return m
}

func romanToInt(s string) int {
    index := 0
    result := 0
    roman2IntMap := getRoman2IntMap()
    for index < len(s) {
        if index+1 == len(s) {
            result += roman2IntMap[s[index:index+1]]
            index++
        } else {
            if num, ok := roman2IntMap[s[index:index+2]]; ok {
                result += num
                index += 2
            } else {
                result += roman2IntMap[s[index:index+1]]
                index++
            }
        }
    }
    return result
}
