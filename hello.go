package main
import (
	"fmt"
	"encoding/json"
)

type Node struct {
	Successor      *Node
	KeyValue       map[int]int
	FingerTable    map[int]int
	Predecessor    *Node
	NodeIdentifier int
}

type Message struct {
	Do             string	`json:"do"`
	Sponsoringnode int		`json:"sponsoring-node"`
	Mode           int		`json:"mode"`
	Respondto      string	`json:"respond-to"`
}


var chrodRingNodes = []*Node{}
var nodeChannelmap = make(map[int]chan Message)


func main() {


	results := make(chan int, 100)
	var chans [6]chan Message
	for i := range chans {
		chans[i] = make(chan Message)
	}


    fmt.Println("hello world")
	ringPos := 1
	if(len(chrodRingNodes) == 0) {
		node := Node{}
		node.NodeIdentifier = ringPos
		node.Successor = &node
		node.Predecessor = &node
		node.KeyValue = make(map[int]int)
		chrodRingNodes = append(chrodRingNodes, &node)
		nodeChannelmap[ringPos] = chans[0]
	}
	
	
	jsonMessage := string(`{"do": "join-ring","sponsoring-node":1, "mode": 3}`)
    message := Message{}
    json.Unmarshal([]byte(jsonMessage), &message)
	
	go nodeFunc(nodeChannelmap[message.Sponsoringnode], results)
	ch := nodeChannelmap[message.Sponsoringnode]
	ch <- message
	
	jsonMessage1 := string(`{"do": "join-ring","sponsoring-node":3, "mode": 5}`)
    message1 := Message{}
    json.Unmarshal([]byte(jsonMessage1), &message1)
	ch1 := nodeChannelmap[message.Sponsoringnode]
	ch1 <- message1
}

//function used to perform all the node operations
func nodeFunc(jobs <-chan Message, results chan<- int) {
	for j := range jobs {
		fmt.Println("check")
		fmt.Println(j)
		
		//snippet invoked when the node operation is join
		if("join-ring" == j.Do) {
			makeNode(j)
			
			//recursively call the node operation function for the respective channel
			go nodeFunc(nodeChannelmap[j.Mode], results)
		}
		fmt.Println(len(nodeChannelmap))
		results <- 1
	}
}

//function used to create a new node in the chord ring and also updates the successor and predecessor of the node to be inserted as well as the affected nodes
func makeNode(msg Message) {
	ringPos := msg.Mode
	newNode := Node{}
	startNode := Node{}
	newNode.NodeIdentifier = ringPos
	sponsoringNode := Node{}
	for _, element := range chrodRingNodes {
		if(msg.Sponsoringnode == element.NodeIdentifier) {
			sponsoringNode = *element
			}
		}
		newNode.Successor, newNode.Predecessor, startNode = findSuccessorAndPredecessorForJoin(ringPos, sponsoringNode)
		
		//snippet used to update the successor and predecessor of the nodes which were disturbed when a new node was brought inside
		if(ringPos < msg.Sponsoringnode) {
			startNode.Successor.Predecessor = &newNode
			startNode.Successor = &newNode
		}
		if(ringPos > msg.Sponsoringnode) {
			startNode.Predecessor.Successor = &newNode
			startNode.Predecessor = &newNode
		}
	
	//add the newly created node to the global dictionary
	chrodRingNodes = append(chrodRingNodes, &newNode)
	
	//Assign a channel to the newly created node
	nodeChannelmap[ringPos] = make(chan Message)
}

func findSuccessorAndPredecessorForJoin(ringPos int, sponsoringNode Node) (*Node, *Node, Node) {

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