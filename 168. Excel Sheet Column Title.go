package main

func convertToTitle(columnNumber int) string {
    var result string
    for columnNumber > 0 {
        result += string('A' + columnNumber/26)
        columnNumber /= 26
    }
    return result
}
