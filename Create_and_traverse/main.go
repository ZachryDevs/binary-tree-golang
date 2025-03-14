package main

import (
	"errors"
	"fmt"
)

func main() {
	// input := []int{5, 3, 7, 2, 4, 6, 8}
	input := []int{1}
	tree := MakeTree()
	for _, num := range input {
		tree.Insert(num)
	}
	vals := tree.Traverse()
	fmt.Println(vals)
	test := tree.Search(10)
	fmt.Printf("Searching Tree for %d returned %t\n", 10, test)
	for _, num := range input {
		fmt.Printf("Searching Tree for %d returned %t\n", num, tree.Search(num))
	}
	maxVal, err := tree.Max()
	if err != nil {
		fmt.Println(err)
	}
	minVal, err := tree.Min()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The max value in this tree is: %d\nThe min value in this tree is: %d\n", *maxVal, *minVal)
}

type Tree struct {
	val   *int
	left  *Tree
	right *Tree
}

func MakeTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(num int) {
	// If the tree is empty (nil pointer), set the root value
	if t.val == nil {
		t.val = &num
		return
	}

	// Insert logic based on comparison of num with the current node's value
	if num < *t.val {
		if t.left == nil {
			t.left = &Tree{}
		}
		t.left.Insert(num)
	} else {
		if t.right == nil {
			t.right = &Tree{}
		}
		t.right.Insert(num)
	}
}

func (t *Tree) Traverse() []int {
	var intList []int
	// Traverse left subtree
	if t.left != nil {
		intList = append(intList, t.left.Traverse()...)
	}

	// Append current node's value (dereferencing t.val to get the actual value)
	if t.val != nil {
		intList = append(intList, *t.val)
	}

	// Traverse right subtree
	if t.right != nil {
		intList = append(intList, t.right.Traverse()...)
	}

	return intList
}

func (t *Tree) Search(num int) bool {
	if t.val == nil {
		return false
	}

	if *t.val == num {
		return true
	} else if num < *t.val {
		if t.left == nil {
			return false
		} else {
			return t.left.Search(num)
		}
	} else {
		if t.right == nil {
			return false
		} else {
			return t.right.Search(num)
		}
	}
}

func (t *Tree) Min() (*int, error) {
	var minNode *int
	var err error
	if t.left != nil {
		minNode, err = t.left.Min()
		if err != nil {
			return nil, err
		}
	} else {
		if t.val == nil {
			return nil, errors.New("internal error")
		}
		return t.val, nil
	}
	return minNode, nil
}

func (t *Tree) Max() (*int, error) {
	var maxNode *int
	var err error
	if t.right != nil {
		maxNode, err = t.right.Max()
		if err != nil {
			return nil, err
		}
	} else {
		if t.val == nil {
			return nil, errors.New("internal error")
		}
		return t.val, nil
	}
	return maxNode, nil
}
