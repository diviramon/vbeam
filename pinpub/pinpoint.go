package main

import (
	"sync"

	"github.com/ipfs/go-cid"
)

type Pinpoint struct {
	RootCID cid.Cid
	Mutex   *sync.Mutex
}
