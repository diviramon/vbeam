package main

import (
	"sync"
	"time"

	"github.com/ipfs/go-cid"
)

type Pinpoint struct {
	RootCID   cid.Cid   `json:"cid"`
	CreatedAt time.Time `json:"createdAt"`
	mu        *sync.Mutex
}
