package main

import "math"

type BSTIterator struct {
    nodes []*TreeNode
    index int
}

func Constructor1(root *TreeNode) BSTIterator {
    nodes := []*TreeNode{&TreeNode{Val: math.MinInt32}}
    var stack []*TreeNode
    for len(stack) > 0 || root != nil {
        if root != nil {
            stack = append(stack, root)
            root = root.Left
        } else {
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            nodes = append(nodes, root)
            root = root.Right
        }
    }
    return BSTIterator{nodes: nodes, index: 0}
}

func (this *BSTIterator) Next() int {
    this.index++
    return this.nodes[this.index].Val
}

func (this *BSTIterator) HasNext() bool {
    if len(this.nodes) < 2 || this.index == len(this.nodes) {
        return false
    }
    return true
}
