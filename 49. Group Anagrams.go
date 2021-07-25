package main

import "sort"

func groupAnagrams(strs []string) [][]string {
    strMap := make(map[string][]string)
    for _, str := range strs {
        bytes := []byte(str)
        sort.Sort(ByteSlice(bytes))
        if _, ok := strMap[string(bytes)]; !ok {
            strMap[string(bytes)] = []string{str}
        } else {
            strMap[string(bytes)] = append(strMap[string(bytes)], str)
        }
    }
    var result [][]string
    for _, v := range strMap {
        result = append(result, v)
    }
    return result
}

type ByteSlice []byte

func (bs ByteSlice) Len() int {
    return len(bs)
}

func (bs ByteSlice) Less(i, j int) bool {
    return bs[i] < bs[j]
}

func (bs ByteSlice) Swap(i, j int) {
    bs[i], bs[j] = bs[j], bs[i]
}
