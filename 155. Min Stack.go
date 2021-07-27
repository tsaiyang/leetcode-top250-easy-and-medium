package main

import "math"

type MinStack struct {
    nums     []int
    minStack []int
}

/** initialize your data structure here. */
func Constructor2() MinStack {
    return MinStack{}
}

func (this *MinStack) Push(val int) {
    this.nums = append(this.nums, val)
    if len(this.minStack) < 1 || val < this.minStack[len(this.minStack)-1] {
        this.minStack = append(this.minStack, val)
    } else {
        this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
    }
}

func (this *MinStack) Pop() {
    if len(this.nums) < 1 {
        return
    }
    this.minStack = this.minStack[:len(this.minStack)-1]
    this.nums = this.nums[:len(this.nums)-1]
}

func (this *MinStack) Top() int {
    if len(this.nums) < 1 {
        return math.MinInt32
    }
    return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack)-1]
}
