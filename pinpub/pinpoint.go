package main

import (
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
)

// Publisher is a given publisher,
// that is, it reference the ID of the IPFS Node and
// its list of topics.
type Publisher struct {
	PubID   string            `json:"PubID"`
	NodeKey peer.ID           `json: "NodeKey"`
	Topics  map[string]string `json: "Topics"`
}

// Pinpoint is a given topic message,
// that is, it references the root CID of the IPFS object that
// contains the current state of the topic
type Pinpoint struct {
	RootCID   string    `json:"cid"`
	CreatedAt time.Time `json:"createdAt"`
	mu        *sync.Mutex
}
