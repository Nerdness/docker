package main

import (
	"github.com/Nerdness/docker/sysinit"
)

var (
	GITCOMMIT string
	VERSION   string
)

func main() {
	// Running in init mode
	sysinit.SysInit()
	return
}
