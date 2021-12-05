package tree

import (
    "fmt"
    "likebeta/gds/util"
    "strings"
)

type BSTNode struct {
    Left  *BSTNode
    Right *BSTNode
    Value interface{}
}

func (node *BSTNode) String() string {
    return fmt.Sprintf("%v", node.Value)
}

type BSTree struct {
    Root       *BSTNode
    Comparator util.Comparator
    size       int
}

func NewBSTree(comparator util.Comparator) *BSTree {
    return &BSTree{Comparator: comparator}
}

func (t *BSTree) IsValid() bool {
    return t.isValid(t.Root, nil, nil)
}

func (t *BSTree) isValid(root *BSTNode, min *BSTNode, max *BSTNode) bool {
    if root == nil {
        return true
    }
    if min != nil && t.Comparator(root.Value, min.Value) <= 0 {
        return false
    } else if max != nil && t.Comparator(root.Value, max.Value) >= 0 {
        return false
    }
    return t.isValid(root.Left, min, root) && t.isValid(root.Right, root, max)
}

func (t *BSTree) Size() int {
    return t.size
}

func (t *BSTree) Add(value interface{}) *BSTree {
    // t.Root = t.addWithRecursion(t.Root, value)
    // return t
    if t.Root == nil {
        t.Root = &BSTNode{Value: value}
    } else {
        pre, curr := t.Root, t.Root
        var v int
        for curr != nil {
            v = t.Comparator(curr.Value, value)
            if v == 0 {
                return t
            } else if v > 0 {
                pre, curr = curr, curr.Left
            } else {
                pre, curr = curr, curr.Right
            }
        }
        if v > 0 {
            pre.Left = &BSTNode{Value: value}
        } else {
            pre.Right = &BSTNode{Value: value}
        }
    }
    t.size++
    return t
}

func (t *BSTree) addWithRecursion(root *BSTNode, value interface{}) *BSTNode {
    if root == nil {
        t.size++
        return &BSTNode{Value: value}
    }

    v := t.Comparator(root.Value, value)
    if v > 0 {
        root.Left = t.addWithRecursion(root.Left, value)
    } else if v < 0 {
        root.Right = t.addWithRecursion(root.Right, value)
    }
    return root
}

func (t *BSTree) Find(value interface{}) bool {
    // node := t.findWithRecursion(t.Root, value)
    // return node != nil
    curr := t.Root
    for curr != nil {
        v := t.Comparator(curr.Value, value)
        if v == 0 {
            return true
        } else if v > 0 {
            curr = curr.Left
        } else {
            curr = curr.Right
        }
    }
    return false
}

func (t *BSTree) Min() interface{} {
    if t.Root == nil {
        return nil
    }
    node := t.Root
    for node.Left != nil {
        node = node.Left
    }
    return node.Value
}

func (t *BSTree) Max() interface{} {
    if t.Root == nil {
        return nil
    }
    node := t.Root
    for node.Right != nil {
        node = node.Right
    }
    return node.Value
}

func (t *BSTree) findWithRecursion(root *BSTNode, value interface{}) *BSTNode {
    if root == nil {
        return nil
    }

    v := t.Comparator(root.Value, value)
    if v == 0 {
        return root
    } else if v > 0 {
        return t.findWithRecursion(root.Left, value)
    } else {
        return t.findWithRecursion(root.Right, value)
    }
}

func (t *BSTree) getLeftMax(node *BSTNode) (*BSTNode, *BSTNode) {
    parent, curr := node, node.Left
    for curr.Right != nil {
        parent = curr
        curr = curr.Right
    }
    return parent, curr
}

func (t *BSTree) Delete(value interface{}) *BSTree {
    // t.Root = t.deleteWithRecursion(t.Root, value)
    // return t
    var parent *BSTNode
    curr := t.Root
    for curr != nil {
        v := t.Comparator(curr.Value, value)
        if v > 0 {
            parent, curr = curr, curr.Left
        } else if v < 0 {
            parent, curr = curr, curr.Right
        } else {
            if curr.Left != nil && curr.Right != nil {
                nodeParent, node := t.getLeftMax(curr)
                curr.Value = node.Value
                if nodeParent == curr {
                    nodeParent.Left = node.Left
                } else {
                    nodeParent.Right = node.Left
                }
            } else {
                var node *BSTNode
                if curr.Left == nil {
                    node = curr.Right
                } else /* if curr.Right == nil */ {
                    node = curr.Left
                }
                if parent == nil {
                    t.Root = node
                } else if parent.Left == curr {
                    parent.Left = node
                } else /*if parent.Right == curr */ {
                    parent.Right = node
                }
            }
            t.size--
            return t
        }
    }
    return t
}

func (t *BSTree) deleteWithRecursion(root *BSTNode, value interface{}) *BSTNode {
    if root == nil {
        return nil
    }
    v := t.Comparator(root.Value, value)
    if v > 0 {
        root.Left = t.deleteWithRecursion(root.Left, value)
    } else if v < 0 {
        root.Right = t.deleteWithRecursion(root.Right, value)
    } else {
        t.size--
        if root.Left == nil {
            return root.Right
        } else if root.Right == nil {
            return root.Left
        } else {
            parent, node := t.getLeftMax(root)
            root.Value = node.Value
            // Root.Left = t.deleteWithRecursion(Root.Left, node.Value)
            if parent == root {
                parent.Left = node.Left
            } else {
                parent.Right = node.Left
            }
        }
    }
    return root
}

func (t *BSTree) Str(width int) string {
    var lines []string
    lines = append(lines, fmt.Sprintf("Binary Search Tree - %d Node:", t.size))
    t.formatInOrder(&lines, t.Root, 0, "H", width)
    return strings.Join(lines, "\n")
}

func (t *BSTree) String() string {
    return t.Str(9)
}

func (t *BSTree) formatInOrder(lines *[]string, node *BSTNode, height int, to string, length int) {
    if node != nil {
        t.formatInOrder(lines, node.Right, height+1, "v", length)
        val := fmt.Sprintf("%s%s%s", to, node.String(), to)
        lenM := len(val)
        lenL := (length - lenM) / 2
        lenR := length - lenM - lenL
        val = strings.Repeat(" ", lenL) + val + strings.Repeat(" ", lenR)
        *lines = append(*lines, strings.Repeat(" ", height*length)+val)
        t.formatInOrder(lines, node.Left, height+1, "^", length)
    }
}
