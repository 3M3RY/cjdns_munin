// Copyright © 2013 Emery Hemingway

package main

import (
	"fmt"
	"github.com/3M3RY/go-cjdns/admin"
	"os"
	"strconv"
	"strings"
)

var modes = make(map[string]func(bool))

func main() {
	cmd := os.Args[0]
	i := strings.IndexRune(cmd, os.PathSeparator)
	for i != -1 {
		cmd = cmd[i+1:]
		i = strings.IndexRune(cmd, os.PathSeparator)
	}

	f, ok := modes[cmd]
	if !ok {
		fmt.Fprintf(os.Stderr, "Unknown command %q, supported commands:\n", cmd)
		for k, _ := range modes {
			fmt.Println(k)
		}
		os.Exit(1)
	}
	config := false
	for _, arg := range os.Args[1:] {
		if arg == "config" {
			config = true
		} else {
			fmt.Fprintln(os.Stderr, "unhandled argument", os.Args[1])
			os.Exit(1)
		}
	}

	f(config)
}

func ConnectToCjdns() *admin.Conn {
	port, _ := strconv.ParseUint(os.Getenv("cjdns_port"), 10, 16)
	cjdnsConf := &admin.CjdnsAdminConfig{
		Addr:     os.Getenv("cjdns_addr"),
		Port:     int(port),
		Password: os.Getenv("cjdns_password"),
	}
	if cjdnsConf.Addr == "" || cjdnsConf.Port == 0 || cjdnsConf.Password == "" {
		fmt.Fprintln(os.Stderr, "failed to read CJDNS admin port configuration from environment")
		os.Exit(1)
	}

	admin, err := admin.Connect(cjdnsConf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return admin
}
