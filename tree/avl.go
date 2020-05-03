package tree

import (
    "fmt"
    "github.com/likebeta/gds/util"
    "strings"
)

type AVLNode struct {
    Left   *AVLNode
    Right  *AVLNode
    Height int
    Value  interface{}
}

type AVLTree struct {
    root       *AVLNode
    Comparator util.Comparator
    size       int
}

func NewAVLTree(comparator util.Comparator) *AVLTree {
    return &AVLTree{Comparator: comparator}
}

func (t *AVLTree) addWithRecursion(root *AVLNode, value interface{}) *AVLNode {
    if root == nil {
        t.size++
        return &AVLNode{Value: value, Height: 1}
    }
    v := t.Comparator(root.Value, value)
    if v > 0 {
        root.Left = t.addWithRecursion(root.Left, value)
        root = t.adjust(root)
    } else if v < 0 {
        root.Right = t.addWithRecursion(root.Right, value)
        root = t.adjust(root)
    }
    return root
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func getHeight(node *AVLNode) int {
    if node == nil {
        return 0
    }
    return node.Height
}

func updateHeight(node *AVLNode) {
    node.Height = max(getHeight(node.Left), getHeight(node.Right)) + 1
}

func getLeftMax(node *AVLNode) *AVLNode {
    curr := node.Left
    for curr.Right != nil {
        curr = curr.Right
    }
    return curr
}

func getRightMin(node *AVLNode) *AVLNode {
    curr := node.Right
    for curr.Left != nil {
        curr = curr.Left
    }
    return curr
}

/*
  		  A                  B
         /                  / \
        B                  C   A
       / \                    /
      C   D                  D
*/
func llAdjust(node *AVLNode) *AVLNode {
    b := node.Left
    node.Left = b.Right
    b.Right = node
    updateHeight(node)
    updateHeight(b)
    return b
}

/*
  	  A                      B
       \                    / \
        B                  A   D
       / \                  \
      C   D                  C
*/
func rrAdjust(node *AVLNode) *AVLNode {
    b := node.Right
    node.Right = b.Left
    b.Left = node
    updateHeight(node)
    updateHeight(b)
    return b
}

/*
		A
       /                   C
      B                  /   \
       \                B     A
        C                \   /
       / \                D E
      D   E
*/
func lrAdjust(node *AVLNode) *AVLNode {
    b := node.Left
    c := b.Right
    node.Left = c.Right
    b.Right = c.Left
    c.Left = b
    c.Right = node
    updateHeight(node)
    updateHeight(b)
    updateHeight(c)
    return c
}

/*
     A
      \                   C
       B                /   \
      /                A     B
     C                  \   /
    / \                  D E
   D   E
*/
func rlAdjust(node *AVLNode) *AVLNode {
    b := node.Right
    c := b.Left
    node.Right = c.Left
    b.Left = c.Right
    c.Left = node
    c.Right = b
    updateHeight(node)
    updateHeight(b)
    updateHeight(c)
    return c
}

func (t *AVLTree) adjust(node *AVLNode) *AVLNode {
    lH, rH := getHeight(node.Left), getHeight(node.Right)
    if lH-rH == 2 {
        if getHeight(node.Left.Left)-getHeight(node.Left.Right) > 0 {
            return llAdjust(node)
        } else {
            return lrAdjust(node)
        }
    } else if lH-rH == -2 {
        if getHeight(node.Right.Left)-getHeight(node.Right.Right) < 0 {
            return rrAdjust(node)
        } else {
            return rlAdjust(node)
        }
    } else {
        node.Height = max(lH, rH) + 1
    }
    return node
}

func (t *AVLTree) Add(value interface{}) *AVLTree {
    t.root = t.addWithRecursion(t.root, value)
    return t
}

func (t *AVLTree) Find(value interface{}) bool {
    curr := t.root
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

func (t *AVLTree) Min() interface{} {
    if t.root == nil {
        return nil
    }
    node := t.root
    for node.Left != nil {
        node = node.Left
    }
    return node.Value
}

func (t *AVLTree) Max() interface{} {
    if t.root == nil {
        return nil
    }
    node := t.root
    for node.Right != nil {
        node = node.Right
    }
    return node.Value
}

func (t *AVLTree) Delete(value interface{}) *AVLTree {
    t.root = t.deleteWithRecursion(t.root, value)
    return t
}

func (t *AVLTree) deleteWithRecursion(root *AVLNode, value interface{}) *AVLNode {
    if root == nil {
        return nil
    }
    v := t.Comparator(root.Value, value)
    if v > 0 {
        root.Left = t.deleteWithRecursion(root.Left, value)
    } else if v < 0 {
        root.Right = t.deleteWithRecursion(root.Right, value)
    } else {
        if root.Left == nil {
            root = root.Right
            t.size--
        } else if root.Right == nil {
            root = root.Left
            t.size--
        } else {
            if getHeight(root.Left) > getHeight(root.Right) {
                node := getLeftMax(root)
                root.Value = node.Value
                root.Left = t.deleteWithRecursion(root.Left, node.Value)
            } else {
                node := getRightMin(root)
                root.Value = node.Value
                root.Right = t.deleteWithRecursion(root.Right, node.Value)
            }
        }
    }
    if root != nil {
        root = t.adjust(root)
    }
    return root
}

func (t *AVLTree) String() string {
    var lines []string
    lines = append(lines, fmt.Sprintf("AVL Tree - %d Node:", t.size))
    t.formatInOrder(&lines, t.root, 0, "H", 9)
    return strings.Join(lines, "\n")
}

func (t *AVLTree) formatInOrder(lines *[]string, node *AVLNode, height int, to string, length int) {
    if node != nil {
        t.formatInOrder(lines, node.Right, height+1, "v", length)
        val := fmt.Sprintf("%s%d%s", to, node.Value, to)
        lenM := len(val)
        lenL := (length - lenM) / 2
        lenR := length - lenM - lenL
        val = strings.Repeat(" ", lenL) + val + strings.Repeat(" ", lenR)
        *lines = append(*lines, strings.Repeat(" ", height*length)+val)
        t.formatInOrder(lines, node.Left, height+1, "^", length)
    }
}
