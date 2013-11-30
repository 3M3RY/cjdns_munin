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
	modes["cjdns_traffic"] = cjdnsTraffic
}

func cjdnsTraffic(config bool) {
	admin := ConnectToCjdns()

	stats, err := admin.InterfaceController_peerStats()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if config {
		fmt.Println("graph_title CJDNS Peer Traffic")
		fmt.Println("graph_args --base 1000") // --lower-limit 0")
		fmt.Println("graph_scale yes")
		fmt.Println("graph_vlabel bytes")
		fmt.Println("graph_category network")
	}

	for _, peer := range stats {
		peerName := strings.Split(peer.PublicKey, ".")[0]

		if config {
			pubKey, _ := key.DecodePublic(peer.PublicKey)
			label := pubKey.IP().String()
			names, _ := net.LookupAddr(label)
			if len(names) > 0 {
				label = names[0]
			}

			fmt.Printf("_%s_in.label %s\n", peerName, label)
			fmt.Printf("_%s_in.type DERIVE\n", peerName)
			fmt.Printf("_%s_in.graph no\n", peerName)
			fmt.Printf("_%s_in.min 0\n", peerName)

			fmt.Printf("_%s_out.label %s\n", peerName, label)
			fmt.Printf("_%s_out.type DERIVE\n", peerName)
			fmt.Printf("_%s_out.negative _%s_in\n", peerName, peerName)
			fmt.Printf("_%s_out.min 0\n", peerName)
		} else {
			fmt.Printf("_%s_in.value %d\n", peerName, peer.BytesIn)
			fmt.Printf("_%s_out.value %d\n", peerName, peer.BytesOut)
		}
	}
	os.Exit(0)
}
