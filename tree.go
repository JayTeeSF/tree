// go run tree.go
package main

import "fmt"

// move this to a tree class

// Tree will be an interface that Branch & Leaf support
type tree interface {
	Tag() string
	SubTrees() []tree
}

type branch struct {
	tag      string
	subTrees []tree //should this be a method
}

type leaf struct {
	tag string
}

func (b branch) SubTrees() []tree {
	return b.subTrees
}

func (b branch) Tag() string {
	return b.tag
}

func (l leaf) SubTrees() []tree {
	return nil
}

func (l leaf) Tag() string {
	return l.tag
}

// need search method(s): dfs and bfs
// Can we do a common walk method (that accept blocks)
// or perhaps that uses go-routines and channels (instead of blocks)

func treeFind(t tree, search string) tree {
	var resultTree tree
	// resultTree := tree(nil)
	if search == t.Tag() {
		return t
	} else {
		for _, subTree := range t.SubTrees() {
			resultTree = treeFind(subTree, search)
			if nil != resultTree {
				return resultTree
			}
		}
	}
	return nil
}

func printTree(t tree, c int) {
	if c > 0 {
		for i := 0; i < c; i++ {
			fmt.Print("\t")
		}
		fmt.Print("-> ")
	}
	fmt.Println(t.Tag())
	count := 0
	for _, subTree := range t.SubTrees() {
		count += 1
		printTree(subTree, count)
	}
}

// call it from a tree printer (or something)
func main() {
	var found tree
	l1 := leaf{tag: "other string"}
	l2 := leaf{tag: "other string"}
	l3 := leaf{tag: "other string to match"}
	l4 := leaf{tag: "string to match"}
	l5 := leaf{tag: "other string"}

	b1 := branch{tag: "other str", subTrees: []tree{l1, l2, l3}}
	b2 := branch{tag: "other str", subTrees: []tree{l4, l5}}
	t1 := branch{tag: "main tree", subTrees: []tree{b1, b2}}
	s := "string to match"

	printTree(t1, 0)
	fmt.Println("Finding 'string to match'")
	found = treeFind(t1, s)
	if nil != found {
		fmt.Println("found it")
		printTree(found, 0)
	} else {
		fmt.Println("NOT found it")
	}
}