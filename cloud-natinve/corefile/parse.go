package main

import (
	"github.com/coredns/corefile-migration/migration/corefile"
)

func parseCorefile(data string) (*corefile.Corefile, error) {
	return corefile.New(data)
}

// setUpstream find forward plugin and edit upstream
func setUpstream(data *corefile.Corefile, upstream string) (string, error) {
	for _, server := range data.Servers {
		// just edit .:53 zone
		if len(server.DomPorts) != 0 && server.DomPorts[0] == ".:53" {
			for _, plugin := range server.Plugins {
				// find forward plugin and edit upstream
				if plugin.Name == "forward" {
					plugin.Args = []string{".", upstream}
				}
			}
		}
	}
	return data.ToString(), nil
}
