package main

import (
	"context"
	"fmt"

	files "github.com/ipfs/go-ipfs-files"
	icorepath "github.com/ipfs/interface-go-ipfs-core/path"
)

func main() {

	fmt.Println("-- Getting an IPFS node running -- ")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Spawn a node using the default path (~/.ipfs)
	// fmt.Println("Spawning node on default repo")
	// ipfs, err := spawnDefault(ctx)
	// if err != nil {
	// 	fmt.Println("No IPFS repo available on the default path -", err)
	// 	return
	// }

	// Spawn a node using a temporary repo
	fmt.Println("Spawning node on a temporary repo")
	ipfs, err := spawnEphemeral(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to spawn ephemeral node: %s", err))
	}

	fmt.Println("IPFS node is running.")

	// Adding a file
	inputBasePath := "./target/"
	inputPathFile := inputBasePath + "whitepaper.pdf"
	inputPathDirectory := inputBasePath + "test-dir"

	someFile, err := getUnixfsNode(inputPathFile)
	if err != nil {
		panic(fmt.Errorf("could not find the File: %s", err))
	}

	cidFile, err := ipfs.Unixfs().Add(ctx, someFile)
	if err != nil {
		panic(fmt.Errorf("could not add File: %s", err))
	}

	fmt.Printf("File added, its CID is: %s\n", cidFile.String())

	// Adding a directory
	someDir, err := getUnixfsNode(inputPathDirectory)
	if err != nil {
		panic(fmt.Errorf("could not find the directory: %s", err))
	}

	cidDir, err := ipfs.Unixfs().Add(ctx, someDir)
	if err != nil {
		panic(fmt.Errorf("could not add directory: %s", err))
	}

	fmt.Printf("Added directory to IPFS with CID %s\n", cidDir.String())

	// Connecting to peers
	bootstrapNodes := []string{
		// IPFS Bootstrapper nodes.
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",

		// IPFS Cluster Pinning nodes
		"/ip4/138.201.67.219/tcp/4001/p2p/QmUd6zHcbkbcs7SMxwLs48qZVX3vpcM8errYS7xEczwRMA",
		"/ip4/138.201.67.219/udp/4001/quic/p2p/QmUd6zHcbkbcs7SMxwLs48qZVX3vpcM8errYS7xEczwRMA",
		"/ip4/138.201.67.220/tcp/4001/p2p/QmNSYxZAiJHeLdkBg38roksAR9So7Y5eojks1yjEcUtZ7i",
		"/ip4/138.201.67.220/udp/4001/quic/p2p/QmNSYxZAiJHeLdkBg38roksAR9So7Y5eojks1yjEcUtZ7i",
		"/ip4/138.201.68.74/tcp/4001/p2p/QmdnXwLrC8p1ueiq2Qya8joNvk3TVVDAut7PrikmZwubtR",
		"/ip4/138.201.68.74/udp/4001/quic/p2p/QmdnXwLrC8p1ueiq2Qya8joNvk3TVVDAut7PrikmZwubtR",
		"/ip4/94.130.135.167/tcp/4001/p2p/QmUEMvxS2e7iDrereVYc5SWPauXPyNwxcy9BXZrC1QTcHE",
		"/ip4/94.130.135.167/udp/4001/quic/p2p/QmUEMvxS2e7iDrereVYc5SWPauXPyNwxcy9BXZrC1QTcHE",

		// Add local nodes here
	}

	go connectToPeers(ctx, ipfs, bootstrapNodes)

	fmt.Println("Connected to peer nodes as bootstrappers.")

	// Retrieving a File
	testCIDstr := "QmV9tSDx9UiPeWExXEeH6aoDvmihvx6jD5eLb4jbTaKGps" // ipfs whitepaper
	outputPath := inputBasePath + "output-dir/whitepaper.pdf"
	testCID := icorepath.New(testCIDstr)

	rootNode, err := ipfs.Unixfs().Get(ctx, testCID)
	if err != nil {
		panic(fmt.Errorf("could not get the test file from the swarm"))
	}

	err = files.WriteTo(rootNode, outputPath)
	if err != nil {
		panic(fmt.Errorf("could not save the fetched CID: %s", err))
	}

	fmt.Printf("Test file retrieved and saved to: %s\n", outputPath)

	fmt.Println("\nEsto es todo amigos!")
}
