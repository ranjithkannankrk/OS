package main
import (
	"math/rand"
	"time"
	"sync"
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
var nodeChannelVarMap = make(map[int]chan Message)
var mutex = &sync.Mutex{}


func main() {

	results := make(chan int, 100)
	if(len(nodeChannelmap) == 0) {
		ringPos := 1
		node := Node{}
		node.NodeIdentifier = ringPos
		node.Successor = &node
		node.Predecessor = &node
		node.KeyValue = make(map[int]int)
		mutex.Lock()
		nodeChannelmap[ringPos] = node
		mutex.Unlock()
		mutex.Lock()
		nodeChannelVarMap[ringPos] = make(chan Message)
		mutex.Unlock()
		mutex.Lock()
		go nodeFunc(nodeChannelVarMap[ringPos], results)
		mutex.Unlock()
	}
	
	
	action := "join-ring"
	for i := 0; i < 5; i++ {
		if(action == "join-ring") {
			message := Message{}
			message.Do = action
			mutex.Lock()
			message.Mode, message.Sponsoringnode = getRandomRingPosAndRandomSponsoringNode()
			mutex.Unlock()
			mutex.Lock()
			ch := nodeChannelVarMap[message.Sponsoringnode]
			mutex.Unlock()
			ch <- message	// for the chosen sponsoring node the request is sent and the respective go routine for the node is invoked when the message is inserted in the channel
		}
	}
	
	time.Sleep(100 * time.Second)
	for k := range nodeChannelmap {
		fmt.Println(nodeChannelmap[k].NodeIdentifier, nodeChannelmap[k].Successor.NodeIdentifier, nodeChannelmap[k].Predecessor.NodeIdentifier)
	}
	
}

//function which is used to perform node operations, this will be invoked as a go routine
func nodeFunc(jobs <-chan Message, results chan<- int) {
	for j := range jobs {
		if("join-ring" == j.Do) {
			mutex.Lock()
			joinNode(j.Mode, j.Sponsoringnode)
			mutex.Unlock()
			mutex.Lock()
			go nodeFunc(nodeChannelVarMap[j.Mode], results)	//recursively call the nodeFunc function for the respective node as a go routine when a new node joins the system
			mutex.Unlock()
		}
		results <- 1
	}
	time.Sleep(100 * time.Second)
}

//This function is invoked when the action is for a node to join the system
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
	nodeChannelVarMap[ringPos] = make(chan Message)

}

// This function updates the successor and predecessor of the neighboring nodes when a new node joins the system
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

//function used to find the successor and predecessor if a node, where ringpos is the node identifier for which 
//the successor and predecessor is to be found
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

//function used to get a random ring position for the node to join and also to randomly choose a sponsoring node
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