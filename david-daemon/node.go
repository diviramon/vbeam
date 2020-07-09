package main

import (
	"context"
	"fmt"

	config "github.com/ipfs/go-ipfs-config"
	libp2p "github.com/ipfs/go-ipfs/core/node/libp2p"
	icore "github.com/ipfs/interface-go-ipfs-core"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi" // This package is needed so that all the preloaded plugins are loaded automatically
	"github.com/ipfs/go-ipfs/repo/fsrepo"
)

// Creates an IPFS node and returns its coreAPI
func createNode(ctx context.Context, repoPath string) (icore.CoreAPI, error) {
	// Open the repo
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}

	//Construct the node
	nodeOptions := &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTOption, // full DHT node (both fetching and storing DHT Records
		// Routing: libp2p.DHTClientOption, // client DHT node (only fetching records)
		Repo: repo,
	}

	node, err := core.NewNode(ctx, nodeOptions)
	if err != nil {
		return nil, err
	}

	// Attach the Core API to the constructed node
	return coreapi.NewCoreAPI(node)
}

// Spawns a node on the default repo location
func spawnDefault(ctx context.Context) (icore.CoreAPI, error) {
	defaultPath, err := config.PathRoot()
	if err != nil {
		return nil, err
	}

	if err := setupPlugins(defaultPath); err != nil {
		return nil, err
	}

	return createNode(ctx, defaultPath)
}

// Spawns an ephimeral node using a tmp repo
func spawnEphemeral(ctx context.Context) (icore.CoreAPI, error) {
	if err := setupPlugins(""); err != nil {
		return nil, err
	}

	// Create a Temporary Repo
	repoPath, err := createTempRepo(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to create temp repo: %s", err)
	}

	// Spawn ephemeral IPFS node
	return createNode(ctx, repoPath)
}
