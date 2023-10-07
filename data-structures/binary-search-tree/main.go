package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value int
	left  *node
	right *node
}

func (n node) String() string {
	return strconv.Itoa(n.value)
}

type BinarySearchTree struct {
	root   *node
	length int
}

func (bst BinarySearchTree) Len() int {
	return bst.length
}

func (bst *BinarySearchTree) Add(value int) {
	root, err := bst.add(bst.root, value)
	if err == nil {
		bst.root = root
		bst.length++
	}
}

func (bst *BinarySearchTree) add(n *node, value int) (*node, error) {
	if n == nil {
		return &node{value: value}, nil
	}

	if value > n.value {
		node, err := bst.add(n.right, value)
		if err != nil {
			// bypass the error so length does not increase
			return n, err
		}
		n.right = node
		return n, nil
	}

	if value < n.value {
		node, err := bst.add(n.left, value)
		if err != nil {
			// bypass the error so length does not increase
			return n, err
		}
		n.left = node
		return n, nil
	}

	return n, errors.New("value already in bst")
}

func (bst *BinarySearchTree) Remove(value int) {
	if bst.Contains(value) {
		bst.length--
	}
	bst.root = bst.remove(bst.root, value)
}

func (bst BinarySearchTree) Height() int {
	return bst.height(bst.root)
}

func (bst BinarySearchTree) height(n *node) int {
	if n == nil {
		return 0
	}

	heightLeft := bst.height(n.left)
	heightRight := bst.height(n.right)
	return 1 + max(heightLeft, heightRight)
}

func (bst BinarySearchTree) IsBalanced() bool {
	balanced, _ := bst.isBalanced(bst.root)
	return balanced
}

func (bst BinarySearchTree) isBalanced(n *node) (bool, int) {
	if n == nil {
		return true, 0
	}

	leftBalanced, leftHeight := bst.isBalanced(n.left)
	rightBalanced, rightHeight := bst.isBalanced(n.right)

	heightDiff := abs(leftHeight - rightHeight)

	balanced := leftBalanced && rightBalanced && heightDiff <= 1

	return balanced, 1 + max(leftHeight, rightHeight)
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func (bst *BinarySearchTree) remove(n *node, value int) *node {
	// find the highest value in the left subtree
	// copy its value to the current root node
	// delete the node found
	// if there is no left subtree, return the right subtree

	if n == nil {
		return n
	}

	if n.value < value {
		n.right = bst.remove(n.right, value)
	} else if n.value > value {
		n.left = bst.remove(n.left, value)
	} else {
		if n.left == nil {
			return n.right
		} else {
			// copy the highest value in the right children to the current node
			tmp := n.left
			for tmp.right != nil {
				tmp = tmp.right
			}
			n.value = tmp.value

			// try to delete the value from the left
			n.left = bst.remove(n.left, tmp.value)
		}
	}

	return n
}

func (bst BinarySearchTree) Contains(value int) bool {
	_, found := bst.search(bst.root, value)
	return found
}

func (bst BinarySearchTree) search(n *node, value int) (*node, bool) {
	if n == nil {
		return nil, false
	}

	if value == n.value {
		return n, true
	}

	if value > n.value {
		return bst.search(n.right, value)
	} else {
		return bst.search(n.left, value)
	}
}

func (bst *BinarySearchTree) Exists(value int) bool {
	return bst.Contains(value)
}

func (bst BinarySearchTree) String() string {
	sb := strings.Builder{}
	bst.values(&sb)
	return sb.String()
}

func (bst BinarySearchTree) values(sb *strings.Builder) {
	bst.valuesByNode(sb, bst.root)
}

func (bst BinarySearchTree) valuesByNode(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}

	bst.valuesByNode(sb, root.left)
	sb.WriteString(" " + root.String() + " ")
	bst.valuesByNode(sb, root.right)
}

func (bst BinarySearchTree) ToSlice() []int {
	return bst.toSlice(bst.root)
}

func (bst BinarySearchTree) toSlice(n *node) []int {
	if n == nil {
		return nil
	}

	data := []int{}

	leftData := bst.toSlice(n.left)
	rightData := bst.toSlice(n.right)

	data = append(data, n.value)
	data = append(data, leftData...)
	data = append(data, rightData...)

	return data
}

func main() {
	bst := BinarySearchTree{}
	bst.Add(3)
	bst.Add(1)
	bst.Add(2)
	bst.Add(4)
	bst.Add(5)
	bst.Add(6)
	bst.Add(7)
	bst.Add(8)

	fmt.Println("items:", bst)
	fmt.Println("height:", bst.Height())
	fmt.Println("balanced:", bst.IsBalanced())

	fmt.Println("sliced:", bst.ToSlice())
	fmt.Println("size:", bst.Len())
}
