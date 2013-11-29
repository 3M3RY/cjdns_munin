// Copyright © 2013 Emery Hemingway

package main

import (
	"fmt"
	"github.com/inhies/go-cjdns/key"
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
		fmt.Println("graph_title Cjdns Peer Link Quality")
		fmt.Println("graph_args --lower-limit 0")
		fmt.Println("graph_scale no")
		fmt.Println("graph_vlabel Link Quality")
		fmt.Println("graph_category network")

		for _, peer := range stats {
			peerName := strings.Split(peer.PublicKey, ".")[0]

			label, _ := key.PubKeyToIP(peer.PublicKey)
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
		peerName := strings.Split(peer.PublicKey, ".")[0]
		//ip := key.DecodePublic(peer.PublicKey).IP.String()
		ip, _ := key.PubKeyToIP(peer.PublicKey)
		for _, r := range table {
			if r.IP == ip {
				fmt.Printf("_%s_link.value %s\n", peerName, r.Link)
				break
			}
		}
	}
}
