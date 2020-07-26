package main

import (
	"context"
	"fmt"
	"os"

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

	for topic := range cfg.Subscriptions {
		fmt.Printf("scanning topic %s...\n", topic)
		pp, err := scanTopic(topic)
		if err != nil {
			errorOut(fmt.Errorf("failed to scan topic %v: %v", topic, err))
		}
		err = ipfs.Pin().Add(context.Background(), ipfs_path.New(pp.RootCID))
		if err != nil {
			errorOut(fmt.Errorf("unable to pin %v: %v", pp.RootCID, err))
		}
	}
}
