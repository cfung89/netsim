package main

import ()

type Node struct {
	ID, Battery int
	Location    [3]int
	status      string
}

func NewNode() *Node {
	return &Node{}
}
