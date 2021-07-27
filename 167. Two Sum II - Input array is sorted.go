package main

func twoSum1(numbers []int, target int) []int {
    if numbers[len(numbers)-1]+numbers[len(numbers)-2] < target {
        return []int{numbers[len(numbers)-1], len(numbers)}
    }
    i := 0
    j := len(numbers) - 1
    for i < j {
        if target == numbers[i]+numbers[j] {
            return []int{i, j}
        }
        if target > numbers[i]+numbers[j] {
            i++
        } else {
            j--
        }
    }
    return []int{-1, -1}
}
