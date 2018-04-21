package main
import (
	"fmt"
)

type Node struct {
	Successor      *Node
	KeyValue       map[int]int
	FingerTable    map[int]int
	Predecessor    *Node
	NodeIdentifier int
}

var chrodRingNodes = []*Node{}


func main() {
    fmt.Println("hello world")
	ringPos := 1
	if(len(chrodRingNodes) == 0) {
		node := Node{}
		node.NodeIdentifier = ringPos
		node.Successor = &node
		node.Predecessor = &node
		node.KeyValue = make(map[int]int)
		chrodRingNodes = append(chrodRingNodes, &node)
	}
	
	fmt.Println(chrodRingNodes[0].Successor.NodeIdentifier)
	fmt.Println(len(chrodRingNodes))
	
	ringPos = 3
	sponsoringNodeId := 1
	newNode := Node{}
	startNode := Node{}
	newNode.NodeIdentifier = ringPos
	sponsoringNode := Node{}
	for _, element := range chrodRingNodes {
			if(sponsoringNodeId == element.NodeIdentifier) {
				sponsoringNode = *element
		}
	}
	
	newNode.Successor, newNode.Predecessor, startNode = findSuccessorAndPredecessor(ringPos, sponsoringNode)
	if(ringPos < sponsoringNodeId) {
		startNode.Successor.Predecessor = &newNode
		startNode.Successor = &newNode
	}
	if(ringPos > sponsoringNodeId) {
		startNode.Predecessor.Successor = &newNode
		startNode.Predecessor = &newNode
	}
	
	chrodRingNodes = append(chrodRingNodes, &newNode)
	
	fmt.Println(chrodRingNodes[0].Successor.NodeIdentifier, chrodRingNodes[1].Successor.NodeIdentifier)
	
}


func findSuccessorAndPredecessor(ringPos int, sponsoringNode Node) (*Node, *Node, Node) {

	succNode := Node{}
	predNode := Node{}
	startNode := sponsoringNode
	
	if(ringPos < sponsoringNode.NodeIdentifier) {
		startNode := sponsoringNode.Predecessor
		for {
			if( ringPos < startNode.NodeIdentifier) {
				succNode = *startNode.Successor
				predNode = *startNode
				break
			} else {
				startNode = startNode.Predecessor
			}
		}
	}
	if(ringPos > sponsoringNode.NodeIdentifier) {
		startNode := sponsoringNode.Successor
		for {
			if( ringPos > startNode.NodeIdentifier) {
				succNode = *startNode
				predNode = *startNode.Predecessor
				break
			} else {
				startNode = startNode.Successor
			}
		}
	}
	return &succNode, &predNode, startNode
}