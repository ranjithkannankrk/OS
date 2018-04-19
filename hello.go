package main
import (
	"fmt"
	"hash/fnv"
)

type Node struct {
	Successor      *Node
	KeyValue       map[int]int
	FingerTable    map[int]int
	Predecessor    *Node
	NodeIdentifier int
}


func main() {
    fmt.Println("hello world")
}


//generate hash value
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}