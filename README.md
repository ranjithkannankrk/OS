# OS
Chord nodes are implemented as goroutines in the GoLang. Each node has a randomly assigned GoLang channel variable for receiving messages from other nodes. Each channel has a unique (string) identifier; a node uses its channel's id as its own identifier. Further, each node maintains a bucket (list of) (key,value) pairs of a hash table that is distributed among the nodes that are members of the Chord ring. There is no limit on a node's bucket size besides the available memory to GoLang processes, while the keys and values are assumed to be strings. 
The GoLang routine (aka coordinator) spawns some Chord nodes, and then, instructs them to join/leave the Chord ring, as well as get/put/remove key-value pairs from/to the distributed hash table.
Chord nodes reveive JSON request messages from the coordinator or other Chord nodes and respond to the specified response channel directly. We assume the time it takes a node to respond to any message is a random variable. The JSON request messages among the Chord nodes and the coordinator are as follows:

{"do": "join-ring", "sponsoring-node": "channel-id" } instructing the receipient node to join the Chord ring by contacting the (existing) Chord sponsoring node listening on the given channel.
{"do": "leave-ring" "mode": "immediate or orderly"} instructing the receipient node to leave the ring immediately (without informing any other nodes) or in an orderly manner (by informing other nodes and transferring its bucket contents to others)
{"do": "stabilize-ring" }
{"do": "init-ring-fingers" }
{"do": "fix-ring-fingers" }
{"do": "ring-notify", "respond-to": "channel-id" }
{"do": "get-ring-fingers", "respond-to": "channel-id" }
{"do": "find-ring-successor", "respond-to": "channel-id"}
{"do": "find-ring-predecessor", "respond-to": "channel-id"}
{"do": "put", "data": { "key" : "a key", "value" : "a value" }, "respond-to": "channel-id"} instructing the receipient node to store the given (key,value) pair in the appropriate ring node.
{"do": "get", "data": { "key" : "a key" }, "respond-to": "channel-id"} instructing the receipient node to retrieve the value associated with the key stored in the ring.
{"do": "remove", "data": { "key" : "a key" }, "respond-to": "channel-id"} instructing the receipient node to remove the (key,value) pair from the ring.
