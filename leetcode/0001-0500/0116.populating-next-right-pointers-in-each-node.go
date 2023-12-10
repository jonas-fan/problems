// #tree #binary-tree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func walk(node *Node, right *Node) *Node {
    if node == nil {
        return nil
    }

    node.Next = right
    node.Left = walk(node.Left, node.Right)

    if node.Next == nil {
        node.Right = walk(node.Right, node.Next)
    } else {
        node.Right = walk(node.Right, node.Next.Left)
    }

    return node
}

func connect(root *Node) *Node {
    return walk(root, nil)
}
