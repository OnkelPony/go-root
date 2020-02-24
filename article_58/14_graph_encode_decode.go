package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Item int

type Node struct {
	Value Item
	Left  *Node
	Right *Node
}

func main() {
	var a, b, c Node

	a.Value = 1
	b.Value = 2
	c.Value = 3
	a.Left = &b
	a.Right = &c
	b.Left = nil
	b.Right = &c
	c.Left = nil
	c.Right = nil

	var x Node
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&a)
	decoder := gob.NewDecoder(&buffer)
	err = decoder.Decode(&x)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Node:        %p %v\n", &x, x)
	fmt.Printf("Left:        %p %v\n", x.Left, *x.Left)
	fmt.Printf("Right:       %p %v\n", x.Right, *x.Right)
	fmt.Printf("Left/Right:  %p %v\n", x.Left.Right, *x.Left.Right)
}
