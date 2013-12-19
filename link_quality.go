// Copyright © 2013 Emery Hemingway

package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func init() {
	modes["cjdns_link_quality"] = cjdnsLinkQuality
}

func cjdnsLinkQuality(config bool) {
	admin := ConnectToCjdns()

	stats, err := admin.InterfaceController_peerStats()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if config {
		fmt.Println("graph_title CJDNS Peer Link Quality")
		fmt.Println("graph_args --lower-limit 0")
		fmt.Println("graph_scale no")
		fmt.Println("graph_vlabel link quality")
		fmt.Println("graph_category network")

		for _, peer := range stats {
			peerName := strings.Split(peer.PublicKey.String(), ".")[0]
			label := peer.PublicKey.IP().String()
			names, _ := net.LookupAddr(label)
			if len(names) > 0 {
				label = names[0]
			}
			fmt.Printf("_%s_link.label %s\n", peerName, label)
			//fmt.Printf("_%s_in.type GUAGE\n", peerName)
		}
		return
	}

	table, err := admin.NodeStore_dumpTable()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	table.SortByPath()

	for _, peer := range stats {
		peerName := strings.Split(peer.PublicKey.String(), ".")[0]
		ip := peer.PublicKey.IP()
		for _, r := range table {
			if ip.Equal(*r.IP) {
				fmt.Printf("_%s_link.value %s\n", peerName, r.Link)
				break
			}
		}
	}
}
