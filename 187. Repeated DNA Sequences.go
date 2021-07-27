package main

func findRepeatedDnaSequences(s string) []string {
    repeatedSet := make(map[string]bool)
    set := make(map[string]bool)
    for i := 0; i < len(s)-9; i++ {
        str := s[i : i+10]
        if set[str] {
            repeatedSet[str] = true
        }
        set[str] = true
    }
    var result []string
    for str, _ := range repeatedSet {
        result = append(result,str)
    }
    return result
}
