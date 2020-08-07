package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/diviramon/vbeam/ipfsglue"
	ipfs_path "github.com/ipfs/interface-go-ipfs-core/path"
)

func errorOut(err error) {
	fmt.Printf("exiting with error:%s\n", err)
	os.Exit(1)
}

func main() {
	fmt.Printf("this is v-beam gateway\n")
	cfg, err := loadConfig()
	if err != nil {
		errorOut(err)
	}
	ipfs, err := ipfsglue.SpawnIPFS(cfg.IpfsRepoPath)

	if err != nil {
		errorOut(fmt.Errorf("cannot spawn the IPFS node: %v", err))
	}

	topics := make(map[string]*pinpoint)
	for topic := range cfg.Subscriptions {
		topics[topic] = new(pinpoint)
	}
	for {
		for label := range cfg.Subscriptions {
			// fmt.Printf("scanning topic %s...\n", label)
			pp, err := scanTopic(label)
			if err != nil {
				errorOut(fmt.Errorf("failed to scan topic %v: %v", label, err))
			}
			topic, ok := topics[label]
			if !ok {
				panic(fmt.Errorf("no topic for the label=%s, bad", label))
			}
			if pp.RootCID != topic.RootCID {
				if topic.RootCID != "" {
					err = ipfs.Pin().Rm(context.Background(), ipfs_path.New(topic.RootCID))
					if err != nil {
						errorOut(fmt.Errorf("unable to unpin old topic with label: %s and CID: %s", label, topic.RootCID))
					}
				}
				err = ipfs.Pin().Add(context.Background(), ipfs_path.New(pp.RootCID))
				if err != nil {
					errorOut(fmt.Errorf("unable to pin %v: %v", pp.RootCID, err))
				}
				topic.RootCID = pp.RootCID
				topic.CreatedAt = time.Now()
				fmt.Printf("updated topic %s to root CID %s...\n", label, pp.RootCID)
			}
		}
		time.Sleep(time.Second)
	}
}
