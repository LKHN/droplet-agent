// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/digitalocean/droplet-agent/internal/log"
	"github.com/digitalocean/droplet-agent/internal/metadata/watcher"
)

func newMetadataWatcher(cfg *watcher.Conf) watcher.MetadataWatcher {
	log.Info("Launching SSH Port Knocking Watcher")
	return watcher.NewSSHWatcher(cfg)
}
