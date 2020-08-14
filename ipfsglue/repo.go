package ipfsglue

import (
	"fmt"
	"path/filepath"

	"github.com/ipfs/go-ipfs/plugin/loader"
)

// setupPlugins loads the plugins found in the input
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

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("Error initializing plugins: %s", err)
	}

	return nil
}
