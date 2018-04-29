package main
import (
	"fmt"
	"math/rand"
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
	
	action := "join-ring"
	for i := 0; i < 5; i++ {
		if(action == "join-ring") {
			message := Message{}
			message.Do = action
			message.Mode, message.Sponsoringnode = getRandomRingPosAndRandomSponsoringNode()
			nodeFunc(message)
		}
	}
	
	for k := range nodeChannelmap {
		fmt.Println(nodeChannelmap[k].NodeIdentifier, nodeChannelmap[k].Successor.NodeIdentifier, nodeChannelmap[k].Predecessor.NodeIdentifier)
	}
	
}

func nodeFunc(message Message) {
	
	if("join-ring" == message.Do) {
		joinNode(message.Mode, message.Sponsoringnode)
	}
}

func joinNode(ringPos int, sponsoringNodeId int) {

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

func getRandomRingPosAndRandomSponsoringNode() (int, int) {
	
	randPos := 1
	randSponsor := 1
	
	for {
		randPos = rand.Intn(9) + 1
		for k := range nodeChannelmap {
			if(randPos == k) {
				randPos = rand.Intn(9) + 1
			}
		}
		if(randPos > 0) {
			break
		}
	}
	
	for k := range nodeChannelmap {
		randSponsor = k
		if(k != 0) {
			break
		}
	}
	return randPos, randSponsor
}