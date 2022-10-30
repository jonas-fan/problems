// #tree #binary-tree #binary-search-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    index   int
    inorder []int
}

func inorder(node *TreeNode, out []int) []int {
    if node == nil {
        return out
    }

    out = inorder(node.Left, out)
    out = append(out, node.Val)
    out = inorder(node.Right, out)

    return out
}


func Constructor(root *TreeNode) BSTIterator {
    return BSTIterator{
        index:   0,
        inorder: inorder(root, []int{}),
    }
}


func (this *BSTIterator) Next() int {
    out := this.inorder[this.index]

    this.index++

    return out
}


func (this *BSTIterator) HasNext() bool {
    return this.index < len(this.inorder)
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
