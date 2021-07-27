package main

func cloneGraph(node *GraphNode) *GraphNode {
    nodeMap := make(map[*GraphNode]*GraphNode)
    queue := []*GraphNode{node}
    nodeMap[node] = &GraphNode{Val: node.Val}
    for len(queue) > 0 {
        n := queue[0]
        queue = queue[1:]
        for _, neighbor := range n.Neighbors {
            if _, ok := nodeMap[neighbor]; !ok {
                queue = append(queue, neighbor)
                nodeMap[neighbor] = &GraphNode{Val: neighbor.Val}
            }
            nodeMap[n].Neighbors = append(nodeMap[n].Neighbors, nodeMap[neighbor])
        }
    }
    return nodeMap[node]
}
