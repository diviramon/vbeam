package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"

	config "github.com/ipfs/go-ipfs-config"

	"github.com/ipfs/go-ipfs/plugin/loader"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
)

func setupPlugins(externalPluginsPath string) error {
	// Load external plugins
	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
	if err != nil {
		return fmt.Errorf("Error loading plugins: %s", err)
	}

	// Load preloaded plugins
	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("Error initializing preloaded plugins: %s", err)
	}

	return nil
}

func createTempRepo(ctx context.Context) (string, error) {
	repoPath, err := ioutil.TempDir("", "ipfs-shell")
	if err != nil {
		return "", fmt.Errorf("Failed to get temp dir: %s", err)
	}

	//Create a temp config with default options and 2048 key pair
	cfg, err := config.Init(ioutil.Discard, 2048)
	if err != nil {
		return "", err
	}

	//Create the repo with the config
	err = fsrepo.Init(repoPath, cfg)
	if err != nil {
		return "", fmt.Errorf("Failed to init ephemeral repo: %s", err)
	}

	return repoPath, nil
}
