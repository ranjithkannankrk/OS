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

var nodeChannelmap = make(map[int] Node)


func main() {

	ringPos := 1
	if(len(nodeChannelmap) == 0) {
		node := Node{}
		node.NodeIdentifier = ringPos
		node.Successor = &node
		node.Predecessor = &node
		node.KeyValue = make(map[int]int)
		nodeChannelmap[ringPos] = node
	}
	
	
	ringPos = 3
	sponsoringNodeId := 1
	newNode := Node{}
	startNodes := Node{}
	successor := Node{}
	predecessor := Node{}
	newNode.NodeIdentifier = ringPos
	sponsoringNode := nodeChannelmap[sponsoringNodeId]
	
	successor, predecessor, startNodes = findSuccessorAndPredecessor(ringPos, sponsoringNode)
	newNode.Successor = &successor
	newNode.Predecessor = &predecessor
	
	if(ringPos > sponsoringNode.NodeIdentifier) {
		if(len(nodeChannelmap) == 1) {
			startNodes.Successor = &newNode
			startNodes.Predecessor = &newNode
		}
	}
	nodeChannelmap[ringPos] = newNode
	nodeChannelmap[startNodes.NodeIdentifier] = startNodes
	
	ringPos1 := 6
	sponsoringNodeId1 := 1
	newNode1 := Node{}
	startNode1 := Node{}
	successor1 := Node{}
	predecessor1 := Node{}
	newNode1.NodeIdentifier = ringPos1
	sponsoringNode1 := nodeChannelmap[sponsoringNodeId1]
	
	successor1, predecessor1, startNode1 = findSuccessorAndPredecessor(ringPos1, sponsoringNode1)
	newNode1.Successor = &successor1
	newNode1.Predecessor = &predecessor1
	
	updateMap(newNode1, startNode1, sponsoringNode1, ringPos1)	
	
	fmt.Println(nodeChannelmap[1].NodeIdentifier, nodeChannelmap[1].Successor.NodeIdentifier, nodeChannelmap[1].Predecessor.NodeIdentifier)
	
	
	ringPos2 := 5
	sponsoringNodeId2 := 1
	newNode2 := Node{}
	startNode2 := Node{}
	successor2 := Node{}
	predecessor2 := Node{}
	newNode2.NodeIdentifier = ringPos2
	sponsoringNode2 := nodeChannelmap[sponsoringNodeId2]
	
	successor2, predecessor2, startNode2 = findSuccessorAndPredecessor(ringPos2, sponsoringNode2)
	newNode2.Successor = &successor2
	newNode2.Predecessor = &predecessor2
	
	fmt.Println(newNode2.NodeIdentifier, newNode2.Successor.NodeIdentifier, newNode2.Predecessor.NodeIdentifier,startNode2.NodeIdentifier)
	
	updateMap(newNode2, startNode2, sponsoringNode2, ringPos2)
	
	fmt.Println(nodeChannelmap[3].NodeIdentifier, nodeChannelmap[3].Successor.NodeIdentifier, nodeChannelmap[3].Predecessor.NodeIdentifier)
	fmt.Println(nodeChannelmap[1].NodeIdentifier, nodeChannelmap[1].Successor.NodeIdentifier, nodeChannelmap[1].Predecessor.NodeIdentifier)
	
	ringPos3 := 4
	sponsoringNodeId3 := 1
	newNode3 := Node{}
	startNode3 := Node{}
	successor3 := Node{}
	predecessor3 := Node{}
	newNode3.NodeIdentifier = ringPos3
	sponsoringNode3 := nodeChannelmap[sponsoringNodeId3]
	
	successor3, predecessor3, startNode3 = findSuccessorAndPredecessor(ringPos3, sponsoringNode3)
	newNode3.Successor = &successor3
	newNode3.Predecessor = &predecessor3

	fmt.Println(newNode3.NodeIdentifier, newNode3.Successor.NodeIdentifier, newNode3.Predecessor.NodeIdentifier,startNode3.NodeIdentifier)
	
	updateMap(newNode3, startNode3, sponsoringNode3, ringPos3)
	
	fmt.Println(nodeChannelmap[1].NodeIdentifier, nodeChannelmap[1].Successor.NodeIdentifier, nodeChannelmap[1].Predecessor.NodeIdentifier)
	
	ringPos4 := 2
	sponsoringNodeId4 := 5
	newNode4 := Node{}
	startNode4 := Node{}
	successor4 := Node{}
	predecessor4 := Node{}
	newNode4.NodeIdentifier = ringPos4
	sponsoringNode4 := nodeChannelmap[sponsoringNodeId4]
	
	successor4, predecessor4, startNode4 = findSuccessorAndPredecessor(ringPos4, sponsoringNode4)
	newNode4.Successor = &successor4
	newNode4.Predecessor = &predecessor4
	
	fmt.Println(newNode4.NodeIdentifier, newNode4.Successor.NodeIdentifier, newNode4.Predecessor.NodeIdentifier,startNode4.NodeIdentifier)
	
	updateMap(newNode4, startNode4, sponsoringNode4, ringPos4)
	
	fmt.Println(nodeChannelmap[3].NodeIdentifier, nodeChannelmap[3].Successor.NodeIdentifier, nodeChannelmap[3].Predecessor.NodeIdentifier)
	fmt.Println(nodeChannelmap[1].NodeIdentifier, nodeChannelmap[1].Successor.NodeIdentifier, nodeChannelmap[1].Predecessor.NodeIdentifier)
	
}

func updateMap(newNode Node, startNode Node, sponsoringNode Node, ringPos int) {

	if(ringPos > sponsoringNode.NodeIdentifier) {
		if(len(nodeChannelmap) == 1) {
			startNode.Successor = &newNode
			startNode.Predecessor = &newNode
		}
		if(startNode.Successor.NodeIdentifier < ringPos) {
			if(ringPos < startNode.NodeIdentifier) {
				updateNode := nodeChannelmap[startNode.Predecessor.NodeIdentifier]
				updateNode.Successor = &newNode
				nodeChannelmap[updateNode.NodeIdentifier] = updateNode
				startNode.Predecessor = &newNode
				
			} else {
				updateNode := nodeChannelmap[startNode.Successor.NodeIdentifier]
				updateNode.Predecessor = &newNode
				nodeChannelmap[updateNode.NodeIdentifier] = updateNode
				startNode.Successor = &newNode
			}
		}
		if(ringPos < startNode.NodeIdentifier) {
			updateNode := nodeChannelmap[startNode.Predecessor.NodeIdentifier]
			updateNode.Successor = &newNode
			nodeChannelmap[updateNode.NodeIdentifier] = updateNode
			startNode.Predecessor = &newNode
		}
	}
	if(ringPos < sponsoringNode.NodeIdentifier) {
		updateNode := nodeChannelmap[startNode.Successor.NodeIdentifier]
		updateNode.Predecessor = &newNode
		nodeChannelmap[updateNode.NodeIdentifier] = updateNode
		startNode.Successor = &newNode
	}
	
	nodeChannelmap[ringPos] = newNode
	nodeChannelmap[startNode.NodeIdentifier] = startNode
}


func findSuccessorAndPredecessor(ringPos int, sponsoringNode Node) (Node, Node, Node) {
	
	startNode := sponsoringNode
	
	if(ringPos > sponsoringNode.NodeIdentifier) {
	i := 1
	for{
		if(len(nodeChannelmap) == 1) {
			succNode := sponsoringNode
			predNode := sponsoringNode
			return succNode, predNode, sponsoringNode
			break
			}
		if(i == len(nodeChannelmap)) {
			if(ringPos < startNode.NodeIdentifier){
				succNode := startNode
				predNode := *startNode.Predecessor
				return succNode, predNode, startNode
			} else {
				succNode := *startNode.Successor
				predNode := startNode
				return succNode, predNode, startNode
			}
		}
		if(ringPos < startNode.NodeIdentifier){
				succNode := startNode
				predNode := *startNode.Predecessor
				return succNode, predNode, startNode
		}
		startNode = nodeChannelmap[startNode.Successor.NodeIdentifier]
		i = i + 1
		}
	}
	if(ringPos < sponsoringNode.NodeIdentifier) {
		
		for {
			if(ringPos > startNode.NodeIdentifier) {
				succNode := *startNode.Successor
				predNode := startNode
				return succNode, predNode, startNode
			}
			startNode = nodeChannelmap[startNode.Predecessor.NodeIdentifier]
		}
		
	}
	
	return sponsoringNode, sponsoringNode, sponsoringNode
}