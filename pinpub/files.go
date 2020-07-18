// files.go contains stuff related to handling the local filesystem files
// like scanning and loading the target directory
//
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
)

const inputBasePath = "./target/"

func getUnixfsNode(path string) (files.Node, error) {
	st, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	f, err := files.NewSerialFile(path, false, st)
	if err != nil {
		return nil, err
	}

	return f, err
}

// WatchDir watches the direcotry and pushes the latest version of
// the pinpoint to the topic channel
func WatchDir(path string, ipfs icore.CoreAPI, topic *Pinpoint) {
	if path == "" {
		path = inputBasePath
	}

	for {
		// Adding a directory
		someDir, err := getUnixfsNode(path)
		if err != nil {
			panic(fmt.Errorf("could not find the directory: %s", err))
		}

		cidDir, err := ipfs.Unixfs().Add(context.Background(), someDir)
		if err != nil {
			panic(fmt.Errorf("could not add directory: %s", err))
		}

		if cidDir.Cid() != topic.RootCID {
			fmt.Println("updating the topic RootCID to ", cidDir.Cid())
			topic.mu.Lock()
			topic.RootCID = cidDir.Cid()
			topic.CreatedAt = time.Now()
			topic.mu.Unlock()
		}
		time.Sleep(1 * time.Second)
	}
}
