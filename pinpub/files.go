// files.go contains stuff related to handling the local filesystem files
// like scanning and loading the target directory
//
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	go_ipfs_files "github.com/ipfs/go-ipfs-files"
	icore "github.com/ipfs/interface-go-ipfs-core"
)

const inputBasePath = "./target/"

func getUnixfsNode(path string) (go_ipfs_files.Node, error) {
	st, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	f, err := go_ipfs_files.NewSerialFile(path, false, st)
	if err != nil {
		return nil, err
	}

	return f, err
}

// WatchDir watches the direcotry and pushes the latest version of
// the pinpoint to the topic channel
func WatchDir(cfg *config, ipfs icore.CoreAPI, topics map[string]*Pinpoint, cid2topic map[string]string) {

	for {

		for label, val := range cfg.Topics {
			topic, ok := topics[label]
			if !ok {
				panic(fmt.Errorf("no topic for the label=%s, bad", label))
			}

			someDir, err := getUnixfsNode(val.SrcDir)
			if err != nil {
				panic(fmt.Errorf("could not find the directory: %s", err))
			}
			cidDir, err := ipfs.Unixfs().Add(context.Background(), someDir)
			someDir.Close()

			if err != nil {
				panic(fmt.Errorf("could not add directory: %s", err))
			}
			if cidDir.Cid().String() != topic.RootCID {
				fmt.Printf("updating the topic %s RootCID to %s\n", label, cidDir.Cid())
				topic.mu.Lock()
				cidString := cidDir.Cid().String()
				topic.RootCID = cidString
				cid2topic[cidString] = label
				topic.CreatedAt = time.Now()
				topic.mu.Unlock()
			}
		}
		time.Sleep(time.Second)
	}
}
