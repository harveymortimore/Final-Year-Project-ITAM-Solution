package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/host"
)

func main() {
	info, err := host.Info()
	if err != nil {
		panic(err)
	}
	fmt.Printf("host=%s os=%s id=%s\n", info.Hostname, info.OS, info.HostID)
}