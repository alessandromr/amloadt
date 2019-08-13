package sharedq

import "log"

//SharedList struct
type SharedList struct {
	index map[string]*Node
	first *Node
	last  *Node
}

//Node struct
type Node struct {
	value    string
	prevNode *Node
	nextNode *Node
}

//Push function
func (l *SharedList) Push(value string) {
	newNode := Node{}
	if l.last != nil {
		newNode = Node{
			value:    value,
			prevNode: l.last,
			nextNode: nil,
		}
		l.last.nextNode = &newNode
		l.last = &newNode
	} else {
		newNode = Node{
			value:    value,
			prevNode: nil,
			nextNode: nil,
		}
		if l.first == nil {
			l.first = &newNode
		}
		l.last = &newNode
	}
	if l.index == nil {
		l.index = make(map[string]*Node)
	}
	l.index[value] = &newNode
}

//Pull function
func (l *SharedList) Pull() Node {
	if l.first != nil {
		return *l.first
	}
	log.Fatalln("Queue is empty")
	return *l.first
}

//Pop function
func (l *SharedList) Pop() Node {
	if l.first == nil && l.last == nil {
		log.Fatalln("Queue is empty")
	}
	var returns Node
	if l.first != nil {
		returns = *l.first
		if l.first.nextNode != nil {
			newfirstNode := l.first.nextNode
			newfirstNode.prevNode = nil
			l.first = newfirstNode
		} else {
			l.first = nil
			l.last = nil
		}
	}
	return returns
}

//DirectPop function
func (l *SharedList) DirectPop(key string) Node {
	if l.first == nil && l.last == nil {
		log.Fatalln("Queue is empty")
	}

	selectedNode := l.index[key]

	if selectedNode.nextNode == nil && selectedNode.prevNode == nil {
		l.first = nil
		l.last = nil
		return *selectedNode
	}

	if selectedNode.prevNode == nil {
		selectedNode.nextNode.prevNode = nil
	}
	if selectedNode.nextNode == nil {
		selectedNode.prevNode.nextNode = nil
	}

	if selectedNode.nextNode != nil && selectedNode.prevNode != nil {
		prevNode := selectedNode.prevNode
		nextNode := selectedNode.nextNode
		prevNode.nextNode = selectedNode.nextNode
		nextNode.prevNode = selectedNode.prevNode
	}

	if selectedNode == l.first {
		if selectedNode.nextNode == nil {
			l.first = nil
		} else {
			l.first = selectedNode.nextNode
		}
	}

	if selectedNode == l.last {
		if selectedNode.prevNode == nil {
			l.last = nil
		} else {
			l.last = selectedNode.prevNode
		}
	}

	return *selectedNode
}
