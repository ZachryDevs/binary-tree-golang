package main

import "fmt"

func main() {
	input := []int{5, 3, 7, 2, 4, 6, 8}
	tree := MakeTree()
	for _, num := range input {
		tree.Insert(num)
	}
	test := tree.Search(10)
	fmt.Printf("Searching Tree for %d returned %t\n", 10, test)
	for _, num := range input {
		fmt.Printf("Searching Tree for %d returned %t\n", num, tree.Search(num))
	}
}

func MakeTree() *Tree {
	return &Tree{}
}

type Tree struct {
	val   *int
	left  *Tree
	right *Tree
}

func (t *Tree) Insert(num int) {
	if t.val == nil {
		t.val = &num
		return
	}

	if num < *t.val {
		if t.left == nil {
			t.left = &Tree{val: &num}
		} else {
			t.left.Insert(num)
		}
	} else {
		if t.right == nil {
			t.right = &Tree{val: &num}
		} else {
			t.right.Insert(num)
		}
	}

}

func (t *Tree) Search(num int) bool {
	if t.val == &num {
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
