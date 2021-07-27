package main

func isIsomorphic(s string, t string) bool {
    sMap := make(map[byte]byte, 256)
    tMap := make(map[byte]byte, 256)
    for i := range s {
        _, ok1 := sMap[s[i]]
        _, ok2 := tMap[t[i]]
        if !ok1 && !ok2 {
            sMap[s[i]] = t[i]
            tMap[t[i]] = s[i]
            continue
        }
        if !ok1 || !ok2{
            return false
        }
    }
    return true
}
