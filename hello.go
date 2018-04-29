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

type Message struct {
	Do             string   `json:"do"`
	Sponsoringnode int      `json:"sponsoring-node"`
	Mode           int      `json:"mode"`
	Respondto      int      `json:"respond-to"`
	TargetId       int      `json:"target-id"`
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
	nodeFunc(ringPos, sponsoringNodeId)
	
	ringPos1 := 6
	sponsoringNodeId1 := 1
	nodeFunc(ringPos1, sponsoringNodeId1)
	
	fmt.Println(nodeChannelmap[1].NodeIdentifier, nodeChannelmap[1].Successor.NodeIdentifier, nodeChannelmap[1].Predecessor.NodeIdentifier)
	
	ringPos2 := 5
	sponsoringNodeId2 := 1
	nodeFunc(ringPos2, sponsoringNodeId2)
	
	fmt.Println(nodeChannelmap[3].NodeIdentifier, nodeChannelmap[3].Successor.NodeIdentifier, nodeChannelmap[3].Predecessor.NodeIdentifier)
	fmt.Println(nodeChannelmap[1].NodeIdentifier, nodeChannelmap[1].Successor.NodeIdentifier, nodeChannelmap[1].Predecessor.NodeIdentifier)
	
	ringPos3 := 4
	sponsoringNodeId3 := 1
	nodeFunc(ringPos3, sponsoringNodeId3)
	
	fmt.Println(nodeChannelmap[1].NodeIdentifier, nodeChannelmap[1].Successor.NodeIdentifier, nodeChannelmap[1].Predecessor.NodeIdentifier)
	
	ringPos4 := 2
	sponsoringNodeId4 := 5
	nodeFunc(ringPos4, sponsoringNodeId4)
	
	fmt.Println(nodeChannelmap[3].NodeIdentifier, nodeChannelmap[3].Successor.NodeIdentifier, nodeChannelmap[3].Predecessor.NodeIdentifier)
	fmt.Println(nodeChannelmap[1].NodeIdentifier, nodeChannelmap[1].Successor.NodeIdentifier, nodeChannelmap[1].Predecessor.NodeIdentifier)
	
}

func nodeFunc(ringPos int, sponsoringNodeId int) {
	newNode := Node{}
	startNode := Node{}
	successor := Node{}
	predecessor := Node{}
	newNode.NodeIdentifier = ringPos
	sponsoringNode := nodeChannelmap[sponsoringNodeId]
	
	successor, predecessor, startNode = findSuccessorAndPredecessor(ringPos, sponsoringNode)
	newNode.Successor = &successor
	newNode.Predecessor = &predecessor
	
	updateMap(newNode, startNode, sponsoringNode, ringPos)
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