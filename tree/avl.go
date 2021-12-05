package tree

import (
    "fmt"
    "likebeta/gds/util"
    "strings"
)

type AVLNode struct {
    Left   *AVLNode
    Right  *AVLNode
    height int
    Value  interface{}
}

func (node *AVLNode) Height() int {
    if node == nil {
        return 0
    }
    return node.height
}

func (node *AVLNode) SyncHeight() int {
    if node == nil {
        return 0
    }
    node.height = max(node.Left.Height(), node.Right.Height()) + 1
    return node.height
}

func (node *AVLNode) String() string {
    return fmt.Sprintf("%v/%d", node.Value, node.height)
}

type AVLTree struct {
    Root       *AVLNode
    Comparator util.Comparator
    size       int
}

func NewAVLTree(comparator util.Comparator) *AVLTree {
    return &AVLTree{Comparator: comparator}
}

func (t *AVLTree) addWithRecursion(root *AVLNode, value interface{}) *AVLNode {
    if root == nil {
        t.size++
        return &AVLNode{Value: value, height: 1}
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
    node.SyncHeight()
    b.SyncHeight()
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
    node.SyncHeight()
    b.SyncHeight()
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
    node.SyncHeight()
    b.SyncHeight()
    c.SyncHeight()
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
    node.SyncHeight()
    b.SyncHeight()
    c.SyncHeight()
    return c
}

func (t *AVLTree) adjust(node *AVLNode) *AVLNode {
    lH, rH := node.Left.Height(), node.Right.Height()
    if lH-rH == 2 {
        if node.Left.Left.Height()-node.Left.Right.Height() > 0 {
            return llAdjust(node)
        } else {
            return lrAdjust(node)
        }
    } else if lH-rH == -2 {
        if node.Right.Left.Height()-node.Right.Right.Height() < 0 {
            return rrAdjust(node)
        } else {
            return rlAdjust(node)
        }
    } else {
        node.height = max(lH, rH) + 1
    }
    return node
}

func (t *AVLTree) Add(value interface{}) *AVLTree {
    t.Root = t.addWithRecursion(t.Root, value)
    return t
}

func (t *AVLTree) Find(value interface{}) bool {
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

func (t *AVLTree) Min() interface{} {
    if t.Root == nil {
        return nil
    }
    node := t.Root
    for node.Left != nil {
        node = node.Left
    }
    return node.Value
}

func (t *AVLTree) Max() interface{} {
    if t.Root == nil {
        return nil
    }
    node := t.Root
    for node.Right != nil {
        node = node.Right
    }
    return node.Value
}

func (t *AVLTree) Delete(value interface{}) *AVLTree {
    t.Root = t.deleteWithRecursion(t.Root, value)
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
            if root.Left.Height() > root.Right.Height() {
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

func (t *AVLTree) Str(width int) string {
    var lines []string
    lines = append(lines, fmt.Sprintf("AVL Tree - %d Node:", t.size))
    t.formatInOrder(&lines, t.Root, 0, "H", width)
    return strings.Join(lines, "\n")
}

func (t *AVLTree) String() string {
    return t.Str(9)
}

func (t *AVLTree) formatInOrder(lines *[]string, node *AVLNode, height int, to string, length int) {
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
