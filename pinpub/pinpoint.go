package main

import (
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
)

type Publisher struct {
	PubID   string            `json:"PubID"`
	NodeKey peer.ID           `json: "NodeKey"`
	Topics  map[string]string `json: "Topics"`
}

type Pinpoint struct {
	RootCID   string    `json:"cid"`
	CreatedAt time.Time `json:"createdAt"`
	mu        *sync.Mutex
}
