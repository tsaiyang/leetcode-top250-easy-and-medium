package main

func getInt2RomanMap() map[int]string {
    m := make(map[int]string, 13)
    m[1] = "I"
    m[4] = "IV"
    m[5] = "V"
    m[9] = "IX"
    m[10] = "X"
    m[40] = "XL"
    m[50] = "L"
    m[90] = "XC"
    m[100] = "C"
    m[400] = "CD"
    m[500] = "D"
    m[900] = "CM"
    m[1000] = "M"
    return m
}

func intToRoman(num int) string {
    int2RomanMap := getInt2RomanMap()
    arr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    result := ""
    for num > 0 {
        for i := range arr {
            if num >= arr[i] {
                result += int2RomanMap[arr[i]]
                num -= arr[i]
                break
            }
        }
    }
    return result
}
