# cjdns_munin
Munin plugins for CJDNS, written in Go.

## Plugins

### cjdns_memory
Tracks memory usage as reported by CJDNS.

### cjdns_link_quality
Charts changes in the link quality of direct peers.

### cjdns_traffic
Charts network usage of direct peers. 

## Install

    go get github.com/3M3RY/cjdns_munin
    sudo cp ${GOPATH}/bin/cjdns_munin /etc/munin/plugins/cjdns_memory
    sudo cp ${GOPATH}/bin/cjdns_munin /etc/munin/plugins/cjdns_link_quality
    sudo cp ${GOPATH}/bin/cjdns_munin /etc/munin/plugins/cjdns_traffic


## Configuration
#### /etc/munin/plugin-conf.d/cjdns
    [cjdns_*]
    env.cjdns_addr 127.0.0.1
    env.cjdns_port 11234
    env.cjdns_password xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
