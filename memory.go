// Copyright © 2013 Emery Hemingway

package main

import (
	"fmt"
	"os"
)

func init() {
	modes["cjdns_memory"] = cjdnsMemory
}

func cjdnsMemory(config bool) {
	admin := ConnectToCjdns()

	if config {
		fmt.Println("graph_title CJDNS memory")
		fmt.Println("graph_args --base 1024 --lower-limit 0")
		fmt.Println("graph_scale yes")
		fmt.Println("graph_vlabel memory in bytes")
		fmt.Println("graph_category network")
		fmt.Println("m.label memory")
		return
	}

	mem, err := admin.Memory()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to retrieve memory count,", err.Error())
	}
	fmt.Printf("m.value %d\n", mem)
}
