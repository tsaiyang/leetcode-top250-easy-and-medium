package main

func addBinary(a string, b string) string {
    result := make([]byte, 0, len(a)+len(b))
    a1 := reverseStr(a)
    b1 := reverseStr(b)
    i1 := 0
    i2 := 0

    var n1 byte = '0'
    var n2 byte = '0'
    var flag byte = '0'
    for i1 < len(a1) || i2 < len(b1) || flag == '1' {
        if i1 >= len(a1) {
            n1 = '0'
        } else {
            n1 = a1[i1]
            i1++
        }

        if i2 >= len(b1) {
            n2 = '0'
        } else {
            n2 = b1[i2]
            i2++
        }

        if n1 == '1' && n2 == '1' {
            result = append(result, flag)
            flag = '1'
        } else if n1 == '1' || n2 == '1' {
            if flag == '1' {
                result = append(result, '0')
            } else {
                result = append(result, '1')
                flag = '0'
            }
        } else {
            if flag == '1' {
                result = append(result, '1')
            } else {
                result = append(result, '0')
            }
            flag = '0'
        }
    }
    return reverseStr(string(result))
}

func reverseStr(s string) string {
    bytes := []byte(s)
    i := 0
    j := len(bytes) - 1
    for i < j {
        bytes[i], bytes[j] = bytes[j], bytes[i]
    }
    return string(bytes)
}
