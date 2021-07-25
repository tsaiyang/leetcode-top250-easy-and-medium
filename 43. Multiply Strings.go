package main

import "strconv"

//func multiply(num1 string, num2 string) string {
//   var fn func(level int) string
//   fn = func(level int) string {
//       if level == -1 {
//           return ""
//       }
//
//   }
//}

func addZeros(n int) string {
    var result string
    for i := 0; i < n; i++ {
        result += string('0')
    }
    return result
}

func reverseString(s string, i, j int) string {
    bytes := []byte(s)
    for i < j {
        bytes[i], bytes[j] = bytes[j], bytes[i]
        i++
        j--
    }
    return string(bytes)
}

func addStr(s1, s2 string) string {
    s1 = reverseString(s1, 0, len(s1)-1)
    s2 = reverseString(s2, 0, len(s2)-1)
    n1, n2, flag := 0, 0, 0
    index1, index2 := 0, 0
    var result string
    for index1 < len(s1) || index2 < len(s2) || flag != 0 {
        n1 = 0
        if index1 < len(s1) {
            n1 = int(s1[index1] - '0')
            index1++
        }

        n2 = 0
        if index2 < len(s2) {
            n2 = int(s2[index2] - '0')
            index2++
        }

        result += strconv.Itoa((n1 + n2 + flag) % 10)
        flag = (n1 + n2 + flag) / 10
    }
    return reverseString(result, 0, len(result)-1)
}
