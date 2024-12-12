package p2p

// Peer is an interface the represent the remote node
type Peer interface{}

// Transport is anything that handles communication between the nodes in the network
// this can be of the form (TCP, UDB, WebSockets, ...)
type Transport interface{}
