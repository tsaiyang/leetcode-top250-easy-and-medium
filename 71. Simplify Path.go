package main

import "strings"

func simplifyPath(path string) string {
    parts := strings.Split(path, "/")
    stack := make([]string, 0, len(parts))
    for _, part := range parts {
        if len(stack) > 0 && part == ".." {
            stack = stack[:len(stack)-1]
            continue
        }
        if part != "." && part != ".." && part != "" {
            stack = append(stack, part)
        }
    }
    if len(stack) < 1 {
        return "/"
    }
    var result string
    for i := range stack {
        result += "/" + stack[i]
    }
    return result
}
