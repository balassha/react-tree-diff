package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

type option int

const (
	INSERT option = 1 << iota
	MOVE option = 1 << iota
	REMOVE   option = 1 << iota
)

// Binary unsorted tree
type ReactTree struct {
	ListOfNodes []string
	NodeSet map[string]bool
}


/*
*	Insert Node to React Tree
*/
func (r *ReactTree) InsertNode(val string, index int) bool {
	if index > len(r.ListOfNodes) || index <= 0 {
		//fmt.Println("length too big or too small")
		return false
	}

	if val == "" {
		return false //cannot insert nil value
	}

	if index > 1 && r.ListOfNodes[index/2] == "" {
		//fmt.Println("Parent is not exist:", nodeIndex, nodeIndex/2)
		return false
	}

	if _, exist := r.NodeSet[val]; exist {
		//fmt.Println("Element duplicated:", val)
		return false
	}

	r.ListOfNodes[index] = val
	r.NodeSet[val] = true
	return true
}

func (r *ReactTree) deleteNode(index int) {
	if r.ListOfNodes[index] == "" {
		return
	}

	nextIndex := index * 2
	if nextIndex < len(r.ListOfNodes) && r.ListOfNodes[nextIndex] != "" {
		r.deleteNode(nextIndex)
	}

	if nextIndex < len(r.ListOfNodes) && r.ListOfNodes[nextIndex+1] != "" {
		r.deleteNode(nextIndex + 1)
	}

	r.deleteUnitNode(index)
}

func (r *ReactTree) deleteUnitNode(index int) {
	if index > len(r.ListOfNodes) {
		return
	}

	if r.ListOfNodes[index] == "" {
		return
	}

	val := r.ListOfNodes[index]
	r.ListOfNodes[index] = ""
	delete(r.NodeSet, val)
}

//Clone current tree to another new one
func (r *ReactTree) Clone() *ReactTree {
	nT := NewReactTree(len(r.ListOfNodes))
	for k, v := range r.ListOfNodes {
		nT.ListOfNodes[k] = v
	}

	for k, v := range r.NodeSet {
		nT.NodeSet[k] = v
	}

	return nT
}

//Remove node via node value, return true if node exist and successful delete
func (r *ReactTree) RemoveNode(val string) bool {
	if len(r.NodeSet) == 0 {
		//fmt.Println("Empty tree deletion")
		return false
	}

	if _, exist := r.NodeSet[val]; !exist {
		//fmt.Println("value not exist for deletion")
		return false
	}

	for index, v := range r.ListOfNodes {
		if v == val {
			r.deleteNode(index)
		}
	}

	return true
}

//Return node index via node value, return -1 if node is not exist
func (r *ReactTree) GetIndexOfNode(searchTarget interface{}) int {
	for index, value := range r.ListOfNodes {
		if value == searchTarget {
			return index
		}
	}
	return -1
}

// Diff Tree will diff with input target tree, if not identical will replace to new one
func (r *ReactTree) DiffTree(targetTree *ReactTree, opt option) bool {
	refTree := r.Clone()

	//fmt.Println("option=", option, " it is match with ", option&REMOVE_NODE)
	for newIndex, value := range targetTree.ListOfNodes {
		if value == "" {
			continue
		}

		oldIndex := refTree.GetIndexOfNode(value)

		//INSERT_MARKUP
		if (opt & INSERT) == INSERT && oldIndex == -1 {
			r.InsertNode(value, newIndex)
			continue
		}

		//MOVE_EXISTING
		if opt & MOVE == MOVE {
			if oldIndex != -1 && oldIndex != newIndex {
				if r.ListOfNodes[oldIndex] == value {
					r.ListOfNodes[oldIndex] = ""
				}
				r.ListOfNodes[newIndex] = value
			}

			//fmt.Println("Enter move:", value, oldIndex, newIndex, r.ListOfNodes, refTree.ListOfNodes)
		}
	}

	//REMOVE_NODE
	if opt & REMOVE == REMOVE {
		//fmt.Println("Enter remove node")
		for k, _ := range r.NodeSet {
			if _, exist := targetTree.NodeSet[k]; !exist {
				r.RemoveNode(k)
			}
		}
	}

	return false
}

func system(s string) {
	cmd := exec.Command(`/bin/sh`, `-c`, s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", out.String())
}

//New a React Diff Tree with define size
//The binary tree with basic alignment with array
//0-> 1, 2
//1-> 3, 4
//2-> 5, 6 ....
func NewReactTree(treeSize int) *ReactTree {
	nodes := make([]string, treeSize)
	newRD := new(ReactTree)
	newRD.ListOfNodes = nodes
	newRD.NodeSet = make(map[string]bool)
	return newRD
}
