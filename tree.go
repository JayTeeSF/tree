// go run tree.go
package main

import (
	"encoding/json"
	"fmt"
)

// move this to a tree class

// Tree will be an interface that Branch & Leaf support
type tree interface {
	Name() string
	Children() []tree
}

type branch struct {
	Tag      string
	SubTrees []tree //should this be a method
}

type leaf struct {
	Tag string
}

func (b branch) Children() []tree {
	return b.SubTrees
}

func (b branch) Name() string {
	return b.Tag
}

func (l leaf) Children() []tree {
	return nil
}

func (l leaf) Name() string {
	return l.Tag
}

// need search method(s): dfs and bfs
// Can we do a common walk method (that accept blocks)
// or perhaps that uses go-routines and channels (instead of blocks)

func treeFind(t tree, search string) tree {
	var resultTree tree
	if search == t.Name() {
		return t
	} else {
		for _, subTree := range t.Children() {
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
	fmt.Println(t.Name())
	count := 0
	for _, subTree := range t.Children() {
		count += 1
		printTree(subTree, count)
	}
}

// call it from a tree printer (or something)
func main() {
	//var found tree
	l1 := &leaf{Tag: "other string"}
	l2 := &leaf{Tag: "other string"}
	l3 := &leaf{Tag: "almost string to match"}
	l4_a := &leaf{Tag: "child string to match"}
	l4_b := &leaf{Tag: "string to match but not quite"}
	l5 := &leaf{Tag: "other string"}

	b1 := &branch{Tag: "other str", SubTrees: []tree{l1, l2, l3}}
	b2_1_b_1 := &branch{Tag: "string to match", SubTrees: []tree{l4_a}}
	b2_1 := &branch{Tag: "string to match", SubTrees: []tree{l4_b, b2_1_b_1}}
	b2 := &branch{Tag: "other str", SubTrees: []tree{b2_1, l5}}
	t1 := &branch{Tag: "main tree", SubTrees: []tree{b1, b2}}
	s := "string to match"

	printTree(t1, 0)
	fmt.Println("Finding 'string to match'")
	found := treeFind(t1, s)
	if nil != found {
		fmt.Println("found it")
		//printTree(found, 0)
		out, err := json.Marshal(&found) // tree is not a struct, so it has no public attrs
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))

		/*
			fmt.Printf("%v", found) // format: {<val> [<vals>]}

			fmt.Printf("%#v", found) // format: main.<type>{<key>:<val> []main.<type>{<key>: [<vals>]}
		*/

		// best:
		//fmt.Printf("%+v", found) // format: {<key>: <val> <key>{[<vals>]}
	} else {
		fmt.Println("NOT found it")
	}
}
